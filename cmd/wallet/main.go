package wallet

import (
	"github.com/goshathebusiness/dombu/pkg/wallet/router"
	"github.com/goshathebusiness/dombu/pkg/wallet/services"
)

func main() {
	services := services.NewServices()

	router := router.NewRouter(services)

	router.Run()
}
