package routes

import (
	"PetSitter/database"
	"PetSitter/models"
	"time"

	"github.com/gofiber/fiber/v2"
)

type Pet struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Tutor     User      `json:"tutor"`
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

func CreateResponsePet(petModel models.Pet, tutor User) Pet {
	return Pet{
		ID:        petModel.ID,
		Name:      petModel.Name,
		Tutor:     tutor,
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
		return c.Status(400).JSON(err.Error())
	}

	database.Petsitter.Db.Create(&pet)
	response := CreateResponsePet(pet, User{})

	return c.Status(200).JSON(response)
}
