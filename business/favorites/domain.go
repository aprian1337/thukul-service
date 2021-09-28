package favorites

import (
	"aprian1337/thukul-service/business/coins"
	coins2 "aprian1337/thukul-service/repository/databases/coins"
	"context"
	"time"
)

type Domain struct {
	ID        int
	UserId    int
	CoinId    int
	Coin      coins2.Coins
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
		Coin: coins2.Coins{
			Symbol: symbol.Symbol,
			Name:   symbol.Name,
		},
		CreatedAt: d.CreatedAt,
		UpdatedAt: d.UpdatedAt,
	}
}
