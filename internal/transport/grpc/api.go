package grpc

import (
	"github.com/godverv/Velez/internal/service"
	"github.com/godverv/Velez/pkg/velez_api"
)

type Api struct {
	velez_api.UnimplementedVelezAPIServer

	containerManager service.ContainerManager
	hardwareManager  service.HardwareManager

	version string
}
