package userservice

import (
	"f-commerce/model"
	"fmt"
)

func (us *userService) UpdateAdmin(token string, admin *model.Admin) error {

	id, err := us.jwt.ParsingPayload(token)
	if err != nil {
		return fmt.Errorf("failed parsing id from JWT: " + err.Error())
	}

	if err := us.Repo.User.UpdateAdmin(id, admin); err != nil {
		return err
	}

	return nil
}
