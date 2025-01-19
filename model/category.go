package model

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	ID        int             `json:"id" gorm:"primaryKey"`
	Name      string          `json:"name" gorm:"unique"`
	CreatedAt time.Time       `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time       `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt *gorm.DeletedAt `json:"-" gorm:"index"`
}
