package models

import (
	"gorm.io/gorm"
	"time"
)

type Salaries struct {
	ID        uint    `gorm:"primaryKey"`
	Minimal   float64 `json:"minimal"`
	Maximal   float64 `json:"maximal"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
