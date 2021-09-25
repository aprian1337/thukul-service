package requests

import (
	"aprian1337/thukul-service/business/users"
)

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func ToDomain(login UserLogin) users.Domain {
	return users.Domain{
		Password: login.Password,
		Email:    login.Password,
	}
}
