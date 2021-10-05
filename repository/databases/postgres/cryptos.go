package postgres

import (
	"aprian1337/thukul-service/business/cryptos"
	"aprian1337/thukul-service/repository/databases/records"
	"context"
	"fmt"
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
func (repo *CryptosRepository) CryptosGetByUser(ctx context.Context, userId int) ([]cryptos.Domain, error) {
	var model []records.Cryptos
	err := repo.ConnPostgres.Joins("Coin").Find(&model, "user_id=?", userId)
	if err.Error != nil {
		return []cryptos.Domain{}, err.Error
	}
	return records.CryptosToListDomain(model), nil
}

func (repo *CryptosRepository) CryptosUpdate(ctx context.Context, domain cryptos.Domain) (cryptos.Domain, error) {
	model := records.Cryptos{}
	err := repo.ConnPostgres.First(&model, "user_id=? AND coin_id=?", domain.UserId, domain.CoinId)
	model.Qty = domain.Qty
	repo.ConnPostgres.Save(&model)
	if err.Error != nil {
		return cryptos.Domain{}, err.Error
	}
	return model.CryptosToDomain(), nil
}

func (repo *CryptosRepository) CryptosCreate(ctx context.Context, domain cryptos.Domain) (cryptos.Domain, error) {
	model := records.Cryptos{}
	data := records.CryptosFromDomain(domain)
	fmt.Println(data)
	err := repo.ConnPostgres.Create(&data)
	if err.Error != nil {
		return cryptos.Domain{}, err.Error
	}
	return model.CryptosToDomain(), nil
}

func (repo *CryptosRepository) CryptosGetDetail(ctx context.Context, userId int, coinId int) (cryptos.Domain, error) {
	model := records.Cryptos{}
	err := repo.ConnPostgres.Joins("Coin").First(&model, "user_id=? AND coin_id=?", userId, coinId)
	if err.Error != nil {
		return cryptos.Domain{}, err.Error
	}
	return model.CryptosToDomain(), nil
}

func (repo *CryptosRepository) CryptosGetLastQty(ctx context.Context, userId int, coinId int) float64 {
	model := records.Cryptos{}
	err := repo.ConnPostgres.Model(&model).Where("user_id=? AND coin_id=?", userId, coinId)
	if err.Error != nil {
		return 0
	}
	return model.Qty
}
