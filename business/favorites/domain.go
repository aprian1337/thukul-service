package favorites

import (
	"aprian1337/thukul-service/business/coins"
	"context"
	"time"
)

type Domain struct {
	ID     int
	UserId int
	CoinId int
	Coins  struct {
		CoinSymbol string
		CoinName   string
	}
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Usecase interface {
	GetList(ctx context.Context, userId int) ([]Domain, error)
	GetById(ctx context.Context, userId int, favoriteId int) (Domain, error)
	Create(ctx context.Context, domain Domain, userId int) (Domain, error)
	Delete(ctx context.Context, userId int, favoriteId int) error
}

type Repository interface {
	Check(ctx context.Context, userId int, coinId int) (int64, error)
	GetList(ctx context.Context, userId int) ([]Domain, error)
	GetById(ctx context.Context, userId int, favoriteId int) (Domain, error)
	Create(ctx context.Context, domain Domain) (Domain, error)
	Delete(ctx context.Context, userId int, favoriteId int) (int64, error)
}

func (d *Domain) AddCoins(symbol coins.Domain) Domain {
	return Domain{
		ID:     d.ID,
		UserId: d.UserId,
		CoinId: d.CoinId,
		Coins: struct {
			CoinSymbol string
			CoinName   string
		}{
			CoinSymbol: symbol.Symbol,
			CoinName:   symbol.Name,
		},
		CreatedAt: d.CreatedAt,
		UpdatedAt: d.UpdatedAt,
	}
}

func AddCoinsList(data []Domain, symbol coins.Domain) []Domain {
	res := []Domain{}
	for _, v := range data {
		res = append(res, v.AddCoins(symbol))
	}
	return res
}
