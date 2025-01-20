package userservice

import (
	"f-commerce/model"
	"fmt"
)

func (us *userService) UpdateCustomer(token string, cust *model.Customer) error {

	id, err := us.jwt.ParsingID(token)
	if err != nil {
		return fmt.Errorf("failed parsing id from JWT: " + err.Error())
	}

	if err := us.Repo.User.UpdateCustomer(id, cust); err != nil {
		return err
	}

	return nil
}

func (us *userService) UpdateProfile(token string, image string) error {

	id, err := us.jwt.ParsingID(token)
	if err != nil {
		return fmt.Errorf("failed parsing id from JWT: " + err.Error())
	}

	return us.Repo.User.UpdateProfile(id, image)
}

func (us *userService) UpdateRole(token string) error {

	id, err := us.jwt.ParsingID(token)
	if err != nil {
		return fmt.Errorf("failed parsing id from JWT: " + err.Error())
	}

	return us.Repo.User.UpdateRole(id)
}

func (us *userService) NonactiveAccount(token string) error {
	id, err := us.jwt.ParsingID(token)
	if err != nil {
		return fmt.Errorf("failed parsing id from JWT: " + err.Error())
	}

	status := "nonactive"
	return us.Repo.User.UpdateStatus(id, status)
}
