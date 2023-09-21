package services

import "github.com/goshathebusiness/dombu/pkg/models"

type Services struct {
	WalletSvc      WalletService
	TransactionSvc TransactionService
}

type BalanceSvc interface {
	GetWalletByID(id uint) (*models.Wallet, error)
	CreateUser(user *models.User) error
	UpdateUser(user *models.User) error
	DeleteUserByID(id uint) error
}
