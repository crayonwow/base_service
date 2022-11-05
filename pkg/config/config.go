package config

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/crayonwow/base_service/pkg/db/mysql"
	"github.com/crayonwow/base_service/pkg/db/redis"
	"github.com/crayonwow/base_service/pkg/logger"
	"github.com/crayonwow/base_service/pkg/messagebroker/rabbitmq"

	"gopkg.in/yaml.v3"
)

type (
	Config struct {
		RawConfig []byte          `yaml:"-"`
		Logger    logger.Config   `yaml:"logger"`
		MySQL     mysql.Config    `yaml:"mysql"`
		Rabbitmq  rabbitmq.Config `yaml:"rabbitmq"`
		Redis     redis.Config    `yaml:"redis"`
	}
)

func configPath() string {
	path := flag.String("configPath", "config/config.yaml", "path to the config file")
	flag.Parse()
	return *path
}

func NewConfig(params configIn) (*Config, error) {
	if params.ConfigPath == "" {
		return nil, errors.New("empty path")
	}
	b, err := os.ReadFile(params.ConfigPath)
	if err != nil {
		return nil, fmt.Errorf("read file: %w", err)
	}

	c := &Config{
		RawConfig: b,
	}
	return c, yaml.Unmarshal(b, c)
}

func LoggerConfig(config *Config) *logger.Config {
	return &config.Logger
}

func MySQLConfig(c *Config) mysql.Config {
	return c.MySQL
}

func RabbitmqConfig(c *Config) rabbitmq.Config {
	return c.Rabbitmq
}
func RedisConfig(c *Config) redis.Config {
	return c.Redis
}
