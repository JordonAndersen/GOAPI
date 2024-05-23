package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/group3/tasks_mgmt/database"
	"github.com/group3/tasks_mgmt/models"
)

func GetTasks(c *fiber.Ctx) error {
	var tasks []models.Task
	database.DB.Db.Find(&tasks)
	return c.JSON(tasks)
}

func CreateTask(c *fiber.Ctx) error {
	task := new(models.Task)
	if err := c.BodyParser(task); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	database.DB.Db.Create(task)
	return c.SendString("Task created")
}

func GetTask(c *fiber.Ctx) error {
	id := c.Params("id")
	var task models.Task
	if database.DB.Db.First(&task, id).Error != nil {
		return c.Status(fiber.StatusNotFound).SendString("Task not found")
	}
	return c.JSON(task)
}

func UpdateTask(c *fiber.Ctx) error {
	id := c.Params("id")
	var task models.Task
	if database.DB.Db.First(&task, id).Error != nil {
		return c.Status(fiber.StatusNotFound).SendString("Task not found")
	}
	if err := c.BodyParser(&task); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	database.DB.Db.Save(&task)
	return c.SendString("Task updated")
}

func DeleteTask(c *fiber.Ctx) error {
	id := c.Params("id")
	if database.DB.Db.Delete(&models.Task{}, id).Error != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Could not delete task")
	}
	return c.SendString("Task deleted")
}
