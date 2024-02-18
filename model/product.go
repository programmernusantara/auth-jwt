package model

import "gorm.io/gorm"

// Product struct digunakan untuk merepresentasikan data produk dalam database
type Product struct {
	gorm.Model
	Title       string `gorm:"not null" json:"title"`
	Description string `gorm:"not null" json:"description"`
	Amount      int    `gorm:"not null" json:"amount"`
}
