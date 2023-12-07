package services

import (
	"context"

	"github.com/goshathebusiness/dombu/pkg/models"
)

type Services struct {
	UserSvc *UserService
}

type UserSvc interface {
	GetUserByID(ctx context.Context, id uint) (*models.User, error)
	CreateUser(ctx context.Context, user *models.User) error
	UpdateUser(ctx context.Context, user *models.User) error
	DeleteUser(ctx context.Context, user *models.User) error
}
