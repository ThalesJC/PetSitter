package routes

import (
	"PetSitter/database"
	"PetSitter/models"
	"time"

	"github.com/gofiber/fiber/v2"
)

type Pet struct {
	ID        uint      `json:"id"`
	UserID    uint      `json:"user_id"`
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
		ID:        petModel.ID,
		UserID:    petModel.UserID,
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
	data := struct {
		UserID    uint    `json:"user_id"`
		Name      string  `json:"name"`
		BirthDate string  `json:"birth_date"`
		Weight    float32 `json:"weight"`
		Size      byte    `json:"size"`
		Species   string  `json:"species"`
		Gender    string  `json:"gender"`
		Neutered  bool    `json:"neutered"`
		CoatColor string  `json:"coat_color"`
		Picture   string  `json:"picture"`
	}{}

	if err := c.BodyParser(&data); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	var user models.User
	if result := database.Petsitter.Db.First(&user, data.UserID); result.Error != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	pet := models.Pet{
		UserID:    data.UserID,
		Name:      data.Name,
		BirthDate: parseBirthDate(data.BirthDate),
		Weight:    data.Weight,
		Size:      data.Size,
		Species:   data.Species,
		Gender:    data.Gender,
		Neutered:  data.Neutered,
		CoatColor: data.CoatColor,
		Picture:   data.Picture,
	}

	database.Petsitter.Db.Create(&pet)

	return c.Status(200).JSON(pet)
}

func GetPets(c *fiber.Ctx) error {
	userID := c.Params("id")

	if userID == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "User ID is required",
		})
	}

	pets := []models.Pet{}

	result := database.Petsitter.Db.Where("user_id = ?", userID).Find(&pets)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Error retrieving pets",
		})
	}

	if len(pets) == 0 {
		return c.Status(404).JSON(fiber.Map{
			"error": "No pets found for this user",
		})
	}

	responsePets := []Pet{}
	for _, pet := range pets {
		var user models.User
		database.Petsitter.Db.First(&user, pet.UserID)
		newPet := CreateResponsePet(pet, user)
		responsePets = append(responsePets, newPet)
	}

	return c.Status(200).JSON(responsePets)
}

func parseBirthDate(dateStr string) time.Time {
	layout := "2006-01-02 15:04:05"
	birthDate, err := time.Parse(layout, dateStr)
	if err != nil {
		return time.Time{}
	}
	return birthDate
}
