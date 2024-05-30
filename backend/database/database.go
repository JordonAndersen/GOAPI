package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/group3/tasks_mgmt/models" // Import models package from your project
	"gorm.io/driver/mysql"                // Import MySQL driver for GORM
	"gorm.io/gorm"                        // Import GORM
)

type Dbinstance struct {
	Db *gorm.DB // Define a struct to hold the database connection
}

var DB Dbinstance // Create a global variable to access the database connection

func ConnectDb() {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(db:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),     // Database user
		os.Getenv("DB_PASSWORD"), // Database password
		os.Getenv("DB_NAME"),     // Database name
	)

	var db *gorm.DB
	var err error

	for {
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{}) // Try to open a connection to the database
		if err == nil {
			break
		}
		log.Println("Failed to connect to database. Retrying in 5 seconds...")
		time.Sleep(5 * time.Second) // Retry every 5 seconds if connection fails
	}

	log.Println("Connected to database")

	log.Println("Running migrations...")
	db.AutoMigrate(&models.User{}, &models.Task{}) // Run migrations to create/update database schema

	DB = Dbinstance{Db: db} // Assign the database connection to the global variable
}
