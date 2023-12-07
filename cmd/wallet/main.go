package main

import (
	"context"
	"flag"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/goshathebusiness/dombu/pkg/config"
	"github.com/goshathebusiness/dombu/pkg/db"
	"github.com/goshathebusiness/dombu/pkg/logging"
	pg "github.com/goshathebusiness/dombu/pkg/wallet/repository/postgres"
	"github.com/goshathebusiness/dombu/pkg/wallet/router"
	"github.com/goshathebusiness/dombu/pkg/wallet/services"
)

func main() {
	logger := logging.NewLogger()

	var configPath string
	flag.StringVar(&configPath, "config-path", "./cmd/wallet/config.yaml", "provides a path to configuration file with extension .yaml")
	flag.Parse()

	var cfg Config
	ctx := context.Background()
	err := config.UnmarshalYAML(configPath, &cfg)
	if err != nil {
		logger.Error(ctx, "unmarshal config error", err)
	}

	database, err := gorm.Open(postgres.New(postgres.Config{DSN: cfg.DB.URL()}))
	if err != nil {
		logger.Error(ctx, "failed to connect to db", err)
	}
	err = db.AutomigrateWallet(database)
	if err != nil {
		logger.Error(ctx, "failed to migrate into db", err)
	}

	svc := services.NewServices(database, logger, pg.NewRepoManager())

	r := router.NewRouter(svc)

	logger.Fatal(r.Run(cfg.Addr))
}
