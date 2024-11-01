// Code generated by RedSock CLI. DO NOT EDIT.

package app

import (
	errors "github.com/Red-Sock/trace-errors"
	"github.com/godverv/Velez/internal/transport"
)

func (a *App) InitServers() (err error) {
	a.Server, err = transport.NewServerManager(a.Ctx, ":53890")
	if err != nil {
		return errors.Wrap(err, "error during server initialization on port: 53890")
	}

	return nil
}
