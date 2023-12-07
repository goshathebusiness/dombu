package main

import (
	"github.com/goshathebusiness/dombu/pkg/config"
	"github.com/goshathebusiness/dombu/pkg/db"
)

type Config struct {
	config.CommonConfig
	DB db.Config `yaml:"DB"`
}
