package service

import (
	"context"

	"github.com/godverv/Velez/pkg/velez_api"
)

type ContainerManager interface {
	LaunchSmerd(ctx context.Context, req *velez_api.CreateSmerd_Request) (*velez_api.Smerd, error)
	ListSmerds(ctx context.Context, req *velez_api.ListSmerds_Request) (*velez_api.ListSmerds_Response, error)
}
