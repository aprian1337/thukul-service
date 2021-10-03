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
	GetBySymbol(ctx context.Context, symbol string) (Domain, error)
}

type Repository interface {
	CoinsGetSymbol(ctx context.Context, symbol string) (Domain, int64, error)
	CoinsCreateSymbol(ctx context.Context, domain Domain) (Domain, error)
}
