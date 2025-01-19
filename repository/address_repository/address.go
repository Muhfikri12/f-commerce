package addressrepository

import (
	"f-commerce/model"
	"fmt"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type AddressRepo interface {
	CreateAddress(add *model.Address) error
	FindAddressByUserID(id int) (*model.Address, error)
	FindAddressByID(id int) (*model.Address, error)
	UpdateAddress(id int, addr *model.Address) error
}

type addressRepo struct {
	db  *gorm.DB
	log *zap.Logger
}

func NewAddressRepo(db *gorm.DB, log *zap.Logger) AddressRepo {
	return &addressRepo{db, log}
}

func (ar *addressRepo) CreateAddress(add *model.Address) error {
	if err := ar.db.Create(&add).Error; err != nil {
		return err
	}

	return nil
}

func (ar *addressRepo) FindAddressByUserID(id int) (*model.Address, error) {

	addr := model.Address{}

	result := ar.db.Where("user_id = ?", id).Find(&addr)

	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("user with id %d not found", id)
	}

	return &addr, nil
}

func (ar *addressRepo) UpdateAddress(id int, addr *model.Address) error {

	result := ar.db.Where("id = ?", id).Updates(&addr)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("address with id %d not found", id)
	}
	return nil
}

func (ar *addressRepo) FindAddressByID(id int) (*model.Address, error) {

	addr := model.Address{}
	result := ar.db.Where("id = ?", id).First(&addr)

	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("address with id %d not found", id)
	}

	return &addr, nil
}
