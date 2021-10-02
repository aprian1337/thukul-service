package postgres

import (
	"aprian1337/thukul-service/business/coins"
	coins2 "aprian1337/thukul-service/repository/databases/records"
	"context"
	"gorm.io/gorm"
)

type CoinsRepository struct {
	conn *gorm.DB
}

func NewPostgresCoinsRepository(conn *gorm.DB) *CoinsRepository {
	return &CoinsRepository{
		conn: conn,
	}
}

func (repo *CoinsRepository) CoinsGetSymbol(ctx context.Context, symbol string) (coins.Domain, int64, error) {
	var data coins2.Coins
	err := repo.conn.First(&data, "symbol=?", symbol)
	return data.CoinsToDomain(), err.RowsAffected, nil
}

func (repo *CoinsRepository) CoinsCreateSymbol(ctx context.Context, domain coins.Domain) (coins.Domain, error) {
	data := coins2.CoinsFromDomain(domain)
	err := repo.conn.Create(&data)
	if err.Error != nil {
		return coins.Domain{}, err.Error
	}
	return data.CoinsToDomain(), nil
}
