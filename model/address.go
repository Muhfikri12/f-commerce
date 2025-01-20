package model

import (
	"time"
)

type Address struct {
	ID        int    `json:"id" gorm:"primaryKey"`
	UserID    int    `json:"user_id"`
	Address   string `binding:"required"`
	City      string `binding:"required"`
	State     string `binding:"required"`
	IsMain    bool   `gorm:"default:false"`
	Latitude  string
	Longitude string
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
