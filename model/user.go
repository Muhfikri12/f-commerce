package model

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type Customer struct {
	ID          int       `json:"-" gorm:"primaryKey"`
	UserID      int       `json:"user_id"`
	Fullname    string    `json:"fullname" binding:"required,min=5"`
	DateOfBirth time.Time `json:"date_of_birth" binding:"required"`
	AddressID   int
	Gender      string
	Image       string
	CreatedAt   time.Time       `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time       `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt   *gorm.DeletedAt `json:"-" gorm:"index"`
}

func (e *Customer) ValidateGender() error {

	if e.Gender != "L" && e.Gender != "P" {
		return errors.New("error payload should be L or P")
	}

	return nil
}

type Register struct {
	Fullname string `binding:"required"`
	Email    string `binding:"required"`
	Password string `binding:"required,min=8"`
}

type CustomerData struct {
	User     User
	Customer Customer
}
