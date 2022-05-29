package graphql

import (
	"aprian1337/thukul-service/app/middlewares"
	"aprian1337/thukul-service/business/users"
	"aprian1337/thukul-service/deliveries/users/requests"
	"aprian1337/thukul-service/deliveries/users/responses"
	"context"
	"fmt"
	"github.com/graphql-go/graphql"
)

type Resolver struct{
	usersUsecase users.UserUsecase
}

func (ctrl *Resolver) GetUsers(params graphql.ResolveParams) (interface{}, error) {
	header := params.Info.RootValue.(map[string]interface{})
	if header["token"]==""{
		return nil, fmt.Errorf("you're not have authorized for this action")
	}
	_, err := middlewares.GetClaimsAdminId(header["token"].(string))

	ctx := context.Background()
	res, err := ctrl.usersUsecase.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (ctrl *Resolver) GetUserDetail(params graphql.ResolveParams) (interface{}, error) {
	header := params.Info.RootValue.(map[string]interface{})
	if header["token"]==""{
		return nil, fmt.Errorf("you're not have authorized for this action")
	}
	id, err := middlewares.GetClaimsUserId(header["token"].(string))
	if err != nil {
		return nil, err
	}
	ctx := context.Background()
	res, err := ctrl.usersUsecase.GetById(ctx, int(id))
	if err != nil {
		return nil, err
	}
	return responses.FromUsersDomain(res), nil
}

func (ctrl *Resolver) LoginUser(params graphql.ResolveParams) (interface{}, error) {
	var(
		email string
		password string
		ok bool
	)
	if email, ok = params.Args["email"].(string); !ok || email == "" {
		return nil, fmt.Errorf("email is required and must be string")
	}
	if password, ok = params.Args["password"].(string); !ok || password == "" {
		return nil, fmt.Errorf("password is required and must be string")
	}
	ctx := context.Background()
	res, token, err := ctrl.usersUsecase.Login(ctx, email, password)
	if err != nil {
		return nil, err
	}
	return responses.FromUsersDomainToLogin(res, token), nil
}

func (ctrl *Resolver) RegisterUser(params graphql.ResolveParams) (interface{}, error) {
	var(
		salaryId int
		name string
		password string
		email string
		phone string
		gender string
		birthday string
		address string
		company string
		ok bool
	)
	if salaryId, ok = params.Args["salaryId"].(int); !ok || salaryId == 0 {
		return nil, fmt.Errorf("salaryId is required and must be int")
	}
	if name, ok = params.Args["name"].(string); !ok || name == "" {
		return nil, fmt.Errorf("name is required and must be string")
	}
	if password, ok = params.Args["password"].(string); !ok || password == "" {
		return nil, fmt.Errorf("password is required and must be string")
	}
	if phone, ok = params.Args["phone"].(string); !ok || phone == "" {
		return nil, fmt.Errorf("phone is required and must be string")
	}
	if email, ok = params.Args["email"].(string); !ok || email == "" {
		return nil, fmt.Errorf("email is required and must be string")
	}
	if gender, ok = params.Args["gender"].(string); !ok || gender == "" {
		return nil, fmt.Errorf("gender is required and must be string")
	}
	if gender != "Male" && gender != "Female"{
		return nil, fmt.Errorf("gender must be \"Female\" or \"Male\"")
	}
	if birthday, ok = params.Args["birthday"].(string); !ok || birthday == "" {
		return nil, fmt.Errorf("birthday is required and must be string")
	}
	if address, ok = params.Args["address"].(string); !ok || address == "" {
		return nil, fmt.Errorf("address is required and must be string")
	}
	if company, ok = params.Args["company"].(string); !ok || company == "" {
		return nil, fmt.Errorf("company is required and must be string")
	}
	ctx := context.Background()
	request := requests.UserRegister{
		SalaryId: salaryId,
		Name:     name,
		Password: password,
		Email:    email,
		Phone:    phone,
		IsAdmin:  0,
		Gender:   gender,
		Birthday: birthday,
		Address:  address,
		Company:  company,
	}
	fmt.Println("EMAILLL : ", email)
	res, err := ctrl.usersUsecase.Create(ctx, request.ToDomain())
	if err != nil {
		return nil, err
	}
	return responses.FromDomainToCreateResponse(res), nil
}

func NewUsersResolver(usersUsecase users.UserUsecase) *Resolver{
	return &Resolver{
		usersUsecase: usersUsecase,
	}
}