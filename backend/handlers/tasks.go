package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/group3/tasks_mgmt/database"
	"github.com/group3/tasks_mgmt/models"
)

// GetTasks retrieves all tasks from the database
func GetTasks(c *fiber.Ctx) error {
	userID, ok := c.Locals("userID").(uint)
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "User ID not found in context"})
	}

	var tasks []models.Task
	database.DB.Db.Where("user_id = ?", userID).Find(&tasks)
	return c.JSON(tasks)
}

// CreateTask creates a new task in the database
func CreateTask(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uint)
	task := new(models.Task)
	if err := c.BodyParser(task); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	task.UserID = userID
	database.DB.Db.Create(&task)
	return c.JSON(task)
}

// GetTask retrieves a task from the database by ID
func GetTask(c *fiber.Ctx) error {
	id := c.Params("id")
	var task models.Task
	if err := database.DB.Db.First(&task, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Task not found"})
	}
	return c.JSON(task)
}

// UpdateTask updates a task in the database by ID
func UpdateTask(c *fiber.Ctx) error {
	id := c.Params("id")
	var task models.Task
	if err := database.DB.Db.First(&task, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Task not found"})
	}

	if err := c.BodyParser(&task); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	database.DB.Db.Save(&task)
	return c.JSON(task)
}

// DeleteTask deletes a task from the database by ID
func DeleteTask(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := database.DB.Db.Delete(&models.Task{}, id).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not delete task"})
	}
	return c.SendStatus(fiber.StatusNoContent)
}
