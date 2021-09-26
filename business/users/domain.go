package users

import (
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
	Birthday  string
	Address   string
	Company   string
	IsValid   int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Usecase interface {
	GetAll(ctx context.Context) ([]Domain, error)
	GetById(ctx context.Context, id uint) (Domain, error)
	Create(ctx context.Context, register *Domain) (Domain, error)
	Login(ctx context.Context, email string, password string) (Domain, string, error)
	Update(ctx context.Context, domain *Domain, id uint) (Domain, error)
	Delete(ctx context.Context, id uint) error
}

type Repository interface {
	GetAll(ctx context.Context) ([]Domain, error)
	GetById(ctx context.Context, id uint) (Domain, error)
	GetByEmail(ctx context.Context, email string) (Domain, error)
	Create(ctx context.Context, register *Domain) (Domain, error)
	Update(ctx context.Context, domain *Domain) (Domain, error)
	Delete(ctx context.Context, id uint) error
}
