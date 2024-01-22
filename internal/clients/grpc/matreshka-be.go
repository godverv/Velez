// Code generated by RedSock CLI. DO NOT EDIT.

package grpc

import (
	"context"

	errors "github.com/Red-Sock/trace-errors"
	pb "github.com/godverv/matreshka-be/pkg/matreshka_api"

	"github.com/godverv/Velez/internal/config"
)

func NewMatreshkaBeAPIClient(ctx context.Context, cfg config.Config) (pb.MatreshkaBeAPIClient, error) {
	connCfg, err := cfg.Resources().GRPC(config.ResourceGrpcMatreshkaBe)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't find key"+config.ResourceGrpcMatreshkaBe+" grpc connection in config")
	}

	conn, err := connect(ctx, connCfg)
	if err != nil {
		return nil, errors.Wrap(err, "error connection to "+connCfg.Module)
	}

	return pb.NewMatreshkaBeAPIClient(conn), nil
}