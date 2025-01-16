package repository

import (
	authrepository "finance/repository/auth_repository"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Repository struct {
	Auth authrepository.AuthRepo
}

func NewAllRepo(db *gorm.DB, log *zap.Logger) *Repository {
	return &Repository{
		Auth: authrepository.NewAuthRepo(db, log),
	}
}
