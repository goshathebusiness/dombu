package services

import (
	"context"

	"gorm.io/gorm"

	"github.com/goshathebusiness/dombu/pkg/logging"
	"github.com/goshathebusiness/dombu/pkg/models"
	"github.com/goshathebusiness/dombu/pkg/wallet/repository"
)

type WalletService struct {
	db          *gorm.DB
	logger      *logging.Logger
	repoManager repository.RepoManager
}

func NewWalletService(db *gorm.DB, logger *logging.Logger, repoManager repository.RepoManager) *WalletService {
	return &WalletService{
		db:          db,
		logger:      logger,
		repoManager: repoManager,
	}
}

func (s *WalletService) GetWalletByID(ctx context.Context, id uint) (*models.Wallet, error) {
	wallet, err := s.repoManager.NewWalletRepo(s.db).GetWalletByID(ctx, id)
	if err != nil {
		return nil, handleError(s.logger, err)
	}

	return wallet, nil
}

func (s *WalletService) CreateWallet(ctx context.Context, wallet *models.Wallet) error {
	err := s.repoManager.NewWalletRepo(s.db).CreateWallet(ctx, wallet)
	if err != nil {
		return handleError(s.logger, err)
	}

	return nil
}

func (s *WalletService) UpdateWallet(ctx context.Context, wallet *models.Wallet) error {
	err := s.repoManager.NewWalletRepo(s.db).UpdateWallet(ctx, wallet)
	if err != nil {
		return handleError(s.logger, err)
	}

	return nil
}

func (s *WalletService) DeleteWallet(ctx context.Context, wallet *models.Wallet) error {
	err := s.repoManager.NewWalletRepo(s.db).DeleteWallet(ctx, wallet)
	if err != nil {
		return handleError(s.logger, err)
	}

	return nil
}
