package coinmarket

import (
	"context"
)

type Domain struct {
	Symbol string
	Name   string
}

type Repository interface {
	GetBySymbol(ctx context.Context, symbol string) (Domain, error)
	GetPrice(ctx context.Context, symbol string, amount float64) (float64, error)
}
