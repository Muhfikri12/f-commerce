package model

import (
	"time"

	"gorm.io/gorm"
)

type Address struct {
	ID        int `json:"id" gorm:"primaryKey"`
	UserID    int `json:"user_id"`
	Address   string
	City      string
	State     string
	IsMain    bool
	Latitude  string
	Longitude string
	CreatedAt time.Time       `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time       `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt *gorm.DeletedAt `json:"-" gorm:"index"`
}
