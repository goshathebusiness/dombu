package postgres

import (
	"context"
	"database/sql"

	"gorm.io/gorm"

	"github.com/goshathebusiness/dombu/pkg/models"
)

type TransactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *TransactionRepository {
	return &TransactionRepository{db}
}

func (r *TransactionRepository) GetTransactionByID(ctx context.Context, id uint) (*models.Transaction, error) {
	var transaction *models.Transaction

	result := r.db.WithContext(ctx).Find(&transaction, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return transaction, nil
}

func (r *TransactionRepository) FetchTransactions(ctx context.Context) ([]*models.Transaction, error) {
	var transactions []*models.Transaction

	result := r.db.WithContext(ctx).Find(&transactions)
	if result.Error != nil {
		return nil, result.Error
	}

	return transactions, nil
}

func (r *TransactionRepository) CreateTransaction(ctx context.Context, transaction *models.Transaction) error {
	result := r.db.WithContext(ctx).Create(&transaction)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *TransactionRepository) UpdateTransaction(ctx context.Context, transaction *models.Transaction) error {
	result := r.db.WithContext(ctx).Model(&transaction).Where("id", transaction.ID)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (r *TransactionRepository) DeleteTransaction(ctx context.Context, transaction *models.Transaction) error {
	result := r.db.WithContext(ctx).Delete(&transaction).Where("id", transaction.ID)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}
