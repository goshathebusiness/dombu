package postgres

import (
	"context"
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

func (r *UserRepository) GetUserByID(ctx context.Context, id uint) (*models.User, error) {
	var user *models.User

	result := r.db.WithContext(ctx).Find(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (r *UserRepository) CreateUser(ctx context.Context, user *models.User) error {
	result := r.db.WithContext(ctx).Create(&user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *UserRepository) UpdateUser(ctx context.Context, user *models.User) error {
	result := r.db.WithContext(ctx).Model(&user).Where("id", user.ID)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (r *UserRepository) DeleteUser(ctx context.Context, user *models.User) error {
	result := r.db.WithContext(ctx).Delete(&user).Where("id", user.ID)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}
