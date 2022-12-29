package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/truongkma/go-fiber-api-template/app/controllers"
)

// PublicRoutes func for describe group of public routes.
func PublicRoutes(a *fiber.App) {
	// Create routes group.
	route := a.Group("/api/v1")

	// Routes for POST method:
	route.Post("/user/sign_up", controllers.UserSignUp) // register a new user
}
