package services

import (
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

func (s *UserService) GetUserByID(id uint) (*models.User, error) {
	user, err := s.repoManager.NewUserRepo(s.db).GetUserByID(id)
	if err != nil {
		return nil, handleError(s.logger, err)
	}

	return user, nil
}

func (s *UserService) CreateUser(user *models.User) error {
	err := s.repoManager.NewUserRepo(s.db).CreateUser(user)
	if err != nil {
		return handleError(s.logger, err)
	}

	return nil
}

func (s *UserService) UpdateUser(user *models.User) error {
	err := s.repoManager.NewUserRepo(s.db).UpdateUser(user)
	if err != nil {
		return handleError(s.logger, err)
	}

	return nil
}

func (s *UserService) DeleteUser(user *models.User) error {
	err := s.repoManager.NewUserRepo(s.db).DeleteUser(user)
	if err != nil {
		return handleError(s.logger, err)
	}

	return nil
}
