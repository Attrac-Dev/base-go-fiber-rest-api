package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	taskRoutes "github.com/skyler-saville/base-api-fiber/internal/routes/task"
)

func SetupRoutes(app *fiber.App) {
	// Group endpoints with param 'api' and log whenever this endpoint is hit.
	api := app.Group("/api", logger.New()) 

	// setup the task routes, can use same syntax to add routes for additional models
	taskRoutes.SetupTaskRoutes(api)
}