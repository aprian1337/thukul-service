package payments

import (
	"aprian1337/thukul-service/business/wallets"
	"context"
)

type Domain struct {
	UserId  int
	Kind    string
	Coin    string
	Qty     float64
	Nominal float64
}

type Usecase interface {
	BuyCoin(ctx context.Context, domain Domain) error
	TopUp(ctx context.Context, domain Domain) (wallets.Domain, error)
}
