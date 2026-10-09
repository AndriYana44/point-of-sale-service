package models

import "time"

type Product struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	SKU        string    `gorm:"size:50;not null;uniqueIndex" json:"sku"`
	Barcode    string    `gorm:"size:100;uniqueIndex" json:"barcode"`
	Name       string    `gorm:"size:150;not null" json:"name"`
	CategoryID uint      `gorm:"not null;index" json:"category_id"`
	Price      float64   `gorm:"type:decimal(15,2);not null" json:"price"`
	IsActive   bool      `gorm:"default:true" json:"is_active"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`

	Category Category `json:"category,omitempty"`
}

func (Product) TableName() string {
	return "mst_products"
}
