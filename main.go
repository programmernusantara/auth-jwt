package main

import (
	"api-fiber-gorm/database" // Mengimport package untuk koneksi database
	"api-fiber-gorm/router"   // Mengimport package untuk routing
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()  // Membuat instance baru dari Fiber framework
	app.Use(cors.New()) // Menambahkan middleware untuk menangani CORS

	database.ConnectDB() // Menghubungkan ke database

	router.SetupRoutes(app) // Menyiapkan routing untuk aplikasi Fiber

	log.Fatal(app.Listen(":3000")) // Mendengarkan koneksi pada port 3000 dan menangani kesalahan jika terjadi
}
