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
	TransactionsCreate(ctx context.Context, domain Domain) (Domain, error)
	TransactionsUpdaterVerify(ctx context.Context, transactionId string) (Domain, error)
	TransactionsUpdaterCompleted(ctx context.Context, transactionId string, status int) (Domain, error)
}

type Repository interface {
	TransactionsCreate(ctx context.Context, domain Domain) (Domain, error)
	TransactionsUpdaterVerify(ctx context.Context, transactionId string) (Domain, error)
	TransactionsUpdaterCompleted(ctx context.Context, transactionId string, status int) (Domain, error)
}
