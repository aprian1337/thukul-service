package pockets

import (
	"context"
	"time"
)

type Domain struct {
	ID           int
	UserId       int
	Name         string
	TotalNominal float64
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type Usecase interface {
	GetList(ctx context.Context, id string) ([]Domain, error)
	GetById(ctx context.Context, id int) (Domain, error)
	Create(ctx context.Context, domain Domain) (Domain, error)
	Update(ctx context.Context, id int, domain Domain) (Domain, error)
	Delete(ctx context.Context, id int) error
}

type Repository interface {
	GetList(ctx context.Context, id int) ([]Domain, error)
	GetById(ctx context.Context, id int) (Domain, error)
	Create(ctx context.Context, domain Domain) (Domain, error)
	Update(ctx context.Context, domain Domain) (Domain, error)
	Delete(ctx context.Context, id int) error
}
