package rabbitmq

type Config struct {
	User string `yaml:"user"`
	Pass string `yaml:"pass"`
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}
