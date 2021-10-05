package records

import (
	"aprian1337/thukul-service/business/wallet_histories"
	wallets_domain "aprian1337/thukul-service/business/wallets"
	"github.com/google/uuid"
	"time"
)

type WalletHistories struct {
	ID            int `gorm:"primaryKey"`
	WalletId      int
	Wallet        Wallets `gorm:"foreignKey:wallet_id"`
	TransactionId *uuid.UUID
	Transaction   Transactions `gorm:"foreignKey:transaction_id"`
	Nominal       float64
	CreatedAt     time.Time `gorm:"autoCreateTime"`
}

func (data *WalletHistories) WalletHistoriesToDomain() wallets_domain.Domain {
	return wallets_domain.Domain{
		Id:     0,
		UserId: 0,
		Total:  0,
	}
}

func WalletHistoriesFromDomain(domain wallet_histories.Domain) WalletHistories {
	return WalletHistories{
		WalletId:      domain.WalletId,
		Nominal:       domain.Nominal,
		TransactionId: domain.TransactionId,
	}
}
