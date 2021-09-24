package users

import (
	"aprian1337/thukul-service/deliveries/users/requests"
	"context"
	"time"
)

type Domain struct {
	ID        uint
	SalaryId  int
	Name      string
	Password  string
	IsAdmin   int
	Email     string
	Phone     string
	Gender    string
	Birthday  time.Time
	Address   string
	Company   string
	IsValid   int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Usecase interface {
	GetAll(ctx context.Context) ([]Domain, error)
	GetById(id uint, ctx context.Context) (Domain, error)
	Create(ctx context.Context, register requests.UserRegister) (Domain, error)
	Login(ctx context.Context, login requests.UserLogin) (Domain, error)
}

type Repository interface {
	GetAll(ctx context.Context) ([]Domain, error)
	GetById(id uint, ctx context.Context) (Domain, error)
	Create(ctx context.Context, register requests.UserRegister) (Domain, error)
	Login(ctx context.Context, login requests.UserLogin) (Domain, error)
}
