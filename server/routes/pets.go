package routes

import (
	"PetSitter/database"
	"PetSitter/models"
	"time"

	"github.com/gofiber/fiber/v2"
)

type Pet struct {
	User      User
	Name      string    `json:"name"`
	BirthDate time.Time `json:"birth_date"`
	Weight    float32   `json:"weight"`
	Size      byte      `json:"size"` // 1: Small, 2: Medium, 3: Large
	Species   string    `json:"species"`
	Gender    string    `json:"gender"`
	Neutered  bool      `json:"neutered"`
	CoatColor string    `json:"coat_color"`
	Picture   string    `json:"picture"`
}

func CreateResponsePet(petModel models.Pet, user models.User) Pet {
	return Pet{
		Name:      petModel.Name,
		BirthDate: petModel.BirthDate,
		Weight:    petModel.Weight,
		Size:      petModel.Size,
		Species:   petModel.Species,
		Gender:    petModel.Gender,
		Neutered:  petModel.Neutered,
		CoatColor: petModel.CoatColor,
		Picture:   petModel.Picture,
	}
}

func CreatePet(c *fiber.Ctx) error {
	pet := models.Pet{}

	err := c.BodyParser(&pet)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if pet.UserID == 0 {
		return c.Status(400).JSON(fiber.Map{
			"error": "UserID is required to associate the pet with a user",
		})
	}

	var user models.User
	if result := database.Petsitter.Db.First(&user, pet.UserID); result.Error != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	database.Petsitter.Db.Create(&pet)

	response := CreateResponsePet(pet, user)

	return c.Status(200).JSON(response)
}
