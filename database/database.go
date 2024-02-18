package database

import (
	"api-fiber-gorm/config" // Mengimport package untuk konfigurasi aplikasi
	"api-fiber-gorm/model"  // Mengimport model yang diperlukan
	"fmt"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB adalah instance GORM yang digunakan untuk koneksi ke database
var DB *gorm.DB

// ConnectDB digunakan untuk membuat koneksi ke database dan melakukan migrasi model
func ConnectDB() {
	var err error
	p := config.Config("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)

	if err != nil {
		panic("failed to parse database port")
	}

	// Membuat DSN (Data Source Name) untuk koneksi ke database PostgreSQL
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Config("DB_HOST"), port, config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_NAME"))
	// Membuka koneksi ke database menggunakan GORM
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Connection Opened to Database")
	// Melakukan migrasi model ke database
	DB.AutoMigrate(&model.Product{}, &model.User{})
	fmt.Println("Database Migrated")
}
