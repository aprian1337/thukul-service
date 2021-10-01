package cryptos

import (
	"aprian1337/thukul-service/business/cryptos"
	"context"
	"gorm.io/gorm"
)

type PostgresCryptosRepository struct {
	ConnPostgres *gorm.DB
}

func NewPostgresCryptosRepository(conn *gorm.DB) *PostgresCryptosRepository {
	return &PostgresCryptosRepository{
		ConnPostgres: conn,
	}
}

func (repo *PostgresCryptosRepository) Update(ctx context.Context, domain cryptos.Domain) (cryptos.Domain, error) {
	model := Cryptos{}
	data := FromDomain(domain)
	err := repo.ConnPostgres.Model(&model).Where("user_id=? AND coin_id=?", data.UserId, data.CoinId)
	if err.Error != nil {
		return cryptos.Domain{}, err.Error
	}
	return model.ToDomain(), nil
}

func (repo *PostgresCryptosRepository) Create(ctx context.Context, domain cryptos.Domain) (cryptos.Domain, error) {
	model := Cryptos{}
	data := FromDomain(domain)
	err := repo.ConnPostgres.Create(&data)
	if err.Error != nil {
		return cryptos.Domain{}, err.Error
	}
	return model.ToDomain(), nil
}

func (repo *PostgresCryptosRepository) GetDetail(ctx context.Context, userId int, coinId int) (cryptos.Domain, error) {
	model := Cryptos{}
	err := repo.ConnPostgres.Model(&model).Where("user_id=? AND coin_id=?", userId, coinId)
	if err.Error != nil {
		return cryptos.Domain{}, err.Error
	}
	return model.ToDomain(), nil
}

func (repo *PostgresCryptosRepository) GetLastQty(ctx context.Context, userId int, coinId int) float64 {
	model := Cryptos{}
	err := repo.ConnPostgres.Model(&model).Where("user_id=? AND coin_id=?", userId, coinId)
	if err.Error != nil {
		return 0
	}
	return model.Qty
}
