package requests

import (
	"aprian1337/thukul-service/business/users"
)

type UserRegister struct {
	SalaryId int    `json:"salary_id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	IsAdmin  int    `json:"is_admin"`
	Gender   string `json:"gender"`
	Birthday string `json:"birthday"`
	Address  string `json:"address"`
	Company  string `json:"company"`
}

func (ur *UserRegister) ToDomain() *users.Domain {
	return &users.Domain{
		SalaryId: ur.SalaryId,
		Name:     ur.Name,
		Password: ur.Password,
		Email:    ur.Email,
		Phone:    ur.Phone,
		Gender:   ur.Gender,
		IsAdmin:  ur.IsAdmin,
		Birthday: ur.Birthday,
		Address:  ur.Address,
		Company:  ur.Company,
	}
}
