package activities

import (
	"context"
	"time"
)

type Domain struct {
	ID        int
	PocketId  int
	Name      string
	Type      string
	Nominal   float64
	Note      string
	Date      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Usecase interface {
	GetList(ctx context.Context, pocketId int) ([]Domain, error)
	GetById(ctx context.Context, pocketId int, id int) (Domain, error)
	GetTotal(ctx context.Context, userId int, pocketId int, kind string) (int64, error)
	Create(ctx context.Context, domain Domain, pocketId int) (Domain, error)
	Update(ctx context.Context, domain Domain, pocketId int, id int) (Domain, error)
	Delete(ctx context.Context, id int, pocketId int) error
}

type Repository interface {
	ActivitiesGetAll(ctx context.Context, pocketId int) ([]Domain, error)
	ActivitiesGetById(ctx context.Context, pocketId int, id int) (Domain, error)
	ActivitiesGetTotal(ctx context.Context, userId int, pocketId int, kind string) (int64, error)
	ActivitiesCreate(ctx context.Context, domain Domain, pocketId int) (Domain, error)
	ActivitiesUpdate(ctx context.Context, domain Domain, pocketId int, id int) (Domain, error)
	ActivitiesDelete(ctx context.Context, id int, pocketId int) (int64, error)
}
