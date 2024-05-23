package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/group3/tasks_mgmt/database"
)

func main() {
	database.ConnectDb()
	app := fiber.New()

	setupRoutes(app)

	app.Listen("0.0.0.0:3000")
}
