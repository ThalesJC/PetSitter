package main

import (
	"PetSitter/database"
	"PetSitter/routes"
	"log"

	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	api := app.Group("/api/v1")

	users := api.Group("/users")
	users.Post("/", routes.CreateUser)
	users.Get("/", routes.GetUsers)
	users.Get("/:id", routes.GetUser)
	users.Put("/:id", routes.UpdateUser)
	users.Delete("/:id", routes.DeleteUser)

	pets := api.Group("/pets")
	pets.Post("/", routes.CreatePet)
}

func main() {
	database.ConnectDB()

	app := fiber.New()

	setupRoutes(app)

	log.Println("🚀 Servidor rodando em http://localhost:8080")

	if err := app.Listen("0.0.0.0:8080"); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
