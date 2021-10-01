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
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Usecase interface {
	UpdateBuyCoin(ctx context.Context, domain Domain) (Domain, error)
	UpdateSellCoin(ctx context.Context, domain Domain) (Domain, error)
}

type Repository interface {
	GetDetail(ctx context.Context, userId int, coinId int) (Domain, error)
	Create(ctx context.Context, domain Domain) (Domain, error)
	Update(ctx context.Context, domain Domain) (Domain, error)
}
