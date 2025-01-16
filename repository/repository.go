package repository

import (
	authrepository "f-commerce/repository/auth_repository"
	userrepositoy "f-commerce/repository/user_repositoy"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Repository struct {
	Auth authrepository.AuthRepo
	User userrepositoy.UserRepo
}

func NewAllRepo(db *gorm.DB, log *zap.Logger) *Repository {
	return &Repository{
		Auth: authrepository.NewAuthRepo(db, log),
		User: userrepositoy.NewUserRepo(db, log),
	}
}
