package wallets

import (
	"context"
)

type Domain struct {
	Id     int
	UserId int
	Total  float64
}

type Usecase interface {
	GetByUserId(ctx context.Context, userId int) (Domain, error)
	UpdateByUserId(ctx context.Context, domain Domain) (Domain, error)
	Create(ctx context.Context, domain Domain) error
}

type Repository interface {
	GetByUserId(ctx context.Context, userId int) (Domain, error)
	UpdateByUserId(ctx context.Context, domain Domain) (Domain, error)
	Create(ctx context.Context, domain Domain) error
}
