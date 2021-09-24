package responses

import (
	"aprian1337/thukul-service/business/users"
	"aprian1337/thukul-service/repository/databases/salaries"
	"time"
)

type UsersRequest struct {
	SalaryId int `json:"salary_id" validate:"numeric"`
	SalaryFk salaries.Salaries
	Name     string    `json:"name"`
	IsAdmin  int       `json:"is_admin" validate:"numeric"`
	Email    string    `json:"email"`
	Phone    string    `json:"phone"`
	Gender   string    `json:"gender"`
	Birthday time.Time `json:"birthday"`
	Address  string    `json:"address"`
	Company  string    `json:"company"`
	IsValid  int       `json:"is_valid"`
}

func FromUsersDomain(domain users.Domain) UsersRequest {
	return UsersRequest{
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

func FromUsersListDomain(domain []users.Domain) []UsersRequest {
	list := []UsersRequest{}
	for _, v := range domain {
		list = append(list, FromUsersDomain(v))
	}
	return list
}
