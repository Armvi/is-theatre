package configs

import (
	"os"
)

const (
	SERVER_PORT = "OC_SERVER_PORT"
)

type ServerConfig struct {
	Port string
}

func NewServerConfig() (*ServerConfig, error) {
	lConfig := ServerConfig{
		Port: os.Getenv(SERVER_PORT),
	}

	lEmptyConfig := ServerConfig{}

	if lConfig == lEmptyConfig {
		return nil, ErrEmptyConfig
	}

	return &lConfig, nil
}
