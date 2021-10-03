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
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Usecase interface {
	CryptosGetDetail(ctx context.Context, userId int, coinId int) (Domain, error)
	UpdateBuyCoin(ctx context.Context, domain Domain) (Domain, error)
	UpdateSellCoin(ctx context.Context, domain Domain) (Domain, error)
}

type Repository interface {
	CryptosGetDetail(ctx context.Context, userId int, coinId int) (Domain, error)
	CryptosCreate(ctx context.Context, domain Domain) (Domain, error)
	CryptosUpdate(ctx context.Context, domain Domain) (Domain, error)
}
