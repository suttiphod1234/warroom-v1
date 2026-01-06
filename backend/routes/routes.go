package routes

import (
	"warroom-backend/handlers"
	"warroom-backend/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	// Auth Routes
	auth := api.Group("/auth")
	auth.Post("/register", handlers.Register)
	auth.Post("/login", handlers.Login)
	auth.Get("/profile", middleware.Protected(), handlers.GetProfile)

	// MLM Routes
	mlm := api.Group("/mlm")
	mlm.Post("/purchase", handlers.Purchase)
}
