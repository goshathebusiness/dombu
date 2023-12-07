package services

import (
	"context"

	"github.com/goshathebusiness/dombu/pkg/models"
)

type Services struct {
	WalletSvc      WalletService
	TransactionSvc TransactionService
}

type BalanceSvc interface {
	GetWalletByID(ctx context.Context, id uint) (*models.Wallet, error)
}

type TransactionSvc interface {
	GetTransactionByID(ctx context.Context, id uint) (*models.Transaction, error)
	CreateTransaction(ctx context.Context, transaction *models.Transaction) error
	UpdateTransaction(ctx context.Context, transaction *models.Transaction) error
	DeleteTransactionByID(ctx context.Context, id uint) error
}
