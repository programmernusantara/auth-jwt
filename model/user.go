package model

import "gorm.io/gorm"

// User struct digunakan untuk merepresentasikan data pengguna dalam database
type User struct {
	gorm.Model
	Username string `gorm:"uniqueIndex;not null" json:"username"`
	Email    string `gorm:"uniqueIndex;not null" json:"email"`
	Password string `gorm:"not null" json:"password"`
	Names    string `json:"names"`
}
