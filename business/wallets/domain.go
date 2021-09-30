package wallets

import (
	"aprian1337/thukul-service/business/wallet_histories"
	"context"
	"github.com/google/uuid"
)

type Domain struct {
	Id                 int
	UserId             int
	Total              float64
	NominalTransaction float64
	Kind               string
	TransactionId      *uuid.UUID
	CoinId             int
}

type Usecase interface {
	GetByUserId(ctx context.Context, userId int) (Domain, error)
	UpdateByUserId(ctx context.Context, domain Domain) (Domain, error)
	Create(ctx context.Context, domain Domain) error
}

type Repository interface {
	GetByUserId(ctx context.Context, userId int) (Domain, error)
	UpdateByUserId(ctx context.Context, domain Domain) (Domain, error)
	Create(ctx context.Context, domain Domain) error
}

func (d *Domain) ToHistoryDomain() wallet_histories.Domain {
	return wallet_histories.Domain{
		WalletId:      d.Id,
		TransactionId: d.TransactionId,
		Type:          d.Kind,
		Nominal:       d.NominalTransaction,
	}
}
