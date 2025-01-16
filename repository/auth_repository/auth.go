package authrepository

import (
	"errors"
	"f-commerce/database"
	"f-commerce/model"
	"fmt"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type AuthRepo interface {
	Login(login *model.Login) (*model.User, error)
}
type authRepo struct {
	db    *gorm.DB
	log   *zap.Logger
	redis *database.Cache
}

func NewAuthRepo(db *gorm.DB, log *zap.Logger, redis *database.Cache) AuthRepo {
	return &authRepo{db, log, redis}
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

func (ar *authRepo) VerificationEmail(verify *model.VerificationEmail) error {

	otp, err := ar.redis.Get(verify.Email)
	if err != nil {
		return err
	}

	if otp != verify.Otp {
		return errors.New("otp invalid or expired")
	}

	if err := ar.redis.Delete(verify.Email); err != nil {
		return fmt.Errorf("failed to delete otp: %v ", err)
	}

	return nil
}
