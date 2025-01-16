package model

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type Login struct {
	Input    string `json:"input" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type User struct {
	Id         int             `json:"id" gorm:"primaryKey"`
	Email      string          `json:"email" gorm:"unique" binding:"required,email"`
	Password   string          `json:"-" binding:"required,min=8"`
	Role       string          `json:"role"`
	Username   string          `json:"username" gorm:"unique"`
	Created_at time.Time       `json:"created_at"`
	Updated_at time.Time       `json:"updated_at"`
	Deleted_at *gorm.DeletedAt `json:"-" gorm:"index"`
}

func (u *User) Validate() error {
	if u.Role != "admin" && u.Role != "employee" {
		return errors.New("role must be either 'admin' or 'employee'")
	}
	return nil
}
