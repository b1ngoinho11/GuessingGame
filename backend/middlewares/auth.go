package middlewares

import (
	"backend/utils"
	"log"

	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(c *fiber.Ctx) error {
	token := c.Cookies("token")
	if token == "" {
		log.Println("[MIDDLEWARE][ERROR] Unauthorized - Missing Token")
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized - Missing Token",
		})
	}

	// Validate token
	if err := utils.ValidateToken(token); err != nil {
		log.Println("[MIDDLEWARE][ERROR] Unauthorized - Invalid Token")
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid Token",
		})
	}

	log.Println("[MIDDLEWARE][INFO] Token validated successfully")
	return c.Next()
}
