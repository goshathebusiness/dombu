package repository

import (
	"gorm.io/gorm"

	"github.com/goshathebusiness/dombu/pkg/models"
)

type RepoManager interface {
	NewUserRepo(db *gorm.DB) UserRepository
}

type UserRepository interface {
	GetUserByID(id uint) (*models.User, error)
	CreateUser(user *models.User) error
	UpdateUser(user *models.User) error
	DeleteUser(user *models.User) error
}
