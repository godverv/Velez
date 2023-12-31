package container_manager_v1

import (
	"context"
	"strings"

	errors "github.com/Red-Sock/trace-errors"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/network"
	v1 "github.com/opencontainers/image-spec/specs-go/v1"

	"github.com/godverv/Velez/internal/client/docker/dockerutils/parser"
	"github.com/godverv/Velez/internal/service"
	"github.com/godverv/Velez/pkg/velez_api"
)

func (c *containerManager) LaunchSmerd(ctx context.Context, req *velez_api.CreateSmerd_Request) (*velez_api.Smerd, error) {
	image, err := c.getImage(ctx, req.ImageName)
	if err != nil {
		return nil, errors.Wrap(err, "error getting image")
	}

	if req.Name == "" {
		req.Name = strings.Split(image.Name, "/")[1]
	}

	for i := range req.Settings.Ports {
		port := c.portManager.GetPort()
		if port == nil {
			return nil, service.ErrNoPortsAvailable
		}

		req.Settings.Ports[i].Host = uint32(*port)
	}

	serviceContainer, err := c.docker.ContainerCreate(ctx,
		&container.Config{
			Image:    image.Name,
			Hostname: req.Name,
			Cmd:      parser.FromCommand(req.Command),
		},
		&container.HostConfig{
			PortBindings: parser.FromPorts(req.Settings),
			Mounts:       parser.FromBind(req.Settings),
		},
		&network.NetworkingConfig{},
		&v1.Platform{},
		req.Name,
	)
	if err != nil {
		return nil, errors.Wrap(err, "error creating container")
	}

	err = c.docker.ContainerStart(ctx, serviceContainer.ID, types.ContainerStartOptions{})
	if err != nil {
		return nil, errors.Wrap(err, "error starting container")
	}

	out := &velez_api.Smerd{
		Uuid:      serviceContainer.ID,
		ImageName: req.ImageName,
	}

	cl, err := c.docker.ContainerList(ctx, types.ContainerListOptions{
		Filters: filters.NewArgs(filters.Arg("id", serviceContainer.ID)),
	})

	if err == nil && len(cl) == 1 && len(cl[0].Names) > 0 {
		out.Name = cl[0].Names[0][1:]
	}

	return out, nil
}
