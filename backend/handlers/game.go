package handlers

import (
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

// Initialize the random number on program start
var hiddenNumber int

func init() {
	rand.Seed(time.Now().UnixNano()) // Ensure different random values each run
	hiddenNumber = rand.Intn(10)
	log.Println("[INIT] Hidden number initialized:", hiddenNumber)
}

// @Summary Guess a number
// @Description User guesses a hidden number. If correct, regenerates a new number.
// @Tags game
// @Produce json
// @Param guess path int true "Guess value"
// @Success 201 {object} GuessResponse "Correct guess! New number generated."
// @Success 200 {object} GuessResponse "Incorrect guess, try again."
// @Failure 400 {object} ErrorResponse "Invalid guess"
// @Router /guess/{guess} [post]
func Guess(c *fiber.Ctx) error {
	log.Println("[GAME] POST /guess - Incoming request")

	// Parse the guess parameter
	guessStr := c.Params("guess")
	guess, err := strconv.Atoi(guessStr)
	if err != nil {
		log.Println("[GAME][ERROR] Invalid guess:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid guess",
		})
	}

	log.Println("[GAME][INFO] User guessed:", guess)

	// Check if the guess is correct
	if guess == hiddenNumber {
		log.Println("[GAME][INFO] Correct guess")

		// Generate a new hidden number
		hiddenNumber = rand.Intn(100)
		log.Println("[GAME][INFO] New hidden number generated:", hiddenNumber)

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"message": "Correct guess! New number generated.",
		})
	}

	log.Println("[GAME][INFO] Incorrect guess")
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Incorrect guess, try again.",
	})
}
