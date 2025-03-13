package main

import (
	_ "backend/docs"
	"backend/handlers"
	"backend/middlewares"
	"log"

	"github.com/gofiber/swagger"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	// Enable CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://127.0.0.1:5000/",
		AllowMethods:     "GET,POST,PUT,DELETE",
		AllowCredentials: true,
	}))

	// Swagger Route
	app.Get("/swagger/*", swagger.HandlerDefault)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Post("/login", handlers.Login)
	app.Post("/guess/:guess", middlewares.AuthMiddleware, handlers.Guess)

	// Start the server
	log.Fatal(app.Listen(":3000"))
}
