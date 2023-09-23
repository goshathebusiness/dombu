package handlers

import (
	"github.com/gin-gonic/gin"

	"github.com/goshathebusiness/dombu/pkg/wallet/services"
)

func GetWalletByIDHandler(svc services.WalletService) gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func GetTransactionByIDHandler(svc services.TransactionService) gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func FetchTransactionByBalanceIDHandler(svc services.TransactionService) gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func CreateTransactionHandler(svc services.TransactionService) gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func UpdateTransactionHandler(svc services.TransactionService) gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func DeleteTransactionByIDHandler(svc services.TransactionService) gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
