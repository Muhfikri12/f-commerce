package model

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type Customer struct {
	ID          int       `json:"-" gorm:"primaryKey"`
	UserID      int       `json:"user_id"`
	User        User      `json:"user" binding:"required"`
	Fullname    string    `json:"fullname" binding:"required,min=5"`
	DateOfBirth time.Time `json:"date_of_birth" binding:"required"`
	AddressID   int
	Gender      string
	Image       string
	Created_at  time.Time       `json:"created_at"`
	Updated_at  time.Time       `json:"updated_at"`
	Deleted_at  *gorm.DeletedAt `json:"-" gorm:"index"`
}

func (e *Customer) ValidateEmploye() error {

	if e.Gender != "L" && e.Gender != "P" {
		return errors.New("error payload should be L or P")
	}

	return nil
}
