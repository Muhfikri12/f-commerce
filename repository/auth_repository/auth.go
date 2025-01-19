package authrepository

import (
	"f-commerce/model"
	"fmt"

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

	result := ar.db.Table("users").
		Where("email = ?", verify.Email).
		Update("status", "active")

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("email %s doesn't exist", verify.Email)
	}

	return nil
}
