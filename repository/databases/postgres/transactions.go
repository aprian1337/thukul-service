package postgres

import (
	"aprian1337/thukul-service/business/transactions"
	"aprian1337/thukul-service/repository/databases/records"
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

func (repo *TransactionRepository) TransactionsById(ctx context.Context, id string) (transactions.Domain, error) {
	model := records.Transactions{}
	err := repo.ConnPostgres.First(&model, "id=?", id)
	if err.Error != nil {
		return transactions.Domain{}, err.Error
	}
	return model.TransactionsToDomain(), nil
}

func (repo *TransactionRepository) TransactionsCreate(ctx context.Context, domain transactions.Domain) (transactions.Domain, error) {
	data := records.TransactionsFromDomain(domain)
	err := repo.ConnPostgres.Create(&data)
	if err.Error != nil {
		return transactions.Domain{}, err.Error
	}
	return data.TransactionsToDomain(), nil
}

func (repo *TransactionRepository) TransactionsUpdaterVerify(ctx context.Context, transactionId string) (transactions.Domain, error) {
	data := records.Transactions{}
	now := time.Now().Format(time.RFC3339)

	err := repo.ConnPostgres.Model(&data).Where("id", transactionId).Update("datetime_verify", now).Update("status", 1)
	if err.Error != nil {
		return transactions.Domain{}, err.Error
	}
	return data.TransactionsToDomain(), nil
}
func (repo *TransactionRepository) TransactionsUpdaterCompleted(ctx context.Context, transactionId string, status int) (transactions.Domain, error) {
	data := records.Transactions{}
	now := time.Now().Local()
	err := repo.ConnPostgres.Model(&data).Where("id", transactionId).Update("datetime_completed", now).Update("status", status)
	if err.Error != nil {
		return transactions.Domain{}, err.Error
	}
	return data.TransactionsToDomain(), nil
}
