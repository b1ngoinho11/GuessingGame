package main

import (
	_ "backend/docs"
	"backend/handlers"
	"backend/middlewares"
	"log"

	"github.com/gofiber/swagger"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

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
