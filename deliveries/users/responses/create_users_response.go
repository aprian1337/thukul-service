package responses

import "aprian1337/thukul-service/business/users"

type UsersCreateResponse struct {
	Id       uint   `json:"id"`
	SalaryId int    `json:"salary_id"`
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

func FromDomainToCreateResponse(domain users.Domain) UsersCreateResponse {
	return UsersCreateResponse{
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
