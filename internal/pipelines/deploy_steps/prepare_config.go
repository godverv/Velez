package deploy_steps

import (
	"context"
	"fmt"
	"strings"

	"go.redsock.ru/rerrors"
	"go.verv.tech/matreshka"

	"github.com/godverv/Velez/internal/clients"
	"github.com/godverv/Velez/internal/domain"
	"github.com/godverv/Velez/internal/service"
	"github.com/godverv/Velez/pkg/velez_api"
)

const (
	CreatedWithVelezLabel = "CREATED_WITH_VELEZ"
	MatreshkaConfigLabel  = "MATRESHKA_CONFIG_ENABLED"
)

type prepareConfig struct {
	configService service.ConfigurationService
	portManager   clients.PortManager

	req *velez_api.CreateSmerd_Request
	dp  *domain.LaunchSmerdState

	lockedPorts []uint32
}

func PrepareVervConfig(
	configService service.ConfigurationService,
	portManager clients.PortManager,
	req *velez_api.CreateSmerd_Request,
	dp *domain.LaunchSmerdState,
) *prepareConfig {
	return &prepareConfig{
		configService: configService,
		portManager:   portManager,

		req: req,
		dp:  dp,
	}
}

func (p *prepareConfig) Do(ctx context.Context) error {
	if p.dp.Image.Config.Labels == nil {
		p.dp.Image.Config.Labels = make(map[string]string)
	}

	if p.dp.Image.Config.Labels[MatreshkaConfigLabel] == "true" {
		err := p.enrichWithMatreshkaConfig(ctx, p.req)
		if err != nil {
			return rerrors.Wrap(err, "error enriching with verv data")
		}
	} else {
		p.req.Labels[MatreshkaConfigLabel] = "false"
		// TODO заиспользовать ручку из VERV-75 для получения конфигурации ресурса
	}

	if p.req.UseImagePorts {
		err := p.getPortsFromImage()
		if err != nil {
			return rerrors.Wrap(err, "error locking ports")
		}
	}

	err := p.lockPorts()
	if err != nil {
		return rerrors.Wrap(err, "error locking ports")
	}

	p.req.Env[matreshka.VervName] = p.req.GetName()
	p.req.Labels[CreatedWithVelezLabel] = "true"

	return nil
}

func (p *prepareConfig) Rollback(_ context.Context) error {
	p.portManager.UnlockPorts(p.lockedPorts)
	return nil
}

func (p *prepareConfig) enrichWithMatreshkaConfig(ctx context.Context, req *velez_api.CreateSmerd_Request) error {
	if req.IgnoreConfig {
		req.Labels[MatreshkaConfigLabel] = "false"
		return nil
	}

	envs, err := p.configService.GetEnvFromApi(ctx, req.GetName())
	if err != nil {
		return rerrors.Wrap(err, "error getting matreshka config from matreshka api")
	}
	serviceName := strings.ToUpper(req.GetName())
	for _, e := range envs {
		if len(e.InnerNodes) == 0 && e.Value != nil {
			req.Env[serviceName+"_"+e.Name] = fmt.Sprint(e.Value)
		}
	}

	return nil
}

func (p *prepareConfig) getPortsFromImage() error {
	portsInReq := map[uint32]*velez_api.Port{}
	for _, port := range p.req.Settings.Ports {
		portsInReq[port.ServicePortNumber] = port
	}

	for port := range p.dp.Image.Config.ExposedPorts {
		portVal := uint32(port.Int())
		_, ok := portsInReq[portVal]
		if ok {
			continue
		}

		portBind := &velez_api.Port{
			ServicePortNumber: portVal,
			Protocol:          velez_api.Port_Protocol(velez_api.Port_Protocol_value[port.Proto()]),
		}
		p.req.Settings.Ports = append(p.req.Settings.Ports, portBind)
		portsInReq[portVal] = portBind
	}

	return nil
}

func (p *prepareConfig) lockPorts() (err error) {
	p.lockedPorts = make([]uint32, 0, len(p.req.Settings.Ports))

	for _, imagePort := range p.req.Settings.Ports {
		if imagePort.ExposedTo == nil {
			var port uint32
			port, err = p.portManager.GetPort()
			imagePort.ExposedTo = &port
		} else {
			err = p.portManager.LockPort(*imagePort.ExposedTo)
		}
		if err != nil {
			err = rerrors.Wrap(err, "error locking host port")
			return
		}

		p.lockedPorts = append(p.lockedPorts, *imagePort.ExposedTo)
	}

	return nil
}
