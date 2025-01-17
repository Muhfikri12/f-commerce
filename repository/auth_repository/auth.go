package authrepository

import (
	"f-commerce/model"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type AuthRepo interface {
	VerificationEmail(verify *model.VerificationEmail) error
}
type authRepo struct {
	db  *gorm.DB
	log *zap.Logger
}

func NewAuthRepo(db *gorm.DB, log *zap.Logger) AuthRepo {
	return &authRepo{db, log}
}

func (ar *authRepo) VerificationEmail(verify *model.VerificationEmail) error {

	if err := ar.db.Table("users").
		Where("email = ?", verify.Email).
		Update("status", "active").Error; err != nil {
		return err
	}

	return nil
}
