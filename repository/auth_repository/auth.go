package authrepository

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type AuthRepo interface {
}
type authRepo struct {
	db  *gorm.DB
	log *zap.Logger
}

func NewAuthRepo(db *gorm.DB, log *zap.Logger) AuthRepo {
	return &authRepo{db, log}
}
