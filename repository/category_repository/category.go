package categoryrepository

import (
	"f-commerce/model"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type CategoryRepo interface {
	CreateCategory(cat *model.Category) error
}

type categoryRepo struct {
	db  *gorm.DB
	log *zap.Logger
}

func NewCategoryRepo(db *gorm.DB, log *zap.Logger) CategoryRepo {
	return &categoryRepo{db, log}
}

func (cr *categoryRepo) CreateCategory(cat *model.Category) error {

	if err := cr.db.Create(&cat).Error; err != nil {
		return err
	}

	return nil
}
