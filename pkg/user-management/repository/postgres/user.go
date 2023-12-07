package postgres

import (
	"database/sql"

	"gorm.io/gorm"

	"github.com/goshathebusiness/dombu/pkg/models"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) GetUserByID(id uint) (*models.User, error) {
	var user *models.User

	result := r.db.Find(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (r *UserRepository) CreateUser(user *models.User) error {
	result := r.db.Create(&user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *UserRepository) UpdateUser(user *models.User) error {
	result := r.db.Model(&user).Where("id", user.ID)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (r *UserRepository) DeleteUser(user *models.User) error {
	result := r.db.Delete(&user).Where("id", user.ID)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}
