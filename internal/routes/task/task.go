package taskRoutes

import (
	"github.com/gofiber/fiber/v2"
	taskHandler "github.com/skyler-saville/base-api-fiber/internal/handlers/task"
)

func SetupTaskRoutes(router fiber.Router) {
	task := router.Group("/tasks")
	// create a task
	task.Post("/", taskHandler.CreateTask)
	// read all tasks
	task.Get("/", taskHandler.GetTasks)
	// read one task
	task.Get("/:taskID", taskHandler.GetTask)
	// update a task
	task.Put("/:taskID", taskHandler.UpdateTask)
	// delete a task
	task.Delete("/:taskID", taskHandler.DeleteTask)
}