package model

import (
	"time"

	"gorm.io/gorm"
)

type Sku struct {
	ID        int            `gorm:"primaryKey" json:"id"`
	ProductId int            `gorm:"not null" json:"productId"`
	Name      string         `gorm:"not null" json:"name"`
	Stock     int            `gorm:"not null" json:"stock"`
	Code      string         `json:"code"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
}
