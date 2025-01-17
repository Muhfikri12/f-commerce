package userrepositoy

import (
	"f-commerce/model"
	"fmt"
	"time"

	"gorm.io/gorm"
)

func (c *userRepo) UpdateCustomer(id int, customer *model.Customer) error {

	err := c.db.Transaction(func(tx *gorm.DB) error {

		customer.UpdatedAt = time.Now()

		result := tx.Table("customers").Where("user_id = ?", id).Updates(&customer)
		if result.Error != nil {
			return result.Error
		}

		if result.RowsAffected == 0 {
			return fmt.Errorf("no user found with ID %d", id)
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
		Updates(map[string]interface{}{
			"image":      image,
			"updated_at": time.Now(),
		})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("id %d does not exist", id)
	}

	return nil
}

func (uc *userRepo) UpdateRole(id int) error {

	result := uc.db.Table("users").Where("id = ?", id).
		Updates(map[string]interface{}{
			"role":       "seller",
			"updated_at": time.Now(),
		})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("user %d does not exist", id)
	}

	return nil
}
