package services

import (
	"gorm.io/gorm"

	"github.com/goshathebusiness/dombu/pkg/logging"
	"github.com/goshathebusiness/dombu/pkg/user-management/repository/postgres"
)

func NewServices(db *gorm.DB, logger *logging.Logger, repoManager *postgres.RepoManager) *Services {
	return &Services{
		UserSvc: NewUserService(db, logger, repoManager),
	}
}
