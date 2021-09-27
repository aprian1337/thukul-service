package wishlists

import (
	"context"
	"time"
)

type Domain struct {
	ID          int
	Name        string
	UserId      int
	Nominal     float64
	TargetDate  string
	Priority    string
	Note        string
	IsDone      int
	PicUrl      string
	WishlistUrl string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Usecase interface {
	GetList(ctx context.Context, userId int) ([]Domain, error)
	GetById(ctx context.Context, userId int, wishlistId int) (Domain, error)
	Create(ctx context.Context, domain Domain, userId int) (Domain, error)
	Update(ctx context.Context, domain Domain, userId int, wishlistId int) (Domain, error)
	Delete(ctx context.Context, userId int, wishlistId int) error
}

type Repository interface {
	GetList(ctx context.Context, userId int) ([]Domain, error)
	GetById(ctx context.Context, userId int, wishlistId int) (Domain, error)
	Create(ctx context.Context, domain Domain, userId int) (Domain, error)
	Update(ctx context.Context, domain Domain, userId int, wishlistId int) (Domain, error)
	Delete(ctx context.Context, userId int, wishlistId int) (int64, error)
}
