package router

import (
	"github.com/gin-gonic/gin"

	"github.com/goshathebusiness/dombu/pkg/user-management/handlers"
	"github.com/goshathebusiness/dombu/pkg/user-management/services"
)

func NewRouter(services *services.Services) *gin.Engine {
	r := gin.Default()

	r.GET("/user/{id}", handlers.GetUserByIDHandler(services.UserSvc))
	r.POST("/user", handlers.CreateUserHandler(services.UserSvc))
	r.PUT("/user", handlers.UpdateUserHandler(services.UserSvc))
	r.DELETE("/user/{id}", handlers.DeleteUserByIDHandler(services.UserSvc))
	return r
}
