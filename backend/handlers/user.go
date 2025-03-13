package handlers

import (
	"backend/models"
	"backend/utils"
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type UserHandler struct {
	DB *gorm.DB
}

// @Summary Login
// @Description Authenticates user and sets a token in the cookie
// @Tags users
// @Accept json
// @Produce json
// @Param request body models.UserRequest true "User request"
// @Success 200 {object} models.LoginResponse "Login successful"
// @Failure 400 {object} models.ErrorResponse "Invalid input"
// @Failure 401 {object} models.ErrorResponse "Invalid credentials"
// @Router /users/login [post]
func (h *UserHandler) Login(c *fiber.Ctx) error {
	log.Println("[USER][API] POST /login - Login request")

	// Parse the request body
	var body models.UserRequest
	if err := c.BodyParser(&body); err != nil {
		log.Println("[USER][ERROR] Failed to parse request body:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	username, password := body.Username, body.Password

	// Match username and password to the existing
	var user models.User
	if err := h.DB.Where("username = ? AND password = ?", username, password).First(&user).Error; err != nil {
		log.Println("[USER][INFO] Failed login attempt for username:", username)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	// Generate token
	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		log.Println("[USER][ERROR] Failed to generate token:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to generate token"})
	}
	log.Println("[USER][INFO] Token generated for user:", username)

	// Set cookie
	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: false,
		Secure:   true,
		SameSite: "None",
	})

	log.Println("[USER][INFO] Login successful for user:", username)
	return c.JSON(fiber.Map{"message": "Login successful"})
}

// @Summary Create User
// @Description Creates a new user
// @Tags users
// @Accept json
// @Produce json
// @Param request body models.UserRequest true "User request"
// @Success 201 {object} models.User "User created"
// @Failure 400 {object} models.ErrorResponse "Invalid input"
// @Failure 500 {object} models.ErrorResponse "Failed to create user"
// @Router /users [post]
func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	log.Println("[USER][API] POST /users - Creating user")

	// Parse the request body
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		log.Println("[USER][ERROR] Failed to parse request body:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	// Create user
	if err := h.DB.Create(&user).Error; err != nil {
		log.Println("[USER][ERROR] Failed to create user:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": fmt.Sprintf("Failed to create user (%v)", err)})
	}

	log.Println("[USER][INFO] User created successfully:", user.ID)
	return c.Status(fiber.StatusCreated).JSON(user)
}

// @Summary Get User
// @Description Retrieves a user by ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} models.User "User found"
// @Failure 404 {object} models.ErrorResponse "User not found"
// @Router /users/{id} [get]
func (h *UserHandler) GetUser(c *fiber.Ctx) error {
	log.Println("[USER][API] GET /users/ - Retrieving user")

	id := c.Params("id")

	// Retrieve user from ID
	var user models.User
	if err := h.DB.First(&user, id).Error; err != nil {
		log.Println("[USER][ERROR] User not found:", id)
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	log.Println("[USER][INFO] User found:", user.ID)
	return c.JSON(user)
}

// @Summary Update User
// @Description Updates a user's information
// @Tags users
// @Accept json
// @Produce json
// @Param request body models.UserRequest true "User request"
// @Success 200 {object} models.User "User updated"
// @Failure 400 {object} models.ErrorResponse "Invalid input"
// @Failure 404 {object} models.ErrorResponse "User not found"
// @Router /users [put]
func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	log.Println("[USER][API] PUT /users - Updating user")

	// Get the token from cookie
	token := c.Cookies("token")
	if token == "" {
		log.Println("[USER][ERROR] No token found in cookies")
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	// Decode the token to get the user ID
	id, err := utils.DecodeToken(token)
	if err != nil {
		log.Println("[USER][ERROR] Failed to decode token:", err)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	// Retrieve user from ID
	var user models.User
	if err := h.DB.First(&user, id).Error; err != nil {
		log.Println("[USER][ERROR] User not found:", id)
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	// Parse the request body
	var updateData models.User
	if err := c.BodyParser(&updateData); err != nil {
		log.Println("[USER][ERROR] Failed to parse request body:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	// Check valid username
	if updateData.Username != "" && updateData.Username != user.Username {
		var existingUser models.User
		if err := h.DB.Where("username = ?", updateData.Username).First(&existingUser).Error; err == nil {
			log.Println("[USER][ERROR] Username already exists:", updateData.Username)
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": "Username already exists"})
		}
		user.Username = updateData.Username
	}

	// Update password
	if updateData.Password != "" {
		user.Password = updateData.Password
	}

	h.DB.Save(&user)

	log.Println("[USER][INFO] User updated successfully:", user.ID)
	return c.JSON(user)
}

// @Summary Delete User
// @Description Deletes a user by ID
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {object} models.ErrorResponse "User deleted"
// @Failure 500 {object} models.ErrorResponse "Failed to delete user"
// @Router /users [delete]
func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	log.Println("[USER][API] DELETE /users - Deleting user")

	// Get the token from cookie
	token := c.Cookies("token")
	if token == "" {
		log.Println("[USER][ERROR] No token found in cookies")
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	// Decode the token to get the user ID
	id, err := utils.DecodeToken(token)
	if err != nil {
		log.Println("[USER][ERROR] Failed to decode token:", err)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	// Delete user from ID
	if err := h.DB.Delete(&models.User{}, id).Error; err != nil {
		log.Println("[USER][ERROR] Failed to delete user:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete user"})
	}

	log.Println("[USER][INFO] User deleted successfully:", id)
	return c.JSON(fiber.Map{"message": "UserID:" + fmt.Sprint(id) + " deleted"})
}

// @Summary Get All Users
// @Description Retrieves a list of all users
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {array} models.User "List of users"
// @Failure 500 {object} models.ErrorResponse "Failed to fetch users"
// @Router /users [get]
func (h *UserHandler) GetAllUsers(c *fiber.Ctx) error {
	log.Println("[USER][API] GET /users - Retrieving all users")

	// Retrieve all users
	var users []models.User
	if err := h.DB.Find(&users).Error; err != nil {
		log.Println("[USER][ERROR] Failed to fetch users:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch users"})
	}

	log.Println("[USER][INFO] Users retrieved successfully")
	return c.JSON(users)
}
