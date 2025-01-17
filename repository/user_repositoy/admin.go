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

		if result.Error != nil {
			return result.Error
		}

		if result.RowsAffected == 0 {
			return fmt.Errorf("error id %d does not exist", id)
		}

		return nil
	})

	if err != nil {
		return err
	}
	return nil
}
