package authrepository

import (
	"errors"
	"finance/model"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type AuthRepo interface {
	Login(login *model.Login) (*model.User, error)
}
type authRepo struct {
	db  *gorm.DB
	log *zap.Logger
}

func NewAuthRepo(db *gorm.DB, log *zap.Logger) AuthRepo {
	return &authRepo{db, log}
}

func (ar *authRepo) Login(login *model.Login) (*model.User, error) {

	user := model.User{}
	if err := ar.db.Table("users").
		Where("email = ? OR username = ?", login.Input, login.Input).
		First(&user).Error; err != nil {
		ar.log.Error("Login error", zap.Error(err))
		return nil, errors.New("invalid email or username")
	}

	return &user, nil
}
