package repository

import (
	"context"

	"gorm.io/gorm"

	"github.com/goshathebusiness/dombu/pkg/models"
)

type RepoManager interface {
	NewUserRepo(db *gorm.DB) UserRepository
}

type UserRepository interface {
	GetUserByID(ctx context.Context, id uint) (*models.User, error)
	CreateUser(ctx context.Context, user *models.User) error
	UpdateUser(ctx context.Context, user *models.User) error
	DeleteUser(ctx context.Context, user *models.User) error
}
