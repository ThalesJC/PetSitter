package main

import (
	"PetSitter/database"

	"github.com/gofiber/fiber/v2"
)

func welcome(context *fiber.Ctx) error {
	return context.SendString("Bem vindo ao PetSitter!")
}

func main() {
	database.ConnectDB()
	app := fiber.New()

	app.Get("/", welcome)

	app.Listen(":8080")
}
