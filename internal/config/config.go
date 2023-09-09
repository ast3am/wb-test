package config

import (
	"github.com/example/pkg/logging"
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	ListenPort string     `yaml:"listen_port"`
	SqlConfig  SqlConfig  `yaml:"sql_config"`
	NatsConfig NatsConfig `yaml:"nats_config"`
	LogLevel   string     `yaml:"log_level"`
}

type SqlConfig struct {
	UsernameDB string `yaml:"username_db"`
	PasswordDB string `yaml:"password_db"`
	HostDB     string `yaml:"host_db"`
	PortDB     string `yaml:"port_db"`
	DBName     string `yaml:"db_name"`
}

type NatsConfig struct {
	ChannelName string `yaml:"channel_name"`
	ClusterID   string `yaml:"cluster_id"`
	ClientID    string `yaml:"client_id"`
	Host        string `yaml:"host"`
	Port        string `yaml:"port"`
}

var cfg *Config

func GetConfig(path string) *Config {
	logger := logging.GetLogger("")
	logger.Info().Msg("read config")
	cfg = &Config{}
	err := cleanenv.ReadConfig(path, cfg)
	if err != nil {
		logger.Fatal("", err)
	}
	return cfg
}
