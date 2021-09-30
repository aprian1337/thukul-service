package wallets

import (
	"aprian1337/thukul-service/business/wallet_histories"
	"context"
	"gorm.io/gorm"
)

type PostgresWalletHistoriesRepository struct {
	ConnPostgres *gorm.DB
}

func NewPostgresWalletHistoriesRepository(conn *gorm.DB) *PostgresWalletHistoriesRepository {
	return &PostgresWalletHistoriesRepository{
		ConnPostgres: conn,
	}
}
func (repo *PostgresWalletHistoriesRepository) Create(ctx context.Context, domain wallet_histories.Domain) error {
	data := FromDomain(domain)
	err := repo.ConnPostgres.Create(&data)
	if err.Error != nil {
		return err.Error
	}
	return nil
}
