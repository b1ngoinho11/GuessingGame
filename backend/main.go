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

	// User API Route
	userRoutes := app.Group("/users")
	{
		userRoutes.Post("/login", userHandler.Login)                              // Login user
		userRoutes.Post("", userHandler.CreateUser)                               // Create user
		userRoutes.Get("", userHandler.GetAllUsers)                               // Get all users
		userRoutes.Get("/:id", userHandler.GetUser)                               // Get user by ID
		userRoutes.Put("", middlewares.AuthMiddleware, userHandler.UpdateUser)    // Update user
		userRoutes.Delete("", middlewares.AuthMiddleware, userHandler.DeleteUser) // Delete user
	}

	// Guess API Route
	guessRoutes := app.Group("/guess")
	{
		guessRoutes.Post("/:guess", middlewares.AuthMiddleware, handlers.Guess) // Make a guess
	}

	// Start the server
	log.Fatal(app.Listen(":3000"))
}
