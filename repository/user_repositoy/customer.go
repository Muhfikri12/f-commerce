package userrepositoy

import "f-commerce/model"

func (c *userRepo) CreateCustomer(cust *model.Customer) error {

	if err := c.db.Create(&cust).Error; err != nil {
		return err
	}

	return nil
}

func (c *userRepo) UpdateCustomer(customer *model.Customer) error {

	return nil
}
