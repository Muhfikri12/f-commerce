package categoryservice

import (
	"f-commerce/model"
	"f-commerce/repository"

	"go.uber.org/zap"
)

type CategoryService interface {
	CreateCategory(cat *model.Category) error
	ReadCategories() (*[]model.Category, error)
}

type categoryService struct {
	repo *repository.Repository
	log  *zap.Logger
}

func NewCategoryService(repo *repository.Repository, log *zap.Logger) CategoryService {
	return &categoryService{repo, log}
}

func (cs *categoryService) CreateCategory(cat *model.Category) error {
	return cs.repo.Category.CreateCategory(cat)
}

func (cs *categoryService) ReadCategories() (*[]model.Category, error) {
	return cs.repo.Category.ReadCategories()
}
