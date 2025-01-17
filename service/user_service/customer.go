package userservice

import (
	"f-commerce/model"
	"fmt"
	"strconv"
)

func (us *userService) UpdateCustomer(token string, cust *model.CustomerData) error {

	idStr, err := us.jwt.ParsingPayload(token)
	if err != nil {
		return fmt.Errorf("failed parsing id from JWT: " + err.Error())
	}

	id, _ := strconv.Atoi(idStr.(string))

	if err := us.Repo.User.UpdateCustomer(id, cust); err != nil {
		return err
	}

	return nil
}

func (us *userService) UpdateProfile(token string, image string) error {

	idStr, err := us.jwt.ParsingPayload(token)
	if err != nil {
		return fmt.Errorf("failed parsing id from JWT: " + err.Error())
	}

	id, _ := strconv.Atoi(idStr.(string))

	if err := us.Repo.User.UpdateProfile(id, image); err != nil {
		return err
	}

	return nil
}
