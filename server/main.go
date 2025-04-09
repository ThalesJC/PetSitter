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

	protected := api.Use(middleware.JWTMiddleware)

	// User routes
	protected.Get("/me", routes.Me)
	protected.Put("/me", routes.UpdateMe)
	protected.Delete("/me", routes.DeleteMe)

	// Pet routes
	protected.Post("/pets", routes.CreatePet)
	protected.Get("/pets", routes.GetAllPets)
	protected.Get("/pets/:id", routes.GetPetByID)
	protected.Put("/pets/:id", routes.UpdatePet)
	protected.Delete("/pets/:id", routes.DeletePet)
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
