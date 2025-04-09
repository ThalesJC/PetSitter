package main

import (
	"PetSitter/database"
	"PetSitter/middleware"
	"PetSitter/routes"
	"log"

	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	api := app.Group("/api/v1")

	api.Post("/register", routes.Register)
	api.Post("/login", routes.Login)
	api.Post("/token/refresh", routes.RefreshToken)

	// TODO: remover essa rota
	api.Get("/user/all", routes.GetUsers)

	protected := api.Group("", middleware.JWTMiddleware)
	protected.Get("/me", routes.Me)
	protected.Put("/me", routes.UpdateMe)
	protected.Delete("/me", routes.DeleteMe)

	pets := api.Group("/pets", middleware.JWTMiddleware)
	pets.Post("/", routes.CreatePet)
	pets.Get("/:id", routes.GetPets)
	pets.Get("/userID::userID/petID::petID", routes.GetPetById)
	pets.Put("/:id", routes.UpdatePet)
	pets.Delete("/:id", routes.DeletePet)
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
