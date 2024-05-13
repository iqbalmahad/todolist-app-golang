package main

import (
	"log"

	"github.com/iqbalmahad/todolist-app-golang.git/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Inisialisasi aplikasi Fiber
	app := fiber.New()

	// Middleware CORS untuk mengizinkan akses dari semua domain

	// Setup rute-rute API
	routes.SetupRoutes(app)

	// Menjalankan aplikasi pada port tertentu (misalnya, port 3000)
	log.Fatal(app.Listen(":3000"))
}
