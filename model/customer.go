package model

import (
	"errors"
	"time"
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
	Status      string
	Updated_at  time.Time
}

func (e *Customer) ValidateEmploye() error {

	if e.Gender != "L" && e.Gender != "P" {
		return errors.New("error payload should be L or P")
	}

	if e.Status != "active" && e.Status != "inactive" {
		return errors.New("error payload should be active or inactive")
	}

	return nil
}
