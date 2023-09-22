package main

import (
	"context"
	"flag"

	"github.com/goshathebusiness/dombu/pkg/config"
	"github.com/goshathebusiness/dombu/pkg/logger"
	"github.com/goshathebusiness/dombu/pkg/user-management/router"
	"github.com/goshathebusiness/dombu/pkg/user-management/services"
)

func main() {
	logger := logger.NewLogger()

	var configPath string
	flag.StringVar(&configPath, "config-path", "./cmd/user-management/config.yaml", "provides a path to configuration file with extension .yaml")
	flag.Parse()

	var cfg config.Config
	err := config.UnmarshalYAML(configPath, &cfg)
	if err != nil {
		logger.Error(context.Background(), "unmarshal config error", err)
	}

	services := services.NewServices()

	router := router.NewRouter(services)

	logger.Error(context.Background(), "router error", router.Run())
}
