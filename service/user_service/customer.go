package userservice

import "f-commerce/model"

func (us *userService) UpdateCustomer(id int, cust *model.CustomerData) error {

	if err := us.Repo.User.UpdateCustomer(id, cust); err != nil {
		return err
	}

	return nil
}
