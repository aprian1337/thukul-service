package cryptos

import (
	"context"
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

func (uc *CryptoUsecase) UpdateBuyCoin(ctx context.Context, domain Domain) (Domain, error) {
	check, _ := uc.Repo.CryptosGetDetail(ctx, domain.UserId, domain.CoinId)
	data := Domain{}
	var err error
	if check.ID == 0 {
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
	domain.Qty = check.Qty + domain.BuyQty
	data, err = uc.Repo.CryptosUpdate(ctx, domain)
	if err != nil {
		return Domain{}, err
	}
	return data, nil
}
