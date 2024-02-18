package handler

import (
	// Mengimport package yang diperlukan untuk menangani permintaan HTTP

	"github.com/gofiber/fiber/v2"
)

// Hello digunakan untuk menangani permintaan status API
func Hello(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "success", "message": "Hello i'm ok!", "data": nil})
}
