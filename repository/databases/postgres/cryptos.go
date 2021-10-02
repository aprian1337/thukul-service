package postgres

import (
	"aprian1337/thukul-service/business/cryptos"
	cryptos2 "aprian1337/thukul-service/repository/databases/records"
	"context"
	"gorm.io/gorm"
)

type CryptosRepository struct {
	ConnPostgres *gorm.DB
}

func NewPostgresCryptosRepository(conn *gorm.DB) *CryptosRepository {
	return &CryptosRepository{
		ConnPostgres: conn,
	}
}

func (repo *CryptosRepository) CryptosUpdate(ctx context.Context, domain cryptos.Domain) (cryptos.Domain, error) {
	model := cryptos2.Cryptos{}
	data := cryptos2.CryptosFromDomain(domain)
	err := repo.ConnPostgres.Model(&model).Where("user_id=? AND coin_id=?", data.UserId, data.CoinId)
	if err.Error != nil {
		return cryptos.Domain{}, err.Error
	}
	return model.CryptosToDomain(), nil
}

func (repo *CryptosRepository) CryptosCreate(ctx context.Context, domain cryptos.Domain) (cryptos.Domain, error) {
	model := cryptos2.Cryptos{}
	data := cryptos2.CryptosFromDomain(domain)
	err := repo.ConnPostgres.Create(&data)
	if err.Error != nil {
		return cryptos.Domain{}, err.Error
	}
	return model.CryptosToDomain(), nil
}

func (repo *CryptosRepository) CryptosGetDetail(ctx context.Context, userId int, coinId int) (cryptos.Domain, error) {
	model := cryptos2.Cryptos{}
	err := repo.ConnPostgres.Model(&model).Where("user_id=? AND coin_id=?", userId, coinId)
	if err.Error != nil {
		return cryptos.Domain{}, err.Error
	}
	return model.CryptosToDomain(), nil
}

func (repo *CryptosRepository) CryptosGetLastQty(ctx context.Context, userId int, coinId int) float64 {
	model := cryptos2.Cryptos{}
	err := repo.ConnPostgres.Model(&model).Where("user_id=? AND coin_id=?", userId, coinId)
	if err.Error != nil {
		return 0
	}
	return model.Qty
}
