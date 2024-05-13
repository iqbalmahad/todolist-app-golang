package models

import (
	"time"
)

type Todo struct {
	ID        uint
	Task      string
	Completed bool
	UserID    string `gorm:"column:user_id"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
