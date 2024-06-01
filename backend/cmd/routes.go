package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/group3/tasks_mgmt/handlers"
)

func setupRoutes(app *fiber.App) {
	// Routes for tasks with user association
	app.Get("/tasks", handlers.GetTasks)
	app.Post("/tasks", handlers.CreateTask)
	app.Get("/tasks/:id", handlers.GetTask)
	app.Put("/tasks/:id", handlers.UpdateTask)
	app.Delete("/tasks/:id", handlers.DeleteTask)

	// Routes for users
	app.Get("/users", handlers.GetUsers)
	app.Post("/users", handlers.CreateUser)
	app.Get("/users/:id", handlers.GetUser)
	app.Put("/users/:id", handlers.UpdateUser)
	app.Delete("/users/:id", handlers.DeleteUser)

	// Routes for tasks without user association
	app.Get("/simple-tasks", handlers.GetSimpleTasks)
	app.Post("/simple-tasks", handlers.CreateSimpleTask)
	app.Get("/simple-tasks/:id", handlers.GetSimpleTask)
	app.Put("/simple-tasks/:id", handlers.UpdateSimpleTask)
	app.Delete("/simple-tasks/:id", handlers.DeleteSimpleTask)
}
