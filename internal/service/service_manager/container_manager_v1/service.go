package container_manager_v1

import (
	"github.com/docker/docker/client"
	"github.com/godverv/matreshka-be/pkg/matreshka_api"

	"github.com/godverv/Velez/internal/config"
	"github.com/godverv/Velez/internal/service/service_manager/container_manager_v1/config_manager"
	"github.com/godverv/Velez/internal/service/service_manager/container_manager_v1/port_manager"
	"github.com/godverv/Velez/internal/service/service_manager/container_manager_v1/resource_manager"
)

type ContainerManager struct {
	docker client.CommonAPIClient

	configManager     *config_manager.Configurator
	resourceManager   *resource_manager.ResourceManager
	containerLauncher ContainerStarter

	matreshkaURL string
}

func NewContainerManager(
	cfg config.Config,
	docker client.CommonAPIClient,
	configClient matreshka_api.MatreshkaBeAPIClient,
	portManager *port_manager.PortManager,
) *ContainerManager {
	cm := config_manager.New(configClient, docker)
	rm := resource_manager.New(docker)
	return &ContainerManager{
		docker: docker,

		resourceManager: rm,
		configManager:   cm,

		containerLauncher: ContainerStarter{
			docker:          docker,
			configManager:   cm,
			resourceManager: rm,
			isNodeModeOn:    cfg.GetEnvironment().NodeMode,
			portManager:     portManager,
		},
	}
}
