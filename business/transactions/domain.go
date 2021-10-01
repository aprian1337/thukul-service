package transactions

import (
	"context"
	"time"
)

type Domain struct {
	Id                string
	UserId            int
	CoinId            int
	Qty               float64
	Status            int
	DatetimeRequest   time.Time
	DatetimeVerify    time.Time
	DatetimeCompleted time.Time
}

type Usecase interface {
	Create(ctx context.Context, domain Domain) (Domain, error)
	UpdaterVerify(ctx context.Context, transactionId string) (Domain, error)
	UpdaterCompleted(ctx context.Context, transactionId string, status string) (Domain, error)
}

type Repository interface {
	Create(ctx context.Context, domain Domain) (Domain, error)
	UpdaterVerify(ctx context.Context, transactionId string) (Domain, error)
	UpdaterCompleted(ctx context.Context, transactionId string, status int) (Domain, error)
}
