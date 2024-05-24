package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/group3/tasks_mgmt/database"
	"github.com/group3/tasks_mgmt/models"
)

func GetUsers(c *fiber.Ctx) error {
	var users []models.User
	database.DB.Db.Find(&users) // Retrieve all users from the database
	return c.JSON(users)        // Return the users as a JSON response
}

func CreateUser(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil { // Parse the request body into the user model
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	database.DB.Db.Create(user)         // Save the new user to the database
	return c.SendString("User created") // Return a success message
}

func GetUser(c *fiber.Ctx) error {
	id := c.Params("id") // Get the user ID from the URL
	var user models.User
	if database.DB.Db.First(&user, id).Error != nil {
		return c.Status(fiber.StatusNotFound).SendString("User not found")
	}
	return c.JSON(user) // Return the user as a JSON response
}

func UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id") // Get the user ID from the URL
	var user models.User
	if database.DB.Db.First(&user, id).Error != nil {
		return c.Status(fiber.StatusNotFound).SendString("User not found")
	}
	if err := c.BodyParser(&user); err != nil { // Parse the request body into the user model
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	database.DB.Db.Save(&user)          // Save the updated user to the database
	return c.SendString("User updated") // Return a success message
}

func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id") // Get the user ID from the URL
	if database.DB.Db.Delete(&models.User{}, id).Error != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Could not delete user")
	}
	return c.SendString("User deleted") // Return a success message
}
