package service_discovery

import (
	"context"
	"os"

	"github.com/godverv/makosh/pkg/makosh_be"
	vervResolver "github.com/godverv/makosh/pkg/resolver"
	"github.com/godverv/makosh/pkg/resolver/makosh_resolver"
	errors "go.redsock.ru/rerrors"
	"go.verv.tech/matreshka"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/godverv/Velez/internal/clients/makosh"
)

type ServiceDiscovery struct {
	Sd *vervResolver.ServiceDiscovery
	makosh_be.MakoshBeAPIClient
}

func SetupServiceDiscovery(addr string, token string, msd matreshka.ServiceDiscovery) (sd ServiceDiscovery, err error) {
	_ = os.Setenv(makosh_resolver.MakoshURL, addr)
	_ = os.Setenv(makosh_resolver.MakoshSecret, token)

	sd.Sd, err = vervResolver.Init()
	if err != nil {
		return sd, errors.Wrap(err, "error initializing verv-service-discovery")
	}
	// TODO VERV-126
	for _, override := range msd.Overrides {
		ec := makosh_resolver.EndpointsContainer{}
		err = ec.SetAddrs(override.Urls...)
		if err != nil {
			return sd, errors.Wrap(err, "error setting endpoints overrides")
		}

		// TODO VERV-125
		customResolver := &makosh_resolver.StaticResolver{
			EndpointsContainer: ec,
		}

		err = sd.Sd.SetCustomResolver(customResolver, msd.Overrides[0].ServiceName)
		if err != nil {
			return sd, errors.Wrap(err, "error setting custom resolver")
		}
	}

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	sd.MakoshBeAPIClient, err = makosh.New(token, opts...)
	if err != nil {
		return sd, errors.Wrap(err, "error creating makosh api client")
	}

	_, err = sd.Version(context.Background(), &makosh_be.Version_Request{})
	if err != nil {
		return sd, errors.Wrap(err, "error pinging service discovery")
	}

	return sd, nil
}

func (s *ServiceDiscovery) GetAddrs(vervName string) []string {
	matreshkaResolverPtr, err := s.Sd.GetResolver(vervName)
	if err != nil {
		return nil
	}

	return matreshkaResolverPtr.Load().GetAddrs()
}
