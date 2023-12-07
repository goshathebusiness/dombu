package postgres

import (
	"gorm.io/gorm"

	"github.com/goshathebusiness/dombu/pkg/wallet/repository"
)

type RepoManager struct{}

func NewRepoManager() *RepoManager {
	return &RepoManager{}
}

func (m *RepoManager) NewWalletRepo(db *gorm.DB) repository.WalletRepository {
	return NewWalletRepository(db)
}

func (m *RepoManager) NewTransactionRepo(db *gorm.DB) repository.TransactionRepository {
	return NewTransactionRepository(db)
}
