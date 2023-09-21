package services

import "github.com/goshathebusiness/dombu/pkg/models"

type Services struct {
	UserSvc UserService
}

type UserSvc interface {
	GetUserByID(id uint) (*models.User, error)
	CreateUser(user *models.User) error
	UpdateUser(user *models.User) error
	DeleteUserByID(id uint) error
}
