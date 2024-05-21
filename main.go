package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"errors"
)




type user struct{
ID int `json:"id"`
Username string `json:"username"
Password string `json:"password"
Email string `json:"email"
}

type task struct{
	ID 	int `json:"id"`
	User_id int  `json:"user_id"`
	Title  string `json:"title"`
	Description string `json:"description"`
	Status 	string `json:"status"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`

}
var users = []user{
	{
		"id": 1,
		"username": "johndoe",
		"password": "secret123",
		"email": "john.doe@example.com"
	  },
	  {
		"id": 2,
		"username": "janedoe",
		"password": "password123",
		"email": "jane.doe@example.com"
	  },
	  {
		"id": 3,
		"username": "alicewonderland",
		"password": "wonderland123",
		"email": "alice@wonderland.com"
	  }
	
}

var tasks = []task{
	{
		"id": 1,
		"user_id": 1,
		"title": "Finish coding project",
		"description": "Complete all functionalities and unit tests for the API",
		"status": "pending",
		"created_at": "2024-05-20T00:00:00Z",
		"updated_at": "2024-05-20T00:00:00Z"
	  },
	  {
		"id": 2,
		"user_id": 2,
		"title": "Grocery shopping",
		"description": "Buy milk, bread, and vegetables",
		"status": "completed",
		"created_at": "2024-05-19T18:00:00Z",
		"updated_at": "2024-05-19T20:00:00Z"
	  },
	  {
		"id": 3,
		"user_id": 3,
		"title": "Read a book",
		"description": "Finish reading the first chapter of 'Pride and Prejudice'",
		"status": "in_progress",
		"created_at": "2024-05-20T10:00:00Z",
		"updated_at": "2024-05-20T12:00:00Z"
	  }

}

func getTasks(c *gin.Context){
	c.IndentedJson(http.StatusOK, tasks)

}


func main() {
	router := gin.Default()
	router.Get("/tasks", getTasks)
	router.Run("localhost:8080")

}

