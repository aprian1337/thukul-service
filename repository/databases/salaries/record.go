package salaries

import (
	"gorm.io/gorm"
	"time"
)

type Salaries struct {
	ID        uint `gorm:"primaryKey"`
	Minimal   float64
	Maximal   float64
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
