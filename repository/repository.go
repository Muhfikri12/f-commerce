package repository

import (
	addressrepository "f-commerce/repository/address_repository"
	authrepository "f-commerce/repository/auth_repository"
	categoryrepository "f-commerce/repository/category_repository"
	userrepositoy "f-commerce/repository/user_repositoy"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Repository struct {
	Auth     authrepository.AuthRepo
	User     userrepositoy.UserRepo
	Address  addressrepository.AddressRepo
	Category categoryrepository.CategoryRepo
}

func NewAllRepo(db *gorm.DB, log *zap.Logger) *Repository {
	return &Repository{
		Auth:     authrepository.NewAuthRepo(db, log),
		User:     userrepositoy.NewUserRepo(db, log),
		Address:  addressrepository.NewAddressRepo(db, log),
		Category: categoryrepository.NewCategoryRepo(db, log),
	}
}
