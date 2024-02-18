package middleware

import (
	"api-fiber-gorm/config" // Import package config untuk mengambil konfigurasi aplikasi

	jwtware "github.com/gofiber/contrib/jwt" // Import package jwtware untuk middleware otentikasi JWT
	"github.com/gofiber/fiber/v2"            // Import package Fiber untuk framework web
)

// Protected digunakan untuk melindungi rute dengan middleware otentikasi JWT
func Protected() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:   jwtware.SigningKey{Key: []byte(config.Config("SECRET"))}, // Mengatur kunci penandatanganan JWT
		ErrorHandler: jwtError,                                                 // Menangani kesalahan otentikasi JWT
	})
}

// jwtError digunakan untuk menangani kesalahan otentikasi JWT
func jwtError(c *fiber.Ctx, err error) error {
	// Memeriksa apakah kesalahan adalah "Missing or malformed JWT"
	if err.Error() == "Missing or malformed JWT" {
		// Mengembalikan tanggapan dengan status BadRequest dan pesan kesalahan yang sesuai
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"status": "error", "message": "Missing or malformed JWT", "data": nil})
	}
	// Jika kesalahan bukan "Missing or malformed JWT", mengembalikan tanggapan dengan status Unauthorized dan pesan kesalahan yang sesuai
	return c.Status(fiber.StatusUnauthorized).
		JSON(fiber.Map{"status": "error", "message": "Invalid or expired JWT", "data": nil})
}
