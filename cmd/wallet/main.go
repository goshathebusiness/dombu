package main

import (
	"context"
	"flag"

	"github.com/goshathebusiness/dombu/pkg/config"
	"github.com/goshathebusiness/dombu/pkg/logging"
	"github.com/goshathebusiness/dombu/pkg/wallet/router"
	"github.com/goshathebusiness/dombu/pkg/wallet/services"
)

func main() {
	logger := logging.NewLogger()

	var configPath string
	flag.StringVar(&configPath, "config-path", "./cmd/wallet/config.yaml", "provides a path to configuration file with extension .yaml")
	flag.Parse()

	var cfg config.CommonConfig
	err := config.UnmarshalYAML(configPath, &cfg)
	if err != nil {
		logger.Error(context.Background(), "unmarshal config error", err)
	}

	svc := services.NewServices()

	r := router.NewRouter(svc)
	logger.Fatal(r.Run(cfg.Addr))
}
