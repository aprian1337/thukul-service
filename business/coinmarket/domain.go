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
}
