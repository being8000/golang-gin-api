package model

import (
	"time"

	"gorm.io/gorm"
)

// TODO db model
type TODO struct {
	gorm.Model
	ID        string    `json:""`
	CreatedAt time.Time `json:""`
	Done      bool      `json:""`
	Subject   string    `json:""`
	Note      string    `json:""`
}
