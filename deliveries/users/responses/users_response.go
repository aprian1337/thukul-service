package responses

import (
	"aprian1337/thukul-service/business/users"
)

type UsersResponse struct {
	Id       uint   `json:"id"`
	Salary   Salary `json:"salary"`
	Name     string `json:"name"`
	IsAdmin  int    `json:"is_admin" validate:"numeric"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Gender   string `json:"gender"`
	Birthday string `json:"birthday"`
	Address  string `json:"address"`
	Company  string `json:"company"`
	IsValid  int    `json:"is_valid"`
}

type Salary struct {
	ID      uint    `json:"id"`
	Minimal float64 `json:"minimal"`
	Maximal float64 `json:"maximal"`
}

type LoginResponse struct {
	SessionToken string
	User         interface{}
}

func FromUsersDomain(domain users.Domain) UsersResponse {
	return UsersResponse{
		Id:      domain.ID,
		Name:    domain.Name,
		IsAdmin: domain.IsAdmin,
		Salary: Salary(struct {
			ID      uint
			Minimal float64
			Maximal float64
		}{ID: domain.Salary.ID, Minimal: domain.Salary.Minimal, Maximal: domain.Salary.Maximal}),
		Email:    domain.Email,
		Phone:    domain.Phone,
		Gender:   domain.Gender,
		Birthday: domain.Birthday,
		Address:  domain.Address,
		Company:  domain.Company,
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

func FromUsersDomainToLogin(domain users.Domain, token string) LoginResponse {
	response := UsersResponse{
		Id:   domain.ID,
		Name: domain.Name,
		Salary: Salary(struct {
			ID      uint
			Minimal float64
			Maximal float64
		}{ID: domain.Salary.ID, Minimal: domain.Salary.Minimal, Maximal: domain.Salary.Maximal}),
		IsAdmin:  domain.IsAdmin,
		Email:    domain.Email,
		Phone:    domain.Phone,
		Gender:   domain.Gender,
		Birthday: domain.Birthday,
		Address:  domain.Address,
		Company:  domain.Company,
		IsValid:  domain.IsValid,
	}
	loginResponse := LoginResponse{}
	loginResponse.SessionToken = token
	loginResponse.User = response
	return loginResponse
}
