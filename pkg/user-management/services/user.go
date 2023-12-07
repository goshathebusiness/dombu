package services

import (
	"context"

	"gorm.io/gorm"

	"github.com/goshathebusiness/dombu/pkg/logging"
	"github.com/goshathebusiness/dombu/pkg/models"
	"github.com/goshathebusiness/dombu/pkg/user-management/repository"
)

type UserService struct {
	db          *gorm.DB
	logger      *logging.Logger
	repoManager repository.RepoManager
}

func NewUserService(db *gorm.DB, logger *logging.Logger, repoManager repository.RepoManager) *UserService {
	return &UserService{
		db:          db,
		logger:      logger,
		repoManager: repoManager,
	}
}

func (s *UserService) GetUserByID(ctx context.Context, id uint) (*models.User, error) {
	user, err := s.repoManager.NewUserRepo(s.db).GetUserByID(ctx, id)
	if err != nil {
		return nil, handleError(s.logger, err)
	}

	return user, nil
}

func (s *UserService) CreateUser(ctx context.Context, user *models.User) error {
	err := s.repoManager.NewUserRepo(s.db).CreateUser(ctx, user)
	if err != nil {
		return handleError(s.logger, err)
	}

	return nil
}

func (s *UserService) UpdateUser(ctx context.Context, user *models.User) error {
	err := s.repoManager.NewUserRepo(s.db).UpdateUser(ctx, user)
	if err != nil {
		return handleError(s.logger, err)
	}

	return nil
}

func (s *UserService) DeleteUser(ctx context.Context, user *models.User) error {
	err := s.repoManager.NewUserRepo(s.db).DeleteUser(ctx, user)
	if err != nil {
		return handleError(s.logger, err)
	}

	return nil
}
