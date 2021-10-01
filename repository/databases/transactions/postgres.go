package transactions

import (
	"aprian1337/thukul-service/business/transactions"
	"context"
	"gorm.io/gorm"
	"time"
)

type PostgresTransactionRepository struct {
	ConnPostgres *gorm.DB
}

func NewPostgresTransactionRepository(conn *gorm.DB) *PostgresTransactionRepository {
	return &PostgresTransactionRepository{
		ConnPostgres: conn,
	}
}

func (repo *PostgresTransactionRepository) Create(ctx context.Context, domain transactions.Domain) (transactions.Domain, error) {
	data := FromDomain(domain)
	err := repo.ConnPostgres.Create(&data)
	if err.Error != nil {
		return transactions.Domain{}, err.Error
	}
	return data.ToDomain(), nil
}

func (repo *PostgresTransactionRepository) UpdaterVerify(ctx context.Context, transactionId string) (transactions.Domain, error) {
	data := Transactions{}
	now := time.Now()
	err := repo.ConnPostgres.Model(&data).Where("id", transactionId).Update("datetime_verify", now).Update("status", 1)
	if err.Error != nil {
		return transactions.Domain{}, err.Error
	}
	return data.ToDomain(), nil
}
func (repo *PostgresTransactionRepository) UpdaterCompleted(ctx context.Context, transactionId string, status int) (transactions.Domain, error) {
	data := Transactions{}
	now := time.Now()
	err := repo.ConnPostgres.Model(&data).Where("id", transactionId).Update("datetime_completed", now).Update("status", status)
	if err.Error != nil {
		return transactions.Domain{}, err.Error
	}
	return data.ToDomain(), nil
}
