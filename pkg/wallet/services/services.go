package services

import (
	"gorm.io/gorm"

	"github.com/goshathebusiness/dombu/pkg/logging"
	"github.com/goshathebusiness/dombu/pkg/wallet/repository/postgres"
)

func NewServices(db *gorm.DB, logger *logging.Logger, repoManager *postgres.RepoManager) *Services {
	return &Services{
		WalletSvc:      NewWalletService(db, logger, repoManager),
		TransactionSvc: NewTransactionService(db, logger, repoManager),
	}
}
