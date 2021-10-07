package responses

import (
	"aprian1337/thukul-service/business/users"
)

type UsersResponse struct {
	Id       uint        `json:"id"`
	Salary   interface{} `json:"salary,omitempty"`
	Name     string      `json:"name"`
	IsAdmin  int         `json:"is_admin" validate:"numeric"`
	Email    string      `json:"email"`
	Phone    string      `json:"phone"`
	Gender   string      `json:"gender"`
	Birthday string      `json:"birthday"`
	Address  string      `json:"address"`
	Company  string      `json:"company"`
	IsValid  int         `json:"is_valid"`
	Wallets  interface{} `json:"wallets,omitempty"`
}

type LoginResponse struct {
	SessionToken string      `json:"session_token"`
	User         interface{} `json:"user"`
}

func FromUsersDomain(domain users.Domain) UsersResponse {
	return UsersResponse{
		Id:       domain.ID,
		Name:     domain.Name,
		IsAdmin:  domain.IsAdmin,
		Salary:   domain.Salary,
		Email:    domain.Email,
		Phone:    domain.Phone,
		Gender:   domain.Gender,
		Birthday: domain.Birthday,
		Address:  domain.Address,
		Company:  domain.Company,
		Wallets:  domain.Wallets,
		IsValid:  domain.IsValid,
	}
}

func FromUsersListDomain(domain []users.Domain) []UsersResponse {
	var list []UsersResponse
	for _, v := range domain {
		list = append(list, FromUsersDomain(v))
	}
	return list
}
