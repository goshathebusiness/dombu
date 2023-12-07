package repository

import (
	"context"

	"gorm.io/gorm"

	"github.com/goshathebusiness/dombu/pkg/models"
)

type RepoManager interface {
	NewWalletRepo(db *gorm.DB) WalletRepository
	NewTransactionRepo(db *gorm.DB) TransactionRepository
}

type WalletRepository interface {
	GetWalletByID(ctx context.Context, id uint) (*models.Wallet, error)
	CreateWallet(ctx context.Context, wallet *models.Wallet) error
	UpdateWallet(ctx context.Context, wallet *models.Wallet) error
	DeleteWallet(ctx context.Context, wallet *models.Wallet) error
}

type TransactionRepository interface {
	GetTransactionByID(ctx context.Context, id uint) (*models.Transaction, error)
	FetchTransactions(ctx context.Context) ([]*models.Transaction, error)
	CreateTransaction(ctx context.Context, transaction *models.Transaction) error
	UpdateTransaction(ctx context.Context, transaction *models.Transaction) error
	DeleteTransaction(ctx context.Context, transaction *models.Transaction) error
}
