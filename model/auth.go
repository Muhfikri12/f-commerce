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

type VerificationEmail struct {
	Email string `binding:"required"`
	Otp   string `binding:"required"`
}

type User struct {
	Id        int    `json:"id" gorm:"primaryKey"`
	Email     string `json:"email" gorm:"unique" binding:"required,email"`
	Password  string `binding:"required,min=8"`
	Role      string `json:"role"`
	Username  string `json:"username" gorm:"unique"`
	Status    string
	CreatedAt time.Time       `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time       `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt *gorm.DeletedAt `json:"-" gorm:"index"`
}

func (u *User) Validate() error {
	if u.Role != "admin" && u.Role != "customer" {
		return errors.New("role must be either 'admin' or 'customer'")
	}

	return nil
}
