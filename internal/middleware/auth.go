package middleware

import (
	"github.com/frkntplglu/emir-backend/internal/models"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

// Protected protect routes
func Protected() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:   jwtware.SigningKey{Key: []byte("MYSECRETKEYFROMCONFIG")},
		ErrorHandler: jwtError,
	})
}

func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).JSON(models.FailureResponse{
			Success: false,
			Error: models.ErrorDetails{
				Message: "Missing or malformed JWT",
				Reason:  err.Error(),
			},
		})

	}
	return c.Status(fiber.StatusUnauthorized).JSON(models.FailureResponse{
		Success: false,
		Error: models.ErrorDetails{
			Message: "Invalid or expired JWT",
			Reason:  err.Error(),
		},
	})
}
