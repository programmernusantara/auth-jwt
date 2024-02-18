package router

import (
	"api-fiber-gorm/handler"    // Mengimport package untuk menangani permintaan HTTP
	"api-fiber-gorm/middleware" // Mengimport package untuk menangani middleware

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// SetupRoutes digunakan untuk menyiapkan rute API
func SetupRoutes(app *fiber.App) {
	// Middleware
	api := app.Group("/api", logger.New()) // Membuat grup rute API dengan logger middleware
	api.Get("/", handler.Hello)            // Menangani permintaan GET ke endpoint root

	// Auth
	auth := api.Group("/auth")         // Membuat grup rute untuk otorisasi
	auth.Post("/login", handler.Login) // Menangani permintaan POST untuk login

	// User
	user := api.Group("/user")                                      // Membuat grup rute untuk operasi pengguna
	user.Get("/:id", handler.GetUser)                               // Menangani permintaan GET untuk mendapatkan pengguna berdasarkan ID
	user.Post("/", handler.CreateUser)                              // Menangani permintaan POST untuk membuat pengguna baru
	user.Patch("/:id", middleware.Protected(), handler.UpdateUser)  // Menangani permintaan PATCH untuk memperbarui pengguna
	user.Delete("/:id", middleware.Protected(), handler.DeleteUser) // Menangani permintaan DELETE untuk menghapus pengguna

	// Product
	product := api.Group("/product")                                      // Membuat grup rute untuk operasi produk
	product.Get("/", handler.GetAllProducts)                              // Menangani permintaan GET untuk mendapatkan semua produk
	product.Get("/:id", handler.GetProduct)                               // Menangani permintaan GET untuk mendapatkan produk berdasarkan ID
	product.Post("/", middleware.Protected(), handler.CreateProduct)      // Menangani permintaan POST untuk membuat produk baru
	product.Delete("/:id", middleware.Protected(), handler.DeleteProduct) // Menangani permintaan DELETE untuk menghapus produk
}
