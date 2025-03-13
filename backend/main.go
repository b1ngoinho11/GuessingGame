package main

import (
	_ "backend/docs"
	"backend/handlers"
	"backend/middlewares"
	"backend/models"
	"log"

	"github.com/gofiber/swagger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	// Database setup
	db, err := gorm.Open(sqlite.Open("database/database.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	log.Println("[INIT][INFO] Connected to the database")
	db.AutoMigrate(&models.User{})

	userHandler := handlers.UserHandler{DB: db}

	app := fiber.New()

	// Enable CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://127.0.0.1:5000/",
		AllowMethods:     "GET,POST,PUT,DELETE",
		AllowCredentials: true,
	}))

	// Swagger Route
	app.Get("/swagger/*", swagger.HandlerDefault)

	// Test Route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// API Route
	app.Post("/guess/:guess", middlewares.AuthMiddleware, handlers.Guess)

	app.Post("/users/login", userHandler.Login)
	app.Post("/users", userHandler.CreateUser)
	app.Get("/users", userHandler.GetAllUsers)
	app.Get("/users/:id", userHandler.GetUser)
	app.Put("/users", middlewares.AuthMiddleware, userHandler.UpdateUser)
	app.Delete("/users", middlewares.AuthMiddleware, userHandler.DeleteUser)

	// Start the server
	log.Fatal(app.Listen(":3000"))
}
