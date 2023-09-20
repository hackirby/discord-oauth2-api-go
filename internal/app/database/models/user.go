package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserID string `json:"user_id"`
}
