package main

import (
	"github.com/goshathebusiness/dombu/pkg/user-management/router"
	"github.com/goshathebusiness/dombu/pkg/user-management/services"
)

func main() {
	services := services.NewServices()

	router := router.NewRouter(services)

	router.Run()
}
