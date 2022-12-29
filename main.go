package main

import (
	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload" // load .env file automatically
	"github.com/truongkma/go-fiber-api-template/pkg/configs"
	"github.com/truongkma/go-fiber-api-template/pkg/middleware"
	"github.com/truongkma/go-fiber-api-template/pkg/utils"
	"os"
)

func main() {
	// Define Fiber config.
	config := configs.FiberConfig()
	app := fiber.New(config)

	// Middlewares.
	middleware.FiberMiddleware(app) // Register Fiber's middleware for app.

	// Start server (with or without graceful shutdown).
	if os.Getenv("STAGE_STATUS") == "dev" {
		utils.StartServer(app)
	} else {
		utils.StartServerWithGracefulShutdown(app)
	}
}
