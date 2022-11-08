package config

import "is-theatre/pkg/configs"

type AppConfig struct {
	DB     *configs.DBConfig
	Server *configs.ServerConfig
}

func NewAppConfig() (*AppConfig, error) {
	dbConfig, err := configs.NewDBConfig()
	if err != nil {
		return nil, err
	}

	serverConfig, err := configs.NewServerConfig()
	if err != nil {
		return nil, err
	}

	return &AppConfig{
		DB:     dbConfig,
		Server: serverConfig,
	}, nil
}
