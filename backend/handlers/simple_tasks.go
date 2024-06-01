package handlers

import (
	"github.com/gofiber/fiber/v2"           // Fiber pkg, a web framework in Go used to build web apps and APIs
	"github.com/group3/tasks_mgmt/database" // Custom pkg containing database connection logic
	"github.com/group3/tasks_mgmt/models"   // Custom pkg containing data models for the app
)

func GetSimpleTasks(c *fiber.Ctx) error { // Fiber context `c` as an argument and returns an error if any occurs
	var tasks []models.SimpleTask // Declares a slice of `SimpleTask` models to hold the retrieved tasks
	database.DB.Db.Find(&tasks)   // Uses GORM to retrieve all tasks from the database
	return c.JSON(tasks)          // Return the tasks as a JSON response
}

func CreateSimpleTask(c *fiber.Ctx) error {
	task := new(models.SimpleTask)             // Creates a new instance of the `SimpleTask` model
	if err := c.BodyParser(task); err != nil { // Parse the request body into the `SimpleTask` model
		return c.Status(fiber.StatusBadRequest).SendString(err.Error()) // If parsing fails, returns a bad request status
	}
	database.DB.Db.Create(task)         // Save the new task to the database
	return c.SendString("Task created") // Return a success message
}

func GetSimpleTask(c *fiber.Ctx) error {
	id := c.Params("id") // Get the task ID from the URL
	var task models.SimpleTask
	if database.DB.Db.First(&task, id).Error != nil {
		return c.Status(fiber.StatusNotFound).SendString("Task not found")
	}
	return c.JSON(task) // Return the task as a JSON response
}

func UpdateSimpleTask(c *fiber.Ctx) error {
	id := c.Params("id")                              // Get the task ID from the URL
	var task models.SimpleTask                        // Declares a `SimpleTask` model to hold the retrieved task
	if database.DB.Db.First(&task, id).Error != nil { // Attempts to find the task by ID
		return c.Status(fiber.StatusNotFound).SendString("Task not found") // If not found, returns a not found status with an error message
	}
	if err := c.BodyParser(&task); err != nil { // Parse the request body into the `SimpleTask` model
		return c.Status(fiber.StatusBadRequest).SendString(err.Error()) // If parsing fails, returns a bad request status
	}
	database.DB.Db.Save(&task)          // Save the updated task to the database
	return c.SendString("Task updated") // Return a success message
}

func DeleteSimpleTask(c *fiber.Ctx) error {
	id := c.Params("id")                                              // Get the task ID from the URL
	if database.DB.Db.Delete(&models.SimpleTask{}, id).Error != nil { // Attempts to delete the task by ID
		return c.Status(fiber.StatusInternalServerError).SendString("Could not delete task") // If deletion fails, returns an internal server error status
	}
	return c.SendString("Task deleted") // Return a success message
}
