package repository

import (
	"f-commerce/database"
	authrepository "f-commerce/repository/auth_repository"
	userrepositoy "f-commerce/repository/user_repositoy"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Repository struct {
	Auth authrepository.AuthRepo
	User userrepositoy.UserRepo
}

func NewAllRepo(db *gorm.DB, log *zap.Logger, redis *database.Cache) *Repository {
	return &Repository{
		Auth: authrepository.NewAuthRepo(db, log, redis),
		User: userrepositoy.NewUserRepo(db, log),
	}
}
