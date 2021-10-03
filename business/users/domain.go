package users

import (
	"aprian1337/thukul-service/business/smtp"
	wallets_domain "aprian1337/thukul-service/business/wallets"
	"context"
	"time"
)

type Domain struct {
	ID        uint
	SalaryId  int
	Salary    Salary
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
	Wallets   Wallets
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Salary struct {
	ID      uint    `json:"id"`
	Minimal float64 `json:"minimal"`
	Maximal float64 `json:"maximal"`
}

type Wallets struct {
	ID    int     `json:"id"`
	Total float64 `json:"total"`
}

type Usecase interface {
	GetAll(ctx context.Context) ([]Domain, error)
	GetByIdWithWallet(ctx context.Context, id int) (Domain, error)
	GetById(ctx context.Context, id int) (Domain, error)
	Create(ctx context.Context, register *Domain) (Domain, error)
	Login(ctx context.Context, email string, password string) (Domain, string, error)
	Update(ctx context.Context, domain *Domain, id uint) (Domain, error)
	Delete(ctx context.Context, id uint) error
}

type Repository interface {
	UsersGetAll(ctx context.Context) ([]Domain, error)
	UsersGetByIdWithWallet(ctx context.Context, id int) (Domain, error)
	UsersGetById(ctx context.Context, id int) (Domain, error)
	UsersGetByEmail(ctx context.Context, email string) (Domain, error)
	UsersCreate(ctx context.Context, register *Domain) (Domain, error)
	UsersUpdate(ctx context.Context, domain *Domain) (Domain, error)
	UsersDelete(ctx context.Context, id uint) error
}

func (d *Domain) ToWalletDomain() wallets_domain.Domain {
	return wallets_domain.Domain{
		UserId: int(d.ID),
	}
}

func (d *Domain) ToSmtpDomain(subject string, message string) smtp.Domain {
	return smtp.Domain{
		MailTo:  []string{d.Email},
		Subject: subject,
		Message: message,
	}
}
