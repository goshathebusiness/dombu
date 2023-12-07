package router

import (
	"github.com/gin-gonic/gin"

	"github.com/goshathebusiness/dombu/pkg/wallet/handlers"
	"github.com/goshathebusiness/dombu/pkg/wallet/services"
)

func NewRouter(services *services.Services) *gin.Engine {
	r := gin.Default()

	r.GET("/wallets/:id", handlers.GetWalletByIDHandler(services.WalletSvc))

	r.GET("/transactions/:id", handlers.GetTransactionByIDHandler(services.TransactionSvc))
	r.GET("/transactions", handlers.FetchTransactionByBalanceIDHandler(services.TransactionSvc))
	r.POST("/transactions", handlers.CreateTransactionHandler(services.TransactionSvc))
	r.PUT("/transactions", handlers.UpdateTransactionHandler(services.TransactionSvc))
	r.DELETE("/transactions/:id", handlers.DeleteTransactionByIDHandler(services.TransactionSvc))

	return r
}
