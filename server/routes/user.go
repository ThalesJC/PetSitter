package routes

import (
	"PetSitter/database"
	"PetSitter/models"

	"github.com/gofiber/fiber/v2"
)

type User struct {
	ID    uint
	Name  string `json:"name"`
	Email string `json:"email"`
}

func Me(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uint)

	var user models.User
	if err := database.Petsitter.Db.First(&user, userID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	return c.JSON(fiber.Map{
		"id":         user.ID,
		"name":       user.Name,
		"email":      user.Email,
		"created_at": user.CreatedAt,
	})
}

func UpdateMe(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uint)

	var req struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid body"})
	}

	var user models.User
	if err := database.Petsitter.Db.First(&user, userID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	user.Name = req.Name
	user.Email = req.Email
	database.Petsitter.Db.Save(&user)

	return c.JSON(fiber.Map{
		"message": "User updated",
	})
}

func DeleteMe(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uint)

	if err := database.Petsitter.Db.Delete(&models.User{}, userID).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete user",
		})
	}

	return c.JSON(fiber.Map{
		"message": "User account deleted",
	})
}

func CreateResponseUser(userModel models.User) User {
	return User{
		ID:    userModel.ID,
		Name:  userModel.Name,
		Email: userModel.Email,
	}
}

func GetUsers(c *fiber.Ctx) error {
	users := []models.User{}

	database.Petsitter.Db.Find(&users)

	responseUsers := []User{}
	for _, user := range users {
		newUser := CreateResponseUser(user)
		responseUsers = append(responseUsers, newUser)
	}

	return c.Status(200).JSON(responseUsers)
}
