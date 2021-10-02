package postgres

import (
	"aprian1337/thukul-service/business/wallets"
	wallets2 "aprian1337/thukul-service/repository/databases/records"
	"context"
	"errors"
	"gorm.io/gorm"
)

type WalletsRepository struct {
	ConnPostgres *gorm.DB
}

func NewPostgresWalletsRepository(conn *gorm.DB) *WalletsRepository {
	return &WalletsRepository{
		ConnPostgres: conn,
	}
}

func (repo *WalletsRepository) GetByUserId(ctx context.Context, userId int) (wallets.Domain, error) {
	var data wallets2.Wallets
	err := repo.ConnPostgres.Find(&data, "user_id=?", userId)
	if err.Error != nil {
		return wallets.Domain{}, err.Error
	}
	if err.RowsAffected == 0 {
		return wallets.Domain{}, errors.New("user id not found")
	}

	return data.WalletsToDomain(), nil
}
func (repo *WalletsRepository) UpdateByUserId(ctx context.Context, domain wallets.Domain) (wallets.Domain, error) {
	data := wallets2.WalletsFromDomain(domain)
	err := repo.ConnPostgres.Model(&data).Where("user_id=?", data.UserId).Update("total", data.Total)
	if err.Error != nil {
		return wallets.Domain{}, err.Error
	}
	return data.WalletsToDomain(), nil
}

func (repo *WalletsRepository) Create(ctx context.Context, domain wallets.Domain) error {
	data := wallets2.WalletsFromDomain(domain)
	err := repo.ConnPostgres.Create(&data)
	if err.Error != nil {
		return err.Error
	}
	return nil
}
