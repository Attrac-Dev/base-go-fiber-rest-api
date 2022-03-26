package taskHandler

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/skyler-saville/base-api-fiber/database"
	"github.com/skyler-saville/base-api-fiber/internal/model"
)

// Create a task
func CreateTask(c *fiber.Ctx) error {
	db := database.DB
	task := new(model.Task)

	// Store the body in the task and return error if encountered
	err := c.BodyParser(task)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status":"error", "message":"Review input", "data": err})
	}
	// Add a uuid to the task
	task.ID = uuid.New()
	// Create the task and return error if encountered
	err2 := db.Create(&task).Error
	if err2 != nil {
		return c.Status(500).JSON(fiber.Map{"status":"error", "message":"could not create the task", "data": err2})
	}

	// Return the created task
	return c.JSON(fiber.Map{"status":"success", "message":"task created", "data": task})
}

// Read all tasks
func GetTasks(c *fiber.Ctx) error {
	db := database.DB
	var tasks []model.Task

	//find all the tasks in the database
	db.Find(&tasks)

	// return an error if no task was located
	if len(tasks) == 0 {
		return c.Status(404).JSON(fiber.Map{"status":"error", "message":"no tasks were found", "data": nil})
	}
	
	// return results if tasks found
	return c.JSON(fiber.Map{"status":"success", "message":"tasks located", "data": tasks})
}

// Read one task
func GetTask(c *fiber.Ctx) error {
	db := database.DB
	var task model.Task

	// Read the param of taskID
	id := c.Params("taskID")

	// Find the task with the provided id
	db.Find(&task, "id = ?", id)

	// If no task is found with the provided id, return an error
	if task.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status":"error", "message":"no task present with the provided id", "data": nil})
	}

	// Return the task with the provided id
	return c.JSON(fiber.Map{"status":"success", "message": "task located", "data": task})
}

// Update task
func UpdateTask(c *fiber.Ctx) error {
	// struct must match original model, except without providing the ID, so it remains the same
	type updateTask struct {
		Title			*string 		`json:"title"`
		Subtitle		string		`json:"subtitle"`
		Text			string		`json:"text"`
		CompletedOnDate	time.Time	`json:"completedondate"`
	}
	db := database.DB
	var task model.Task

	// Read the provided param
	id := c.Params("taskID")

	// find the task with the provided id
	db.Find(&task, "id = ?", id)

	// if no task was found with the provided id, return an error
	if task.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status":"error", "message":"no task was found matching provided id", "data": nil})
	}

	// store the body containing the updated data and return error if encountered
	var updateTaskData updateTask
	err := c.BodyParser(&updateTaskData)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status":"error", "message":"server-side error. Review your input.", "data": nil})
	}

	// Edit the task
	task.Title = updateTaskData.Title
	task.Subtitle = updateTaskData.Subtitle
	task.Text = updateTaskData.Text

	// Save the Changes
	db.Save(&task)

	// Return the updated task
	return c.JSON(fiber.Map{"status": "success", "message": "task updated", "data": task})
}

// Delete a task
func DeleteTask(c *fiber.Ctx) error {
	db := database.DB
	var task model.Task

	// read the provided param
	id := c.Params("taskID")

	// find the task with the provided id
	db.Find(&task, "id =?", id)

	// if no task is present, return an error
	if task.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status":"error", "message":"no task found", "data": nil})
	}

	// delete the task and return an error if encountered
	err := db.Delete(&task, "id=?", id).Error

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status":"error", "message":"no task was found", "data":nil})
	}

	// return success message after delete
	return c.JSON(fiber.Map{"status":"success", "message":"task was deleted"})
}