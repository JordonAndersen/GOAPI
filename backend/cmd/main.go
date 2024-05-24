package main

import (
	"github.com/gofiber/fiber/v2"           // Import Fiber web framework
	"github.com/group3/tasks_mgmt/database" // Import the database package from your project
)

func main() {
	database.ConnectDb() // Connect to the database
	app := fiber.New()   // Create a new Fiber application

	setupRoutes(app) // Set up the API routes

	app.Listen("0.0.0.0:3000") // Start the server on port 3000
}
