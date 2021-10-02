package postgres

import (
	"aprian1337/thukul-service/business/transactions"
	transactions2 "aprian1337/thukul-service/repository/databases/records"
	"context"
	"gorm.io/gorm"
	"time"
)

type TransactionRepository struct {
	ConnPostgres *gorm.DB
}

func NewPostgresTransactionRepository(conn *gorm.DB) *TransactionRepository {
	return &TransactionRepository{
		ConnPostgres: conn,
	}
}

func (repo *TransactionRepository) TransactionsCreate(ctx context.Context, domain transactions.Domain) (transactions.Domain, error) {
	data := transactions2.TransactionsFromDomain(domain)
	err := repo.ConnPostgres.Create(&data)
	if err.Error != nil {
		return transactions.Domain{}, err.Error
	}
	return data.TransactionsToDomain(), nil
}

func (repo *TransactionRepository) TransactionsUpdaterVerify(ctx context.Context, transactionId string) (transactions.Domain, error) {
	data := transactions2.Transactions{}
	now := time.Now()
	err := repo.ConnPostgres.Model(&data).Where("id", transactionId).Update("datetime_verify", now).Update("status", 1)
	if err.Error != nil {
		return transactions.Domain{}, err.Error
	}
	return data.TransactionsToDomain(), nil
}
func (repo *TransactionRepository) TransactionsUpdaterCompleted(ctx context.Context, transactionId string, status int) (transactions.Domain, error) {
	data := transactions2.Transactions{}
	now := time.Now()
	err := repo.ConnPostgres.Model(&data).Where("id", transactionId).Update("datetime_completed", now).Update("status", status)
	if err.Error != nil {
		return transactions.Domain{}, err.Error
	}
	return data.TransactionsToDomain(), nil
}
