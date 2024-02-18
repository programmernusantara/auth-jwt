package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// Config digunakan untuk mendapatkan nilai variabel lingkungan berdasarkan kuncinya
func Config(key string) string {
	// Memuat file .env untuk mengambil variabel lingkungan
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Print("Error loading .env file")
	}
	return os.Getenv(key)
}
