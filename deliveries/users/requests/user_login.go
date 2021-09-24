package requests

import (
	"aprian1337/thukul-service/business/users"
	"time"
)

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func ToDomain(login UserLogin) users.Domain {
	return users.Domain{
		ID:        0,
		SalaryId:  0,
		Name:      "",
		Password:  login.Password,
		IsAdmin:   0,
		Email:     login.Password,
		Phone:     "",
		Gender:    "",
		Birthday:  time.Time{},
		Address:   "",
		Company:   "",
		IsValid:   0,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}
}
