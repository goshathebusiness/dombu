package services

import (
	"context"

	"gorm.io/gorm"

	"github.com/goshathebusiness/dombu/pkg/wallet/repository"

	"github.com/goshathebusiness/dombu/pkg/logging"
	"github.com/goshathebusiness/dombu/pkg/models"
)

type TransactionService struct {
	db          *gorm.DB
	logger      *logging.Logger
	repoManager repository.RepoManager
}

func NewTransactionService(db *gorm.DB, logger *logging.Logger, repoManager repository.RepoManager) *TransactionService {
	return &TransactionService{
		db:          db,
		logger:      logger,
		repoManager: repoManager,
	}
}

func (s *TransactionService) GetTransactionByID(ctx context.Context, id uint) (*models.Transaction, error) {
	transaction, err := s.repoManager.NewTransactionRepo(s.db).GetTransactionByID(ctx, id)
	if err != nil {
		return nil, handleError(s.logger, err)
	}

	return transaction, nil
}

func (s *TransactionService) FetchTransactions(ctx context.Context) ([]*models.Transaction, error) {
	transactions, err := s.repoManager.NewTransactionRepo(s.db).FetchTransactions(ctx)
	if err != nil {
		return nil, handleError(s.logger, err)
	}

	return transactions, nil
}

func (s *TransactionService) CreateTransaction(ctx context.Context, transaction *models.Transaction) error {
	err := s.repoManager.NewTransactionRepo(s.db).CreateTransaction(ctx, transaction)
	if err != nil {
		return handleError(s.logger, err)
	}

	return nil
}

func (s *TransactionService) UpdateTransaction(ctx context.Context, transaction *models.Transaction) error {
	err := s.repoManager.NewTransactionRepo(s.db).UpdateTransaction(ctx, transaction)
	if err != nil {
		return handleError(s.logger, err)
	}

	return nil
}

func (s *TransactionService) DeleteTransaction(ctx context.Context, transaction *models.Transaction) error {
	err := s.repoManager.NewTransactionRepo(s.db).DeleteTransaction(ctx, transaction)
	if err != nil {
		return handleError(s.logger, err)
	}

	return nil
}
