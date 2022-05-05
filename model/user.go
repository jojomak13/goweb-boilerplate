package model

import (
	"time"
)

type User struct {
	ID        uint `gorm:"primarykey"`
	Name      string `json:"name" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	CreatedAt time.Time
	UpdatedAt time.Time
}