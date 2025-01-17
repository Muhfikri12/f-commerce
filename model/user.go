package model

import (
	"time"

	"gorm.io/gorm"
)

type Customer struct {
	ID          int       `json:"-" gorm:"primaryKey"`
	UserID      int       `json:"user_id"`
	Fullname    string    `json:"fullname" binding:"required,min=5"`
	DateOfBirth time.Time `json:"date_of_birth" binding:"required"`
	Gender      string
	Phone       string
	Image       string
	CreatedAt   time.Time       `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time       `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt   *gorm.DeletedAt `json:"-" gorm:"index"`
}

type Admin struct {
	ID        int             `json:"-" gorm:"primaryKey"`
	UserID    int             `json:"user_id"`
	Fullname  string          `json:"fullname" binding:"required,min=5"`
	Phone     string          `json:"phone" binding:"required"`
	CreatedAt time.Time       `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time       `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt *gorm.DeletedAt `json:"-" gorm:"index"`
}

type Register struct {
	Fullname string `binding:"required"`
	Email    string `binding:"required"`
	Password string `binding:"required,min=8"`
	Role     string `binding:"required"`
}

type CustomerData struct {
	User     User
	Customer Customer
}
