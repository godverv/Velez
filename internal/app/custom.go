// Code generated by RedSock CLI.
// DO EDIT, DON'T DELETE THIS FILE.

package app

import (
	"context"

	"github.com/sirupsen/logrus"
	errors "go.redsock.ru/rerrors"
	"go.redsock.ru/toolbox/closer"
	"google.golang.org/grpc"

	"github.com/godverv/Velez/internal/backservice/service_discovery"
	"github.com/godverv/Velez/internal/clients"
	"github.com/godverv/Velez/internal/clients/matreshka"
	"github.com/godverv/Velez/internal/pipelines"
	"github.com/godverv/Velez/internal/security"
	"github.com/godverv/Velez/internal/service"
	"github.com/godverv/Velez/internal/service/service_manager"
	"github.com/godverv/Velez/internal/transport/control_plane_api_impl"
	"github.com/godverv/Velez/internal/transport/velez_api_impl"
	"github.com/godverv/Velez/pkg/docs"
	"github.com/godverv/Velez/pkg/velez_api"
)

type Custom struct {
	// NodeClients - hardware scanner, docker and wrappers
	NodeClients clients.NodeClients

	// Service discovery client
	ServiceDiscovery service_discovery.ServiceDiscovery
	// Configuration client
	MatreshkaClient matreshka.Client
	// ClusterClients - contains verv cluster's dependencies
	ClusterClients clients.ClusterClients

	// Services - contains business logic services
	Services  service.Services
	Pipeliner pipelines.Pipeliner
	// Api implementation
	ApiGrpcImpl         *velez_api_impl.Impl
	ControlPlaneApiImpl *control_plane_api_impl.Impl
}

func (c *Custom) Init(a *App) (err error) {
	c.NodeClients, err = clients.NewNodeClientsContainer(a.Ctx, a.Cfg)
	if err != nil {
		return errors.Wrap(err, "error initializing internal clients")
	}

	err = c.setupVervNodeEnvironment()
	if err != nil {
		return errors.Wrap(err, "error setting up node environment")
	}

	err = c.initServiceDiscovery(a)
	if err != nil {
		return errors.Wrap(err, "error initializing service discovery")
	}

	err = c.initConfigurationService(a)
	if err != nil {
		return errors.Wrap(err, "error initializing configuration service")
	}

	c.ClusterClients = clients.NewClusterClientsContainer(c.ServiceDiscovery, c.MatreshkaClient)

	c.initVelezServices(a)

	err = c.initApiServer(a)
	if err != nil {
		return errors.Wrap(err)
	}

	return nil
}

func (c *Custom) Start(ctx context.Context) error {

	return nil
}

func (c *Custom) Stop() error {

	return nil
}

func (c *Custom) initVelezServices(a *App) {
	var err error
	c.Services, err = service_manager.New(a.Ctx, c.NodeClients, c.ClusterClients)
	if err != nil {
		logrus.Fatalf("error initializing service manager: %v", err)
	}

	c.Pipeliner = pipelines.NewPipeliner(c.NodeClients.Docker(), c.NodeClients.PortManager(), c.Services)

	logrus.Warn("shut down on exit is set to: ", a.Cfg.Environment.ShutDownOnExit)
	if a.Cfg.Environment.ShutDownOnExit {
		closer.Add(smerdsDropper(c.Services))
	}
}

func (c *Custom) initApiServer(a *App) error {
	c.ApiGrpcImpl = velez_api_impl.NewImpl(a.Cfg, c.Services, c.Pipeliner)
	c.ControlPlaneApiImpl = control_plane_api_impl.New(c.ServiceDiscovery)

	var opts []grpc.ServerOption
	if !a.Cfg.Environment.DisableAPISecurity {
		opts = append(opts, security.GrpcIncomingInterceptor(c.NodeClients.SecurityManager().ValidateKey))
	}

	a.ServerMaster.AddImplementation(c.ApiGrpcImpl, opts...)
	a.ServerMaster.AddImplementation(c.ControlPlaneApiImpl, opts...)

	a.ServerMaster.AddHttpHandler(docs.Swagger())

	return nil
}

func smerdsDropper(manager service.Services) func() error {
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
