package cryptos

import (
	"context"
	"time"
)

type Domain struct {
	ID        int
	UserId    int
	CoinId    int
	Symbol    string
	Qty       float64
	BuyQty    float64
	SellQty   float64
	Coin      Coin
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Coin struct {
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
}

type Usecase interface {
	CryptosGetByUser(ctx context.Context, userId int) ([]Domain, error)
	CryptosGetDetail(ctx context.Context, userId int, coinId int) (Domain, error)
	UpdateBuyCoin(ctx context.Context, domain Domain) (Domain, error)
	UpdateSellCoin(ctx context.Context, domain Domain) (Domain, error)
}

type Repository interface {
	CryptosGetByUser(ctx context.Context, userId int) ([]Domain, error)
	CryptosGetDetail(ctx context.Context, userId int, coinId int) (Domain, error)
	CryptosCreate(ctx context.Context, domain Domain) (Domain, error)
	CryptosUpdate(ctx context.Context, domain Domain) (Domain, error)
}
