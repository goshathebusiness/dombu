package handlers

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/goshathebusiness/dombu/pkg/models"
	"github.com/goshathebusiness/dombu/pkg/wallet/services"
)

func GetTransactionByIDHandler(svc *services.TransactionService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseUint(c.Param("id"), 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
		}

		transaction, err := svc.GetTransactionByID(c, uint(id))
		if err != nil {
			if errors.Is(err, services.ErrNotFound) {
				c.JSON(http.StatusNotFound, nil)
			}
			c.JSON(http.StatusInternalServerError, err)
		}

		c.JSON(http.StatusOK, transaction)
	}
}

func FetchTransactionsHandler(svc *services.TransactionService) gin.HandlerFunc {
	return func(c *gin.Context) {
		transactions, err := svc.FetchTransactions(c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
		}

		c.JSON(http.StatusOK, transactions)
	}
}

func CreateTransactionHandler(svc *services.TransactionService) gin.HandlerFunc {
	return func(c *gin.Context) {
		jsonData, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
		}

		var transaction *models.Transaction
		err = json.Unmarshal(jsonData, transaction)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
		}

		err = svc.CreateTransaction(c, transaction)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
		}

		c.JSON(http.StatusOK, transaction)
	}
}

func UpdateTransactionHandler(svc *services.TransactionService) gin.HandlerFunc {
	return func(c *gin.Context) {
		jsonData, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
		}

		var transaction *models.Transaction
		err = json.Unmarshal(jsonData, transaction)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
		}

		err = svc.UpdateTransaction(c, transaction)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
		}

		c.JSON(http.StatusOK, transaction)
	}
}

func DeleteTransactionHandler(svc *services.TransactionService) gin.HandlerFunc {
	return func(c *gin.Context) {
		jsonData, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
		}

		var transaction *models.Transaction
		err = json.Unmarshal(jsonData, transaction)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
		}

		err = svc.DeleteTransaction(c, transaction)
		if err != nil {
			if errors.Is(err, services.ErrNotFound) {
				c.JSON(http.StatusNotFound, nil)
			}
			c.JSON(http.StatusInternalServerError, err)
		}

		c.JSON(http.StatusOK, nil)
	}
}
