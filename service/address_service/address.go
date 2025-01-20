package addressservice

import (
	"f-commerce/helper"
	"f-commerce/model"
	"f-commerce/repository"

	"go.uber.org/zap"
)

type AddressService interface {
	CreateAddress(token string, add *model.Address) error
	FindAddressByUserID(token string) (*model.Address, error)
	UpdateAddress(id int, addr *model.Address) error
	FindAddressByID(id int) (*model.Address, error)
	DeleteAddress(id int) error
}

type addressService struct {
	repo *repository.Repository
	log  *zap.Logger
	jwt  *helper.Jwt
}

func NewAddressService(repo *repository.Repository, log *zap.Logger, jwt *helper.Jwt) AddressService {
	return &addressService{repo, log, jwt}
}

func (as *addressService) CreateAddress(token string, add *model.Address) error {

	id, err := as.jwt.ParsingPayload(token)
	if err != nil {
		return err
	}

	add.UserID = id

	_, err = as.repo.Address.FindAddressByUserID(id)
	if err != nil {
		add.IsMain = true
	}

	if err := as.repo.Address.CreateAddress(add); err != nil {
		return err
	}

	return nil
}

func (as *addressService) FindAddressByUserID(token string) (*model.Address, error) {

	id, err := as.jwt.ParsingPayload(token)
	if err != nil {
		return nil, err
	}

	addr, err := as.repo.Address.FindAddressByUserID(id)
	if err != nil {
		return nil, err
	}

	return addr, nil
}

func (as *addressService) UpdateAddress(id int, addr *model.Address) error {
	return as.repo.Address.UpdateAddress(id, addr)
}

func (as *addressService) FindAddressByID(id int) (*model.Address, error) {
	return as.repo.Address.FindAddressByID(id)
}

func (as *addressService) DeleteAddress(id int) error {
	return as.repo.Address.DeleteAddress(id)
}
