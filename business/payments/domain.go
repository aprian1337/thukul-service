package payments

import (
	"aprian1337/thukul-service/business/transactions"
	"aprian1337/thukul-service/business/wallet_histories"
	"aprian1337/thukul-service/business/wallets"
	"context"
	"github.com/google/uuid"
	"time"
)

type Domain struct {
	UserId  int
	Kind    string
	Coin    string
	Qty     float64
	Price   float64
	Nominal float64
}

type Usecase interface {
	BuyCoin(ctx context.Context, domain Domain) error
	SellCoin(ctx context.Context, domain Domain) error
	TopUp(ctx context.Context, domain Domain) (wallets.Domain, error)
	Confirm(ctx context.Context, encode string, encrypt string) (wallets.Domain, error)
}

func (d *Domain) ToTransactionDomain(coinId int, kind string) transactions.Domain {
	return transactions.Domain{
		UserId: d.UserId,
		Price:  d.Price,
		Kind:   kind,
		CoinId: coinId,
		Qty:    d.Qty,
	}
}

func ToWalletHistoriesDomain(walletId int, transactionId string, nominal float64) wallet_histories.Domain {
	transactionUuid, _ := uuid.Parse(transactionId)
	return wallet_histories.Domain{
		WalletId:      walletId,
		TransactionId: &transactionUuid,
		Nominal:       nominal,
		CreatedAt:     time.Time{},
	}
}
