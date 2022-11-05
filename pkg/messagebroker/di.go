package messagebroker

import (
	"github.com/crayonwow/base_service/pkg/di"
	"github.com/crayonwow/base_service/pkg/messagebroker/rabbitmq"
)

func Module() di.Module {
	return di.NewModule().Append(rabbitmq.Module())
}
