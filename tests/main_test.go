//go:build integration

package tests

import (
	"context"
	"os"
	"testing"

	"github.com/docker/docker/api/types/container"
	"github.com/sirupsen/logrus"

	"github.com/godverv/Velez/internal/app"
	"github.com/godverv/Velez/internal/clients/docker/dockerutils"
	"github.com/godverv/Velez/internal/service/service_manager/container_manager_v1"
	"github.com/godverv/Velez/pkg/velez_api"
)

const (
	integrationTest = "integration_test"
)

type TestEnv struct {
	*app.App
}

var testEnv TestEnv

func TestMain(m *testing.M) {
	testEnv.App = app.New()
	testEnv.clean()

	var code int
	defer func() {
		testEnv.clean()
		os.Exit(code)
	}()

	code = m.Run()
}

func (t *TestEnv) callCreate(ctx context.Context, req *velez_api.CreateSmerd_Request) (smerd *velez_api.Smerd, err error) {
	if req.Labels == nil {
		req.Labels = map[string]string{}
	}

	req.Labels[integrationTest] = "true"
	return testEnv.GrpcApi.CreateSmerd(ctx, req)
}

func (t *TestEnv) clean() {
	ctx := context.Background()

	listReq := &velez_api.ListSmerds_Request{
		Label: map[string]string{
			integrationTest: "true",
		},
	}
	cList, err := dockerutils.ListContainers(ctx, t.Clients.Docker(), listReq)
	if err != nil {
		logrus.Fatal(err)
	}

	for _, cont := range cList {
		err = t.Clients.Docker().ContainerRemove(ctx, cont.ID,
			container.RemoveOptions{
				Force: true,
			})
		if err != nil {
			logrus.Fatal(err)
		}
	}

}

func (t *TestEnv) getExpectedLabels() map[string]string {
	return map[string]string{
		container_manager_v1.CreatedWithVelezLabel: "true",
		integrationTest: "true",
	}
}
