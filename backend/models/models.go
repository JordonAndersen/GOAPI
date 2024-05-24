package models

import (
	"time"
)

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Username string `json:"username" gorm:"unique;not null" validate:"required,min=3"`
	Password string `json:"password" gorm:"not null" validate:"required,min=8"`
	Email    string `json:"email" validate:"required,email"`
	Tasks    []Task `json:"tasks" gorm:"foreignKey:UserID"`
}

type Task struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	UserID      uint      `json:"user_id"`
	Title       string    `json:"title" gorm:"not null" validate:"required"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
