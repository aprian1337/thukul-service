package salaries

import (
	"gorm.io/gorm"
	"time"
)

type Db struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Minimal   float64        `json:"minimal"`
	Maximal   float64        `json:"maximal"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (Db) TableName() string {
	return "salaries"
}
