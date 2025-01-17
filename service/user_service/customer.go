package userservice

import (
	"f-commerce/model"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func (us *userService) UpdateCustomer(token string, cust *model.CustomerData) error {

	password, err := bcrypt.GenerateFromPassword([]byte(cust.User.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	cust.User.Password = string(password)

	id, err := us.jwt.ParsingPayload(token)
	if err != nil {
		return fmt.Errorf("failed parsing id from JWT: " + err.Error())
	}

	if err := us.Repo.User.UpdateCustomer(id, cust); err != nil {
		return err
	}

	return nil
}

func (us *userService) UpdateProfile(token string, image string) error {

	id, err := us.jwt.ParsingPayload(token)
	if err != nil {
		return fmt.Errorf("failed parsing id from JWT: " + err.Error())
	}

	if err := us.Repo.User.UpdateProfile(id, image); err != nil {
		return err
	}

	return nil
}

func (us *userService) UpdateRole(token string) error {

	id, err := us.jwt.ParsingPayload(token)
	if err != nil {
		return fmt.Errorf("failed parsing id from JWT: " + err.Error())
	}

	if err := us.Repo.User.UpdateRole(id); err != nil {
		return err
	}

	return nil
}
