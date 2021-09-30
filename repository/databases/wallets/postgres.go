package wallets

import (
	"aprian1337/thukul-service/business/wallets"
	"context"
	"errors"
	"gorm.io/gorm"
)

type PostgresWalletsRepository struct {
	ConnPostgres *gorm.DB
}

func NewPostgresWalletsRepository(conn *gorm.DB) *PostgresWalletsRepository {
	return &PostgresWalletsRepository{
		ConnPostgres: conn,
	}
}

func (repo *PostgresWalletsRepository) GetByUserId(ctx context.Context, userId int) (wallets.Domain, error) {
	var data Wallets
	err := repo.ConnPostgres.Find(&data, "user_id=?", userId)
	if err.Error != nil {
		return wallets.Domain{}, err.Error
	}
	if err.RowsAffected == 0 {
		return wallets.Domain{}, errors.New("user id not found")
	}

	return data.ToDomain(), nil
}
func (repo *PostgresWalletsRepository) UpdateByUserId(ctx context.Context, domain wallets.Domain) (wallets.Domain, error) {
	var data Wallets
	err := repo.ConnPostgres.Model(&Wallets{}).Where("user_id = ? AND id = ?", domain.UserId, domain.Id).Update("total", domain.Total)
	if err.Error != nil {
		return wallets.Domain{}, err.Error
	}
	return data.ToDomain(), nil
}

func (repo *PostgresWalletsRepository) Create(ctx context.Context, domain wallets.Domain) error {
	data := FromDomain(domain)
	err := repo.ConnPostgres.Create(&data)
	if err.Error != nil {
		return err.Error
	}
	return nil
}
