package service_discovery

import (
	"context"
	"strings"
	"time"

	rtb "github.com/Red-Sock/toolbox"
	errors "github.com/Red-Sock/trace-errors"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	pb "github.com/godverv/makosh/pkg/makosh_be"
	v1 "github.com/opencontainers/image-spec/specs-go/v1"
	"github.com/sirupsen/logrus"

	"github.com/godverv/Velez/internal/backservice/env"
	"github.com/godverv/Velez/internal/clients"
	"github.com/godverv/Velez/internal/clients/docker/dockerutils"
	"github.com/godverv/Velez/internal/config"
	"github.com/godverv/Velez/internal/service/service_manager/container_manager_v1"
	"github.com/godverv/Velez/pkg/velez_api"
)

const (
	Name     = "makosh"
	image    = "godverv/makosh:v0.0.3"
	duration = time.Second * 5

	makoshContainerAuthTokenEnvVariable = "MAKOSH_ENVIRONMENT_AUTH-TOKEN"
)

var (
	ErrRequireMakoshPortExportToRunAsDaemon = errors.New("makosh port must be exported in order to run velez as daemon")
)

type ServiceDiscoveryTask struct {
	AuthToken string
	Address   string

	dockerAPI      client.CommonAPIClient
	image          string
	portToExposeTo *string
	duration       time.Duration

	ApiClient *ApiClient
}

func newKeepAliveTask(cfg config.Config, nodeClients clients.NodeClients) (*ServiceDiscoveryTask, error) {
	serviceDiscoveryTask := &ServiceDiscoveryTask{
		dockerAPI: nodeClients.Docker(),
		image:     rtb.Coalesce(cfg.Environment.MakoshImageName, image),
		duration:  duration,
	}

	var err error
	serviceDiscoveryTask.portToExposeTo, err = getPortToExposeTo(cfg.Environment, nodeClients)
	if err != nil {
		return nil, errors.Wrap(err, "error getting port to expose for makosh")
	}

	serviceDiscoveryTask.Address, err = getTargetURL(
		cfg.Environment,
		nodeClients,
		serviceDiscoveryTask.portToExposeTo)
	if err != nil {
		return nil, errors.Wrap(err, "error getting target URL")
	}

	serviceDiscoveryTask.AuthToken = string(rtb.RandomBase64(256))

	return serviceDiscoveryTask, nil
}

func (s *ServiceDiscoveryTask) Start() error {
	ctx := context.Background()

	_, err := dockerutils.PullImage(ctx, s.dockerAPI, s.image, false)
	if err != nil {
		return errors.Wrap(err, "error pulling matreshka image")
	}

	hostConf := &container.HostConfig{}

	if s.portToExposeTo != nil {
		hostConf.PortBindings = nat.PortMap{
			"80/tcp": []nat.PortBinding{
				{
					HostPort: *s.portToExposeTo,
				},
			},
		}
	}

	createCfg := &container.Config{
		Hostname: Name,
		Image:    s.image,
		Labels: map[string]string{
			container_manager_v1.CreatedWithVelezLabel: "true",
		},
		Env: []string{makoshContainerAuthTokenEnvVariable + "=" + s.AuthToken},
	}

	cont, err := s.dockerAPI.ContainerCreate(ctx,
		createCfg,
		hostConf,
		&network.NetworkingConfig{},
		&v1.Platform{},
		Name,
	)
	if err != nil {
		return errors.Wrap(err, "error creating makosh container")
	}

	err = s.dockerAPI.ContainerStart(ctx, cont.ID, container.StartOptions{})
	if err != nil {
		return errors.Wrap(err, "error starting makosh container")
	}

	err = s.dockerAPI.NetworkConnect(ctx,
		env.VervNetwork,
		cont.ID,
		&network.EndpointSettings{Aliases: []string{Name}})
	if err != nil {
		return errors.Wrap(err, "error connecting makosh container to verv network")
	}

	return nil
}

func (s *ServiceDiscoveryTask) GetName() string {
	return Name
}

func (s *ServiceDiscoveryTask) IsAlive() bool {
	name := Name
	ctx := context.Background()

	cont, err := s.dockerAPI.ContainerInspect(ctx, name)
	if err != nil {
		if strings.Contains(err.Error(), "No such container") {
			return false
		}
		logrus.Error(errors.Wrap(err, "error getting makosh container"))
		return false
	}

	if cont.State.Status != velez_api.Smerd_running.String() {
		return false
	}

	if !rtb.Contains(cont.Config.Env, makoshContainerAuthTokenEnvVariable+"="+s.AuthToken) {
		return false
	}

	s.ApiClient, err = newApiClient(s.Address, s.AuthToken)
	if err != nil {
		logrus.Error(errors.Wrap(err, "error creating api client"))
		return false
	}

	resp, err := s.ApiClient.Version(context.Background(), &pb.Version_Request{})
	if err != nil {
		return false
	}
	if resp == nil {
		return false
	}

	return true
}

func (s *ServiceDiscoveryTask) Kill() error {
	ctx := context.Background()
	rmOpts := container.RemoveOptions{
		RemoveVolumes: true,
		Force:         true,
	}

	err := s.dockerAPI.ContainerRemove(ctx, Name, rmOpts)
	if err != nil {
		if !strings.Contains(err.Error(), "No such container") {
			return errors.Wrap(err, "error dropping result")
		}
	}

	return nil
}
