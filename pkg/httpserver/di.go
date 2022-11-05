package httpserver

import (
	"github.com/crayonwow/base_service/pkg/di"
	"go.uber.org/dig"
)

func Module() di.Module {
	return di.NewModule(
		di.NewDependency(register),
		di.NewDependency(newHTTPServer),
		di.NewDependency(appAdapter),
	)
}

type (
	appIn struct {
		dig.In

		Server *httpServer
	}

	appOut struct {
		dig.In

		Server di.Application `group:"applications"`
	}

	registratorIn struct {
		dig.In

		Registrators []Registrator `group:"registrators"`
	}
)

func appAdapter(in appIn) appOut {
	return appOut{
		Server: in.Server,
	}
}
