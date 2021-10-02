package postgres

import (
	"aprian1337/thukul-service/business/wallet_histories"
	"aprian1337/thukul-service/repository/databases/records"
	"context"
	"gorm.io/gorm"
)

type WalletHistoriesRepository struct {
	ConnPostgres *gorm.DB
}

func NewPostgresWalletHistoriesRepository(conn *gorm.DB) *WalletHistoriesRepository {
	return &WalletHistoriesRepository{
		ConnPostgres: conn,
	}
}
func (repo *WalletHistoriesRepository) WalletHistoriesCreate(ctx context.Context, domain wallet_histories.Domain) error {
	data := records.WalletHistoriesFromDomain(domain)
	err := repo.ConnPostgres.Create(&data)
	if err.Error != nil {
		return err.Error
	}
	return nil
}
