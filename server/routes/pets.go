package routes

import (
	"PetSitter/database"
	"PetSitter/models"

	"github.com/gofiber/fiber/v2"
)

func CreatePet(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uint)

	var pet models.Pet
	if err := c.BodyParser(&pet); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	pet.UserID = userID
	if err := database.Petsitter.Db.Create(&pet).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create pet"})
	}

	return c.Status(fiber.StatusCreated).JSON(pet)
}

func GetAllPets(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uint)

	var pets []models.Pet
	if err := database.Petsitter.Db.Where("user_id = ?", userID).Find(&pets).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve pets"})
	}

	return c.JSON(pets)
}

func GetPetByID(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uint)
	petID := c.Params("id")

	var pet models.Pet
	if err := database.Petsitter.Db.Where("id = ? AND user_id = ?", petID, userID).First(&pet).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Pet not found"})
	}

	return c.JSON(pet)
}

func UpdatePet(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uint)
	petID := c.Params("id")

	var pet models.Pet
	if err := database.Petsitter.Db.Where("id = ? AND user_id = ?", petID, userID).First(&pet).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Pet not found"})
	}

	if err := c.BodyParser(&pet); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	if err := database.Petsitter.Db.Save(&pet).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update pet"})
	}

	return c.JSON(pet)
}

func DeletePet(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uint)
	petID := c.Params("id")

	var pet models.Pet
	if err := database.Petsitter.Db.Where("id = ? AND user_id = ?", petID, userID).First(&pet).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Pet not found"})
	}

	if err := database.Petsitter.Db.Delete(&pet).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete pet"})
	}

	return c.JSON(fiber.Map{"message": "Pet deleted successfully"})
}
