package requests

import (
	"aprian1337/thukul-service/business/users"
	"time"
)

type UserRegister struct {
	SalaryId int    `json:"salary_id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Gender   string `json:"gender"`
	Birthday string `json:"birthday"`
	Address  string `json:"address"`
	Company  string `json:"company"`
}

func (ur *UserRegister) ToDomain() *users.Domain {
	birthday, errTime := time.Parse("2006-01-02", ur.Birthday)
	if errTime != nil {
		panic(errTime)
	}
	return &users.Domain{
		SalaryId: ur.SalaryId,
		Name:     ur.Name,
		Password: ur.Password,
		Email:    ur.Email,
		Phone:    ur.Phone,
		Gender:   ur.Gender,
		Birthday: birthday,
		Address:  ur.Address,
		Company:  ur.Company,
	}
}
