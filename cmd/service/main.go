package main

import (
	"context"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/docker/docker/client"
	"github.com/sirupsen/logrus"

	"github.com/godverv/Velez/internal/backservice/configuration"
	"github.com/godverv/Velez/internal/backservice/env"
	"github.com/godverv/Velez/internal/backservice/portainer"
	"github.com/godverv/Velez/internal/backservice/security"
	"github.com/godverv/Velez/internal/backservice/watchtower"
	"github.com/godverv/Velez/internal/clients/docker"
	grpcClients "github.com/godverv/Velez/internal/clients/grpc"
	"github.com/godverv/Velez/internal/config"
	"github.com/godverv/Velez/internal/cron"
	"github.com/godverv/Velez/internal/service"
	"github.com/godverv/Velez/internal/service/service_manager"
	"github.com/godverv/Velez/internal/service/service_manager/container_manager_v1/port_manager"
	"github.com/godverv/Velez/internal/transport"
	"github.com/godverv/Velez/internal/transport/grpc"
	"github.com/godverv/Velez/internal/utils/closer"
	"github.com/godverv/Velez/pkg/velez_api"
)

func main() {
	logrus.Println("starting app")

	// Core dependencies
	aCore := mustInitCore()

	// Verv Environment
	mustInitEnvironment(aCore)

	// Startup ctx
	{
		if aCore.cfg.GetAppInfo().StartupDuration == 0 {
			logrus.Fatalf("no startup duration in config")
		}

		var cancel func()
		aCore.ctx, cancel = context.WithTimeout(context.Background(), aCore.cfg.GetAppInfo().StartupDuration)
		closer.Add(func() error { cancel(); return nil })
	}

	// Service layer
	serviceManager := mustInitServiceManager(aCore)

	// Back service
	initBackServices(aCore, serviceManager.GetContainerManagerService())

	// API
	mgr := mustInitAPI(
		aCore,
		serviceManager,
	)

	err := mgr.Start(aCore.ctx)
	if err != nil {
		logrus.Fatalf("error starting api: %s", err)
	}

	ctx := context.Background()

	waitingForTheEnd()
	logrus.Println("shutting down the app")

	err = mgr.Stop(ctx)
	if err != nil {
		logrus.Errorf("error stopping api: %s", err)
	}

	if err = closer.Close(); err != nil {
		logrus.Fatalf("errors while shutting down application %s", err)
	}
}

// rscli comment: an obligatory function for tool to work properly.
// must be called in the main function above
// also this is an LP song name reference, so no rules can be applied to the function name
func waitingForTheEnd() {
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-done
}

type applicationCore struct {
	ctx context.Context

	cfg config.Config

	dockerAPI       client.CommonAPIClient
	securityManager security.Manager
	portManager     *port_manager.PortManager
}

func mustInitCore() (c applicationCore) {
	var err error

	// Load config
	{
		c.cfg, err = config.Load()
		if err != nil {
			logrus.Fatalf("error reading config %s", err.Error())
		}

	}

	// Docker api
	{
		c.dockerAPI, err = docker.NewClient()
		if err != nil {
			logrus.Fatalf("erorr getting docker api client: %s", err)
		}
		closer.Add(c.dockerAPI.Close)
	}

	// Security access layer
	if !c.cfg.GetEnvironment().DisableAPISecurity {
		c.securityManager = security.NewSecurityManager(c.cfg.GetEnvironment().CustomPassToKey)

		err = c.securityManager.Start()
		if err != nil {
			logrus.Fatalf("error starting security manager: %s", err)
		}

		closer.Add(c.securityManager.Stop)
	}

	// port manager
	{
		c.portManager, err = port_manager.NewPortManager(context.Background(), c.cfg, c.dockerAPI)
		if err != nil {
			logrus.Fatalf("error creating port manager %s", err)
		}
	}
	return
}

func mustInitEnvironment(aCore applicationCore) {
	err := env.StartNetwork(aCore.dockerAPI)
	if err != nil {
		logrus.Fatalf("error creating network: %s", err)
	}

	err = env.StartVolumes(aCore.dockerAPI)
	if err != nil {
		logrus.Fatalf("error creating volumes %s", err)
	}

	if !aCore.cfg.GetEnvironment().NodeMode {
		return
	}

	var portToExposeTo string
	if aCore.cfg.GetEnvironment().ExposeMatreshkaPort {
		p := uint64(aCore.cfg.GetEnvironment().MatreshkaPort)

		if p == 0 {
			portFromPool := aCore.portManager.GetPort()
			if portFromPool == nil {
				logrus.Fatalf("no available port for config to expose")
				return
			}

			p = uint64(*portFromPool)
		}

		portToExposeTo = strconv.FormatUint(p, 10)
	}

	conf := configuration.New(aCore.dockerAPI, portToExposeTo)
	err = conf.Start()
	if err != nil {
		logrus.Fatalf("error launching config backservice: %s", err)
	}

	go cron.KeepAlive(context.Background(), conf)
}

func mustInitServiceManager(aCore applicationCore) service.Services {
	matreshkaApi, err := grpcClients.NewMatreshkaBeAPIClient(aCore.ctx, aCore.cfg)
	if err != nil {
		logrus.Fatalf("error getting matreshka api: %s", err)
	}

	services, err := service_manager.New(
		aCore.cfg,
		aCore.dockerAPI,
		matreshkaApi,
		aCore.portManager,
	)
	if err != nil {
		logrus.Fatalf("error creating service manager: %s", err)
	}

	logrus.Warn("shut down on exit is ", aCore.cfg.GetEnvironment().ShutDownOnExit)

	if aCore.cfg.GetEnvironment().ShutDownOnExit {
		closer.Add(smerdsDropper(services.GetContainerManagerService()))
	}

	return services
}

func initBackServices(aCore applicationCore, cm service.ContainerManager) {
	ctx, c := context.WithCancel(context.Background())
	closer.Add(func() error {
		c()
		return nil
	})

	if aCore.cfg.GetEnvironment().WatchTowerEnabled {
		go cron.KeepAlive(ctx, watchtower.New(aCore.cfg, cm))
	}

	if aCore.cfg.GetEnvironment().PortainerEnabled {
		go cron.KeepAlive(ctx, portainer.New(cm))
	}
}

func smerdsDropper(manager service.ContainerManager) func() error {
	return func() error {
		logrus.Infof("ShutDownOnExit env variable is set to TRUE. Dropping launched smerds")
		logrus.Infof("Listing launched smerds")
		ctx := context.Background()

		smerds, err := manager.ListSmerds(ctx, &velez_api.ListSmerds_Request{})
		if err != nil {
			return err
		}

		names := make([]string, 0, len(smerds.Smerds))

		for _, sm := range smerds.Smerds {
			names = append(names, sm.Name)
		}

		logrus.Infof("%d smerds is active. %v", len(smerds.Smerds), names)

		dropReq := &velez_api.DropSmerd_Request{
			Uuids: make([]string, len(smerds.Smerds)),
		}

		for i := range smerds.Smerds {
			dropReq.Uuids[i] = smerds.Smerds[i].Uuid
		}

		logrus.Infof("Dropping %d smerds", len(smerds.Smerds))

		dropSmerds, err := manager.DropSmerds(ctx, dropReq)
		if err != nil {
			return err
		}

		logrus.Infof("%d smerds dropped successfully", len(dropSmerds.Successful))
		if len(dropSmerds.Successful) != 0 {
			logrus.Infof("Dropped smerds: %v", dropSmerds.Successful)
		}

		if len(dropSmerds.Failed) != 0 {
			logrus.Errorf("%d smerds failed to drop", len(dropSmerds.Failed))
			for _, f := range dropSmerds.Failed {
				logrus.Errorf("error dropping %s. Cause: %s", f.Uuid, f.Cause)
			}
		}

		return nil
	}
}

func mustInitAPI(
	aCore applicationCore,
	services service.Services,
) transport.Server {
	mgr := transport.NewManager()

	grpcConf, err := aCore.cfg.GetServers().GRPC(config.ServerGrpc)
	if err != nil {
		logrus.Fatalf("error getting grpc from config: %s", err)
	}

	srv, err := grpc.NewServer(
		aCore.cfg,
		grpcConf,
		services,
		aCore.securityManager,
	)
	if err != nil {
		logrus.Fatalf("error creating grpc server: %s", err)
	}

	mgr.AddServer(srv)

	return mgr
}
