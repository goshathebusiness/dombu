package handlers

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/goshathebusiness/dombu/pkg/models"
	"github.com/goshathebusiness/dombu/pkg/user-management/services"
)

func GetUserByIDHandler(svc *services.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseUint(c.Param("id"), 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
		}

		user, err := svc.GetUserByID(c, uint(id))
		if err != nil {
			if errors.Is(err, services.ErrNotFound) {
				c.JSON(http.StatusNotFound, nil)
			}
			c.JSON(http.StatusInternalServerError, err)
		}

		c.JSON(http.StatusOK, user)
	}
}

func CreateUserHandler(svc *services.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		jsonData, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
		}

		var user *models.User
		err = json.Unmarshal(jsonData, user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
		}

		err = svc.CreateUser(c, user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
		}

		c.JSON(http.StatusOK, user)
	}
}

func UpdateUserHandler(svc *services.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		jsonData, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
		}

		var user *models.User
		err = json.Unmarshal(jsonData, user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
		}

		err = svc.UpdateUser(c, user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
		}

		c.JSON(http.StatusOK, user)
	}
}

func DeleteUserHandler(svc *services.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		jsonData, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
		}

		var user *models.User
		err = json.Unmarshal(jsonData, user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
		}

		err = svc.DeleteUser(c, user)
		if err != nil {
			if errors.Is(err, services.ErrNotFound) {
				c.JSON(http.StatusNotFound, nil)
			}
			c.JSON(http.StatusInternalServerError, err)
		}

		c.JSON(http.StatusOK, nil)
	}
}
