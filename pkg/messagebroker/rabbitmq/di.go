package rabbitmq

import (
	"github.com/crayonwow/base_service/pkg/di"
)

func Module() di.Module {
	return di.NewModule(di.NewDependency(newRabbit))
}
