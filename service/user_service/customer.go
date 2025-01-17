package userservice

import "f-commerce/model"

func (us *userService) CreateCustomer(cust *model.Customer) error {

	if err := us.Repo.User.CreateCustomer(cust); err != nil {
		return err
	}
	return nil
}
