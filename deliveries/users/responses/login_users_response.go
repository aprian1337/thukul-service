package responses

import "aprian1337/thukul-service/business/users"

func FromUsersDomainToLogin(domain users.Domain, token string) LoginResponse {
	response := UsersResponse{
		Id:       domain.ID,
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
