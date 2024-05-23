package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/group3/tasks_mgmt/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Dbinstance struct {
	Db *gorm.DB
}

var DB Dbinstance

func ConnectDb() {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(backend-db-1:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	var db *gorm.DB
	var err error

	for {
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err == nil {
			break
		}
		log.Println("Failed to connect to database. Retrying in 5 seconds...")
		time.Sleep(5 * time.Second)
	}

	log.Println("Connected to database")

	log.Println("Running migrations...")
	db.AutoMigrate(&models.User{}, &models.Task{})

	DB = Dbinstance{
		Db: db,
	}
}
