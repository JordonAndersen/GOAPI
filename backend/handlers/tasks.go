package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/group3/tasks_mgmt/database"
	"github.com/group3/tasks_mgmt/models"
)

func GetTasks(c *fiber.Ctx) error {
	var tasks []models.Task
	database.DB.Db.Find(&tasks) // Retrieve all tasks from the database
	return c.JSON(tasks)        // Return the tasks as a JSON response
}

func CreateTask(c *fiber.Ctx) error {
	task := new(models.Task)
	if err := c.BodyParser(task); err != nil { // Parse the request body into the task model
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	database.DB.Db.Create(task)         // Save the new task to the database
	return c.SendString("Task created") // Return a success message
}

func GetTask(c *fiber.Ctx) error {
	id := c.Params("id") // Get the task ID from the URL
	var task models.Task
	if database.DB.Db.First(&task, id).Error != nil {
		return c.Status(fiber.StatusNotFound).SendString("Task not found")
	}
	return c.JSON(task) // Return the task as a JSON response
}

func UpdateTask(c *fiber.Ctx) error {
	id := c.Params("id") // Get the task ID from the URL
	var task models.Task
	if database.DB.Db.First(&task, id).Error != nil {
		return c.Status(fiber.StatusNotFound).SendString("Task not found")
	}
	if err := c.BodyParser(&task); err != nil { // Parse the request body into the task model
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	database.DB.Db.Save(&task)          // Save the updated task to the database
	return c.SendString("Task updated") // Return a success message
}

func DeleteTask(c *fiber.Ctx) error {
	id := c.Params("id") // Get the task ID from the URL
	if database.DB.Db.Delete(&models.Task{}, id).Error != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Could not delete task")
	}
	return c.SendString("Task deleted") // Return a success message
}
