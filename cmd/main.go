package main

import (
	"is-theatre/config"
	"is-theatre/internal/app"
	"is-theatre/pkg/logger/logger_logrus"
)

func main() {
	log, err := logger_logrus.NewLogrusLogger()

	cfg, err := config.NewAppConfig()
	if err != nil {
		(*log).Error("config init error:", err)
	}

	app.Run(cfg, log)
}
