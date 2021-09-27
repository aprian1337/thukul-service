package favorites

import (
	"context"
	"time"
)

type Domain struct {
	ID        int
	UserId    int
	CoinId    int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Usecase interface {
	GetList(ctx context.Context, userId int) ([]Domain, error)
	GetById(ctx context.Context, userId int, favoriteId int) (Domain, error)
	Create(ctx context.Context, domain Domain, userId int, coin string) (Domain, error)
	Delete(ctx context.Context, userId int, wishlistId int) error
}

type Repository interface {
	GetList(ctx context.Context, userId int) ([]Domain, error)
	GetById(ctx context.Context, userId int, wishlistId int) (Domain, error)
	Create(ctx context.Context, domain Domain, userId int) (Domain, error)
	Delete(ctx context.Context, userId int, wishlistId int) (int64, error)
}
