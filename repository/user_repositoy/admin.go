package userrepositoy

import (
	"f-commerce/model"
	"fmt"
	"time"

	"gorm.io/gorm"
)

func (ur *userRepo) UpdateAdmin(id int, admin *model.Admin) error {

	err := ur.db.Transaction(func(tx *gorm.DB) error {

		admin.UpdatedAt = time.Now()

		result := tx.Table("admins").
			Where("user_id = ?", id).
			Updates(&admin)

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
