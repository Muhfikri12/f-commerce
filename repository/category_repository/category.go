package categoryrepository

import (
	"f-commerce/model"
	"fmt"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type CategoryRepo interface {
	CreateCategory(cat *model.Category) error
	ReadCategories() (*[]model.Category, error)
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

func (cr *categoryRepo) ReadCategories() (*[]model.Category, error) {

	cat := []model.Category{}

	result := cr.db.Find(&cat)

	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("category not found")
	}

	return &cat, nil
}
