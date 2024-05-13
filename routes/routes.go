package routes

import (
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/iqbalmahad/todolist-app-golang.git/controllers"
	"github.com/iqbalmahad/todolist-app-golang.git/databases"
	"github.com/iqbalmahad/todolist-app-golang.git/repositories"
)

func SetupRoutes(app *fiber.App) {
	db := databases.DBInit()
	userRepo := repositories.NewUserRepo(db)
	userTodo := repositories.NewTodoRepo(db)
	userController := controllers.NewUserController(userRepo) // Inisialisasi UserController
	todoController := controllers.NewTodoController(userTodo) // Inisialisasi TodoController
	authRepo := repositories.NewAuthRepo(db)
	authController := controllers.NewAutController(authRepo)

	app.Post("/login", authController.Login)
	app.Get("/", authController.Accessible)

	app.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte("secret")},
	}))
	api := app.Group("/api") // Grupkan rute API di /api

	// Rute untuk pengguna (users)
	users := api.Group("/users")
	users.Post("/", userController.Create)      // Rute untuk membuat pengguna
	users.Get("/", userController.Index)        // Rute untuk mendapatkan semua pengguna
	users.Get("/:id", userController.Read)      // Rute untuk mendapatkan pengguna berdasarkan ID
	users.Put("/:id", userController.Update)    // Rute untuk mengupdate pengguna berdasarkan ID
	users.Delete("/:id", userController.Delete) // Rute untuk menghapus pengguna berdasarkan ID

	// Rute untuk tugas (todos)
	todos := api.Group("/todos")
	todos.Post("/", todoController.Create)      // Rute untuk membuat tugas
	todos.Get("/", todoController.Index)        // Rute untuk mendapatkan semua tugas
	todos.Get("/:id", todoController.Read)      // Rute untuk mendapatkan tugas berdasarkan ID
	todos.Put("/:id", todoController.Update)    // Rute untuk mengupdate tugas berdasarkan ID
	todos.Delete("/:id", todoController.Delete) // Rute untuk menghapus tugas berdasarkan ID

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE",
		AllowHeaders: "Content-Type,Authorization",
	}))
}
