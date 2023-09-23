package router

import (
	"github.com/gin-gonic/gin"

	"github.com/goshathebusiness/dombu/pkg/wallet/handlers"
	"github.com/goshathebusiness/dombu/pkg/wallet/services"
)

func NewRouter(services *services.Services) *gin.Engine {
	r := gin.Default()

	r.GET("/wallet/:id", handlers.GetWalletByIDHandler(services.WalletSvc))

	r.GET("/transaction/:id", handlers.GetTransactionByIDHandler(services.TransactionSvc))
	r.GET("/transaction", handlers.FetchTransactionByBalanceIDHandler(services.TransactionSvc))
	r.POST("/transaction", handlers.CreateTransactionHandler(services.TransactionSvc))
	r.PUT("/transaction", handlers.UpdateTransactionHandler(services.TransactionSvc))
	r.DELETE("/transaction/:id", handlers.DeleteTransactionByIDHandler(services.TransactionSvc))

	return r
}
