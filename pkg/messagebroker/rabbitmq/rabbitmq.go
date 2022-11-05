package rabbitmq

import (
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

const url = "amqp://%s:%s@%s:%s/"

func newRabbit(cfg Config) (*amqp.Connection, error) {
	return amqp.Dial(fmt.Sprintf(url, cfg.User, cfg.Pass, cfg.Host, cfg.Port))
}
