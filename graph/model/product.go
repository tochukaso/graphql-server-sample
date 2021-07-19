package model

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID        int            `gorm:"primaryKey" json:"id"`
	Name      string         `gorm:"not null;index" json:"name"`
	Price     int            `gorm:"not null" json:"price"`
	Code      string         `json:"code"`
	Detail    string         `json:"detail"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
}
