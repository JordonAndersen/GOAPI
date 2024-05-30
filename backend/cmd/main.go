package main

import (
	"github.com/gofiber/fiber/v2" // Import Fiber web framework
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/group3/tasks_mgmt/database" // Import the database package from your project
)

func main() {
	database.ConnectDb() // Connect to the database
	app := fiber.New()   // Create a new Fiber application

	// CORS configuration
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173",
		AllowHeaders: "Origin, Content-Type, Accept",
		AllowMethods: "GET, POST, PUT, DELETE",
	}))

	setupRoutes(app) // Set up the API routes

	app.Listen("0.0.0.0:3000") // Start the server on port 3000
}
