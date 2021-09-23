package users

import (
	"aprian1337/thukul-service/models/salaries"
)

type Request struct {
	SalaryId int         `json:"salary_id" validate:"numeric"`
	SalaryFk salaries.Db `gorm:"foreignKey:SalaryId"`
	Name     string      `json:"name"`
	Password string      `json:"password"`
	IsAdmin  int         `json:"is_admin" validate:"numeric"`
	Email    string      `json:"email"`
	Phone    string      `json:"phone"`
	Gender   string      `json:"gender"`
	Birthday string      `json:"birthday"`
	Address  string      `json:"address"`
	Company  string      `json:"company"`
	IsValid  int         `json:"is_valid"`
}
