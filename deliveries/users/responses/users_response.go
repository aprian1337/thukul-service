package responses

import (
	"aprian1337/thukul-service/business/users"
	"aprian1337/thukul-service/repository/databases/records"
)

type UsersResponse struct {
	Id       uint             `json:"id"`
	SalaryId int              `json:"salary_id" validate:"numeric"`
	Salary   records.Salaries `json:"salary"`
	Name     string           `json:"name"`
	IsAdmin  int              `json:"is_admin" validate:"numeric"`
	Email    string           `json:"email"`
	Phone    string           `json:"phone"`
	Gender   string           `json:"gender"`
	Birthday string           `json:"birthday"`
	Address  string           `json:"address"`
	Company  string           `json:"company"`
	IsValid  int              `json:"is_valid"`
}

type LoginResponse struct {
	SessionToken string
	User         interface{}
}

func FromUsersDomain(domain users.Domain) UsersResponse {
	return UsersResponse{
		Id:       domain.ID,
		SalaryId: domain.SalaryId,
		Name:     domain.Name,
		IsAdmin:  domain.IsAdmin,
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
		Id:       domain.ID,
		SalaryId: domain.SalaryId,
		Name:     domain.Name,
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
