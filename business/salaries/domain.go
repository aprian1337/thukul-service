package salaries

import (
	"context"
	"time"
)

type Domain struct {
	ID        uint
	Minimal   float64
	Maximal   float64
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Usecase interface {
	GetList(ctx context.Context) ([]Domain, error)
	GetById(ctx context.Context, id uint) (Domain, error)
	Create(ctx context.Context, domain Domain) (Domain, error)
	Update(ctx context.Context, id uint, domain Domain) (Domain, error)
	Delete(ctx context.Context, id uint) error
}

type Repository interface {
	SalariesGetList(ctx context.Context) ([]Domain, error)
	SalariesGetById(ctx context.Context, id uint) (Domain, error)
	SalariesCreate(ctx context.Context, domain Domain) (Domain, error)
	SalariesUpdate(ctx context.Context, domain Domain) (Domain, error)
	SalariesDelete(ctx context.Context, id uint) (int64, error)
}
