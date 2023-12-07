package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/goshathebusiness/dombu/pkg/wallet/services"
)

func GetWalletByIDHandler(svc *services.WalletService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseUint(c.Param("id"), 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
		}

		wallet, err := svc.GetWalletByID(c, uint(id))
		if err != nil {
			if errors.Is(err, services.ErrNotFound) {
				c.JSON(http.StatusNotFound, nil)
			}
			c.JSON(http.StatusInternalServerError, err)
		}

		c.JSON(http.StatusOK, wallet)
	}
}
