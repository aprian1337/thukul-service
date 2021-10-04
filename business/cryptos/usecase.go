package cryptos

import (
	"context"
	"fmt"
	"time"
)

type CryptoUsecase struct {
	Repo    Repository
	Timeout time.Duration
}

func NewCryptoUsecase(repo Repository, timeout time.Duration) *CryptoUsecase {
	return &CryptoUsecase{
		Repo:    repo,
		Timeout: timeout,
	}
}

func (uc *CryptoUsecase) CryptosGetByUser(ctx context.Context, userId int) ([]Domain, error) {
	data, err := uc.Repo.CryptosGetByUser(ctx, userId)
	if err != nil {
		return []Domain{}, nil
	}
	return data, nil
}

func (uc *CryptoUsecase) CryptosGetDetail(ctx context.Context, userId int, coinId int) (Domain, error) {
	data, err := uc.Repo.CryptosGetDetail(ctx, userId, coinId)
	if err != nil {
		return Domain{}, err
	}
	return data, nil
}

func (uc *CryptoUsecase) UpdateBuyCoin(ctx context.Context, domain Domain) (Domain, error) {
	check, _ := uc.Repo.CryptosGetDetail(ctx, domain.UserId, domain.CoinId)
	data := Domain{}
	var err error
	if check.ID == 0 {
		domain.Qty = domain.BuyQty
		data, err = uc.Repo.CryptosCreate(ctx, domain)
		if err != nil {
			return Domain{}, err
		}
		return data, nil
	} else {
		domain.Qty = check.Qty + domain.BuyQty
		data, err = uc.Repo.CryptosUpdate(ctx, domain)
		if err != nil {
			return Domain{}, err
		}
		return data, nil
	}
}

func (uc *CryptoUsecase) UpdateSellCoin(ctx context.Context, domain Domain) (Domain, error) {
	check, _ := uc.Repo.CryptosGetDetail(ctx, domain.UserId, domain.CoinId)
	data := Domain{}
	var err error
	domain.Qty = check.Qty - domain.SellQty
	data, err = uc.Repo.CryptosUpdate(ctx, domain)
	if err != nil {
		return Domain{}, err
	}
	return data, nil
}
