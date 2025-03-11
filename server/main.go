package main

import (
	"PetSitter/database"
	"PetSitter/routes"

	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	// Users Routes
	app.Post("api/v1/users", routes.CreateUser)
	app.Get("api/v1/users", routes.GetUsers)
	app.Get("api/v1/users/:id", routes.GetUser)
	app.Put("api/v1/users/:id", routes.UpdateUser)
	app.Delete("api/v1/users/:id", routes.DeleteUser)
}

func main() {
	database.ConnectDB()
	app := fiber.New()

	setupRoutes(app)

	app.Listen(":8080")
}
