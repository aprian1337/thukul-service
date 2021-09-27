package coins

import (
	"aprian1337/thukul-service/business/coins"
	"context"
	"gorm.io/gorm"
)

type PostgresCoinsRepository struct {
	conn *gorm.DB
}

func NewPostgresCoinsRepository(conn *gorm.DB) *PostgresCoinsRepository {
	return &PostgresCoinsRepository{
		conn: conn,
	}
}

func (repo *PostgresCoinsRepository) GetSymbol(ctx context.Context, symbol string) (coins.Domain, int64, error) {
	var data Coins
	err := repo.conn.First(&data, "symbol=?", symbol)
	return data.ToDomain(), err.RowsAffected, nil
}

func (repo *PostgresCoinsRepository) CreateSymbol(ctx context.Context, domain coins.Domain) (coins.Domain, error) {
	data := FromDomain(domain)
	err := repo.conn.Create(&data)
	if err.Error != nil {
		return coins.Domain{}, err.Error
	}
	return data.ToDomain(), nil
}
