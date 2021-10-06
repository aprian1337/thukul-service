package coins

import (
	"context"
	"time"
)

type Domain struct {
	Id        int
	Symbol    string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Usecase interface {
	GetAllSymbol(ctx context.Context) ([]Domain, error)
	GetBySymbol(ctx context.Context, symbol string) (Domain, error)
}

type Repository interface {
	GetAllSymbol(ctx context.Context) ([]Domain, error)
	CoinsGetSymbol(ctx context.Context, symbol string) (Domain, int64, error)
	CoinsCreateSymbol(ctx context.Context, domain Domain) (Domain, error)
}
