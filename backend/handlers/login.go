package handlers

import (
	"backend/utils"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

// @Summary Login
// @Description Authenticates user and sets a token in the cookie
// @Tags auth
// @Accept json
// @Produce json
// @Param request body LoginRequest true "Login request"
// @Success 200 {object} LoginResponse "Login successful"
// @Failure 400 {object} ErrorResponse "Invalid input"
// @Failure 401 {object} ErrorResponse "Invalid credentials"
// @Router /login [post]
func Login(c *fiber.Ctx) error {
	log.Println("[AUTH] POST /login - Incoming request")
	var body map[string]string
	if err := c.BodyParser(&body); err != nil {
		log.Println("[AUTH][ERROR] Failed to parse request body:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid input",
		})
	}

	username, password := body["username"], body["password"]
	log.Println("[AUTH][INFO] Login attempt for username:", username)

	// Validate username and password
	if username == "admin" && password == "admin" {
		// Generate token
		token, err := utils.GenerateToken(username)
		if err != nil {
			log.Println("[AUTH][ERROR] Failed to generate token:", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to generate token",
			})
		}

		log.Println("[AUTH][INFO] Token generated for user:", username)

		// Set the token as a cookie
		c.Cookie(&fiber.Cookie{
			Name:     "token",
			Value:    token,
			Expires:  time.Now().Add(time.Hour * 24),
			HTTPOnly: true,
			Secure:   true,
			SameSite: "Strict",
		})

		// Return successful JSON
		log.Println("[AUTH][INFO] Login successful for user:", username)
		return c.JSON(fiber.Map{
			"message": "Login successful",
		})
	}

	log.Println("[AUTH][WARNING] Failed login attempt for username:", username)

	// Return invalid
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"error": "Invalid credentials",
	})
}
