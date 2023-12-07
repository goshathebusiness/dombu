package postgres

import (
	"context"
	"database/sql"

	"gorm.io/gorm"

	"github.com/goshathebusiness/dombu/pkg/models"
)

type WalletRepository struct {
	db *gorm.DB
}

func NewWalletRepository(db *gorm.DB) *WalletRepository {
	return &WalletRepository{db}
}

func (r *WalletRepository) GetWalletByID(ctx context.Context, id uint) (*models.Wallet, error) {
	var wallet *models.Wallet

	result := r.db.WithContext(ctx).Find(&wallet, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return wallet, nil
}

func (r *WalletRepository) CreateWallet(ctx context.Context, wallet *models.Wallet) error {
	result := r.db.WithContext(ctx).Create(&wallet)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *WalletRepository) UpdateWallet(ctx context.Context, wallet *models.Wallet) error {
	result := r.db.WithContext(ctx).Model(&wallet).Where("id", wallet.ID)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (r *WalletRepository) DeleteWallet(ctx context.Context, wallet *models.Wallet) error {
	result := r.db.WithContext(ctx).Delete(&wallet).Where("id", wallet.ID)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}
