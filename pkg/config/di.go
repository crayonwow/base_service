package config

import (
	"github.com/crayonwow/base_service/pkg/di"

	"go.uber.org/dig"
)

func Module() di.Module {
	return di.NewModule(
		di.NewDependency(NewConfig),
		di.NewDependency(LoggerConfig),
		di.NewDependency(configPath, dig.Name("config_path")),
		di.NewDependency(MySQLConfig),
		di.NewDependency(RabbitmqConfig),
		di.NewDependency(RedisConfig),
	)
}

type (
	configIn struct {
		dig.In

		ConfigPath string `name:"config_path"`
	}
)
