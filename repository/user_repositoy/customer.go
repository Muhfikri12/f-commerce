package userrepositoy

import (
	"f-commerce/model"
	"fmt"

	"gorm.io/gorm"
)

func (c *userRepo) UpdateCustomer(id int, customer *model.CustomerData) error {

	err := c.db.Transaction(func(tx *gorm.DB) error {

		result := tx.Table("users").Where("id = ?", id).Updates(&customer.User)
		if result.Error != nil {
			return result.Error
		}

		if result.RowsAffected == 0 {
			return fmt.Errorf("no user found with ID %d", id)
		}

		if err := tx.Table("customers").Where("user_id = ?", id).Updates(&customer.Customer).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func (c *userRepo) UpdateProfile(id int, image string) error {

	result := c.db.Table("customers").
		Where("user_id = ?", id).
		Update("image", image)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("id %d does not exist", id)
	}

	return nil
}
