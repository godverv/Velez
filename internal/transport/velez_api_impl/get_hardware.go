package velez_api_impl

import (
	"context"

	"github.com/godverv/Velez/pkg/velez_api"
)

func (a *Impl) GetHardware(context.Context, *velez_api.GetHardware_Request) (*velez_api.GetHardware_Response, error) {
	return a.hardwareManager.GetHardware()
}
