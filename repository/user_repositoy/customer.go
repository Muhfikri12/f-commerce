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

		if result.RowsAffected == 0 {
			return fmt.Errorf("user with id %d not found", id)
		}

		if result.Error != nil {
			return result.Error
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

	if result.RowsAffected == 0 {
		return fmt.Errorf("user with id %d not found", id)
	}

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (uc *userRepo) UpdateRole(id int) error {

	result := uc.db.Table("users").Where("id = ?", id).
		Updates(map[string]interface{}{
			"role":       "seller",
			"updated_at": time.Now(),
		})

	if result.RowsAffected == 0 {
		return fmt.Errorf("user with id %d not found", id)
	}

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (uc *userRepo) UpdateStatus(id int, status string) error {

	result := uc.db.Table("users").Where("id = ?", id).
		Updates(map[string]interface{}{
			"status":     status,
			"updated_at": time.Now(),
		})

	if result.RowsAffected == 0 {
		return fmt.Errorf("user with id %d not found", id)
	}

	if result.Error != nil {
		return result.Error
	}
	return nil
}
