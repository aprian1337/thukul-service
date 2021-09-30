package wallets

import (
	"aprian1337/thukul-service/business/wallet_histories"
	wallets_domain "aprian1337/thukul-service/business/wallets"
	"aprian1337/thukul-service/repository/databases/transactions"
	"aprian1337/thukul-service/repository/databases/wallets"
	"time"
)

type WalletHistories struct {
	ID            int `gorm:"primaryKey"`
	WalletId      int
	Wallet        wallets.Wallets `gorm:"foreignKey:wallet_id"`
	TransactionId int
	Transaction   transactions.Transactions `gorm:"foreignKey:transaction_id"`
	Type          string
	Nominal       float64
	CreatedAt     time.Time `gorm:"autoCreateTime"`
}

func (data *WalletHistories) ToDomain() wallets_domain.Domain {
	return wallets_domain.Domain{
		Id:     0,
		UserId: 0,
		Total:  0,
	}
}

func FromDomain(domain wallet_histories.Domain) WalletHistories {
	return WalletHistories{
		WalletId:      domain.WalletId,
		TransactionId: domain.TransactionId,
		Type:          domain.Type,
		Nominal:       domain.Nominal,
	}
}
