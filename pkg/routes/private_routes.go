package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/truongkma/go-fiber-api-template/app/controllers"
	"github.com/truongkma/go-fiber-api-template/pkg/middleware"
)

// PrivateRoutes func for describe group of private routes.
func PrivateRoutes(a *fiber.App) {
	// Create routes group.
	route := a.Group("/api/v1")

	route.Delete("/user/sign_out", middleware.JWTProtected(), controllers.UserSignOut) // de-authorization user
}
