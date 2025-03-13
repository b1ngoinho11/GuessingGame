package handlers

import (
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

var hiddenNumber int

func init() {
	// Generate hidden number
	rand.Seed(time.Now().UnixNano())
	hiddenNumber = rand.Intn(10)
	log.Println("[INIT][INFO] Hidden number initialized:", hiddenNumber)
}

// @Summary Guess a number
// @Description User guesses a hidden number. If correct, regenerates a new number.
// @Tags guess
// @Produce json
// @Param guess path int true "Guess value"
// @Success 201 {object} models.GuessResponse "Correct guess! New number generated."
// @Success 200 {object} models.GuessResponse "Incorrect guess, try again."
// @Failure 400 {object} models.ErrorResponse "Invalid guess"
// @Router /guess/{guess} [post]
func Guess(c *fiber.Ctx) error {
	log.Println("[GUESS][API] POST /guess - Incoming guess")

	// Parse the guess parameter
	guessStr := c.Params("guess")
	guess, err := strconv.Atoi(guessStr)
	if err != nil {
		log.Println("[GUESS][ERROR] Invalid guess:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid guess",
		})
	}

	// Check if the guess is incorrect
	if guess != hiddenNumber {
		log.Println("[GUESS][INFO] Incorrect guess")
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Incorrect guess, try again.",
		})
	}

	log.Println("[GUESS][INFO] Correct guess")

	// Generate a new hidden number
	hiddenNumber = rand.Intn(10)
	log.Println("[GUESS][INFO] New hidden number generated:", hiddenNumber)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Correct guess! New number generated.",
	})
}
