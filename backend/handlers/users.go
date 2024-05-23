package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/group3/tasks_mgmt/database"
	"github.com/group3/tasks_mgmt/models"
)

func GetUsers(c *fiber.Ctx) error {
	var users []models.User
	database.DB.Db.Find(&users)
	return c.JSON(users)
}

func CreateUser(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	database.DB.Db.Create(user)
	return c.SendString("User created")
}

func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User
	if database.DB.Db.First(&user, id).Error != nil {
		return c.Status(fiber.StatusNotFound).SendString("User not found")
	}
	return c.JSON(user)
}

func UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User
	if database.DB.Db.First(&user, id).Error != nil {
		return c.Status(fiber.StatusNotFound).SendString("User not found")
	}
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	database.DB.Db.Save(&user)
	return c.SendString("User updated")
}

func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	if database.DB.Db.Delete(&models.User{}, id).Error != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Could not delete user")
	}
	return c.SendString("User deleted")
}
