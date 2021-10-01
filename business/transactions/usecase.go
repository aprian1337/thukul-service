package transactions

import (
	businesses "aprian1337/thukul-service/business"
	"context"
	"time"
)

type TransactionUsecase struct {
	Repo    Repository
	Timeout time.Duration
}

func NewTransactionUsecase(repo Repository, timeout time.Duration) *TransactionUsecase {
	return &TransactionUsecase{
		Repo:    repo,
		Timeout: timeout,
	}
}

func (uc *TransactionUsecase) Create(ctx context.Context, domain Domain) (Domain, error) {
	d, err := uc.Repo.Create(ctx, domain)
	if err != nil {
		return Domain{}, err
	}
	return d, nil
}

func (uc *TransactionUsecase) UpdaterVerify(ctx context.Context, transactionId string) (Domain, error) {
	d, err := uc.Repo.UpdaterVerify(ctx, transactionId)
	if err != nil {
		return Domain{}, err
	}
	return d, nil
}

func (uc *TransactionUsecase) UpdaterCompleted(ctx context.Context, transactionId string, status int) (Domain, error) {
	if status != 2 && status != -1 {
		return Domain{}, businesses.ErrBadRequest
	}
	d, err := uc.Repo.UpdaterCompleted(ctx, transactionId, status)
	if err != nil {
		return Domain{}, err
	}
	return d, nil
}
