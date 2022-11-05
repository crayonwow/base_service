package user

import (
	"github.com/crayonwow/base_service/pkg/di"
	"github.com/crayonwow/base_service/pkg/httpserver"

	"go.uber.org/dig"
)

func Module() di.Module {
	return di.NewModule(
		di.NewDependency(newCache),
		di.NewDependency(newRepo),
		di.NewDependency(newService),
		di.NewDependency(newController),
		di.NewDependency(registratorAdapter),
	)
}

type (
	registratorIn struct {
		dig.In

		Controller *controller
	}

	registratorOut struct {
		dig.Out

		Controller httpserver.Registrator `group:"registrators"`
	}
)

func registratorAdapter(in registratorIn) registratorOut {
	return registratorOut{
		Controller: in.Controller,
	}
}
