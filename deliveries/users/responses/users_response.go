package responses

import (
	"aprian1337/thukul-service/business/users"
	"aprian1337/thukul-service/repository/databases/salaries"
	"time"
)

//DELETE
type LoginResponse struct {
	Message string `json:"message"`
	Login   int    `json:"login"`
	Data    interface{}
}

type UsersResponse struct {
	Id       uint `json:"id"`
	SalaryId int  `json:"salary_id" validate:"numeric"`
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
	list := []UsersResponse{}
	for _, v := range domain {
		list = append(list, FromUsersDomain(v))
	}
	return list
}

func ToLoginResponse(domain users.Domain) LoginResponse {
	login := LoginResponse{}
	if domain.ID > 0 {
		login.Login = 1
		login.Message = "Login Successfuly"
		login.Data = UsersResponse{
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
	} else {
		login.Login = 0
		login.Message = "Login Failed"
		login.Data = nil
	}
	return login
}
