package portainer

import (
	"context"
	"time"

	errors "github.com/Red-Sock/trace-errors"

	"github.com/godverv/Velez/internal/service"
	"github.com/godverv/Velez/pkg/velez_api"
)

const (
	portainerName       = "portainer-agent"
	portainerAgentImage = "portainer/agent:2.19.4"

	portainerDuration = time.Second * 5
)

type Portainer struct {
	cm service.ContainerManager

	duration time.Duration
}

func New(cm service.ContainerManager) *Portainer {
	w := &Portainer{
		cm:       cm,
		duration: portainerDuration,
	}

	return w
}

func (b *Portainer) Start() error {
	ctx := context.Background()

	name := portainerName
	req := &velez_api.CreateSmerd_Request{
		Name:      name,
		ImageName: portainerAgentImage,
		Settings: &velez_api.Container_Settings{
			Ports: []*velez_api.PortBindings{
				{
					Container: 9001,
				},
			},
			Mounts: []*velez_api.MountBindings{
				{
					Host:      "/var/run/docker.sock",
					Container: "/var/run/docker.sock",
				},
				{
					Host:      "/var/lib/docker/volumes",
					Container: "/var/lib/docker/volumes",
				},
			},
		},
	}
	_, err := b.cm.LaunchSmerd(ctx, req)
	if err != nil {
		return errors.Wrap(err, "error launching portainer's smerd")
	}

	return nil
}

func (b *Portainer) GetName() string {
	return portainerName
}

func (b *Portainer) GetDuration() time.Duration {
	return b.duration
}

func (b *Portainer) IsAlive() (bool, error) {
	name := portainerName

	smerds, err := b.cm.ListSmerds(context.Background(), &velez_api.ListSmerds_Request{Name: &name})
	if err != nil {
		return false, errors.Wrap(err, "error listing smerds with name "+name)
	}

	for _, smerd := range smerds.Smerds {
		if smerd.Name == name && smerd.Status == velez_api.Smerd_running {
			return true, nil
		}
	}

	return false, nil
}

func (b *Portainer) Kill() error {
	dropRes, err := b.cm.DropSmerds(context.Background(), &velez_api.DropSmerd_Request{
		Name: []string{portainerName},
	})
	if err != nil {
		return errors.Wrap(err, "error dropping result")
	}

	if len(dropRes.Failed) != 0 {
		return errors.New(dropRes.Failed[0].Cause)
	}

	return nil
}
