package routes

import (
	"PetSitter/utils"
	"time"

	"github.com/gofiber/fiber/v2"
)

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token"`
}

func RefreshToken(c *fiber.Ctx) error {
	var req RefreshTokenRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	claims, err := utils.ParseJWT(req.RefreshToken)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid refresh token",
		})
	}

	newToken, err := utils.GenerateJWT(claims.UserID, claims.Email, 15*time.Minute)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not generate new token",
		})
	}

	return c.JSON(fiber.Map{
		"token": newToken,
	})
}
