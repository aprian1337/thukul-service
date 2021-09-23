package users

import (
	"aprian1337/thukul-service/zmvc/models/salaries"
	"time"
)

type Db struct {
	ID        uint        `gorm:"primaryKey" json:"id"`
	SalaryId  int         `json:"salary_id" validate:"numeric"`
	SalaryFk  salaries.Db `gorm:"foreignKey:SalaryId"`
	Name      string      `json:"name"`
	Password  string      `json:"password"`
	IsAdmin   int         `json:"is_admin" gorm:"type:smallint; default:0" validate:"numeric"`
	Email     string      `json:"email"`
	Phone     string      `json:"phone" gorm:"size:18"`
	Gender    string      `json:"gender" gorm:"size:8"`
	Birthday  time.Time   `json:"birthday" gorm:"type:date"`
	Address   string      `json:"address" gorm:"type:text"`
	Company   string      `json:"company"`
	IsValid   int         `json:"is_valid" gorm:"type:smallint; default:0"`
	CreatedAt time.Time   `gorm:"autoCreateTime"`
	UpdatedAt time.Time   `gorm:"autoUpdateTime"`
	DeletedAt time.Time   `gorm:"index"`
}

func (Db) TableName() string {
	return "users"
}
