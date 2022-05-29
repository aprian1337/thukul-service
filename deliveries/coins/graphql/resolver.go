package graphql

import (
	"aprian1337/thukul-service/business/coins"
	"context"
	"fmt"
	"github.com/graphql-go/graphql"
)

type Resolver struct {
	coinsUsecase coins.Usecase
}

func (ctrl *Resolver) GetBySymbol(params graphql.ResolveParams) (interface{}, error) {
	var(
		symbol string
		ok bool
	)
	if symbol, ok = params.Args["symbol"].(string); !ok || symbol == "" {
		return nil, fmt.Errorf("symbol is required and must be string typee")
	}
	ctx := context.Background()
	resp, err := ctrl.coinsUsecase.GetBySymbol(ctx, symbol)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (ctrl *Resolver) GetAllSymbol(params graphql.ResolveParams) (interface{}, error) {
	ctx := context.Background()
	resp, err := ctrl.coinsUsecase.GetAllSymbol(ctx)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func NewCoinsResolver(usecase coins.CoinUsecase) *Resolver {
	return &Resolver{
		coinsUsecase: &usecase,
	}
}
