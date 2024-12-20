package velez_api_impl

import (
	"context"

	"github.com/godverv/Velez/pkg/velez_api"
)

func (a *Impl) DropSmerd(ctx context.Context, req *velez_api.DropSmerd_Request) (*velez_api.DropSmerd_Response, error) {
	return a.srv.DropSmerds(ctx, req)
}