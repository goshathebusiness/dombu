package postgres

import (
	"gorm.io/gorm"

	"github.com/goshathebusiness/dombu/pkg/user-management/repository"
)

type RepoManager struct{}

func NewRepoManager() *RepoManager {
	return &RepoManager{}
}

func (m *RepoManager) NewUserRepo(db *gorm.DB) repository.UserRepository {
	return NewUserRepository(db)
}
