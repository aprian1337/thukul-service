package pockets

import (
	"context"
	"time"
)

type Domain struct {
	ID        int
	UserId    int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Usecase interface {
	GetList(ctx context.Context, id string) ([]Domain, error)
	GetById(ctx context.Context, userId int, pocketId int) (Domain, error)
	GetTotalByActivities(ctx context.Context, userId int, pocketId int, kind string) (int64, error)
	Create(ctx context.Context, domain Domain) (Domain, error)
	Update(ctx context.Context, domain Domain, userId int, pocketId int) (Domain, error)
	Delete(ctx context.Context, userId int, pocketId int) error
}

type Repository interface {
	PocketsGetList(ctx context.Context, id int) ([]Domain, error)
	PocketsGetById(ctx context.Context, userId int, pocketId int) (Domain, error)
	PocketsCreate(ctx context.Context, domain Domain) (Domain, error)
	PocketsUpdate(ctx context.Context, domain Domain, userId int, pocketId int) (Domain, error)
	PocketsDelete(ctx context.Context, userId int, pocketId int) (int64, error)
}
