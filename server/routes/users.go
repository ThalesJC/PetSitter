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
	Pets  []Pet  `json:"pets"`
}

func CreateResponseUser(userModel models.User) User {
	pets := make([]Pet, len(userModel.Pets))

	for i, pet := range userModel.Pets {
		pets[i] = Pet{
			Name:      pet.Name,
			User:      CreateResponseUser(pet.User),
			BirthDate: pet.BirthDate,
			Size:      pet.Size,
			Species:   pet.Species,
			Gender:    pet.Gender,
			Neutered:  pet.Neutered,
			CoatColor: pet.CoatColor,
			Picture:   pet.Picture,
		}
	}

	return User{
		ID:    userModel.ID,
		Name:  userModel.Name,
		Email: userModel.Email,
		Pets:  pets,
	}

}

func CreateUser(c *fiber.Ctx) error {
	user := models.User{}

	err := c.BodyParser(&user)
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Petsitter.Db.Create(&user)
	response := CreateResponseUser(user)

	return c.Status(200).JSON(response)
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

func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	user := models.User{}

	database.Petsitter.Db.Preload("Pets").First(&user, id)

	if user.ID == 0 {
		return c.Status(404).JSON("User not found")
	}

	response := CreateResponseUser(user)

	return c.Status(200).JSON(response)
}

func UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	user := models.User{}

	database.Petsitter.Db.First(&user, id)

	if user.ID == 0 {
		return c.Status(404).JSON("User not found")
	}

	err := c.BodyParser(&user)
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Petsitter.Db.Save(&user)

	return c.Status(200).JSON(user)
}

func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	user := models.User{}

	database.Petsitter.Db.First(&user, id)

	if user.ID == 0 {
		return c.Status(404).JSON("User not found")
	}

	database.Petsitter.Db.Delete(&user)

	return c.Status(204).JSON(nil)
}
