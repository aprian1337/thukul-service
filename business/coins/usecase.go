package coins

import (
	businesses "aprian1337/thukul-service/business"
	"aprian1337/thukul-service/business/coinmarket"
	"context"
	"strings"
	"time"
)

type CoinUsecase struct {
	Repo           Repository
	CoinMarketRepo coinmarket.Repository
	Timeout        time.Duration
}

func NewCoinUsecase(repo Repository, coinMarketRepo coinmarket.Repository, time time.Duration) *CoinUsecase {
	return &CoinUsecase{
		Repo:           repo,
		CoinMarketRepo: coinMarketRepo,
		Timeout:        time,
	}
}

func (uc *CoinUsecase) GetBySymbol(ctx context.Context, symbol string) (Domain, error) {
	symbol = strings.ToUpper(symbol)
	data, count, err := uc.Repo.CoinsGetSymbol(ctx, symbol)
	if err != nil {
		return Domain{}, err
	}
	if count == 0 {
		symbolData, errApi := uc.CoinMarketRepo.GetBySymbol(ctx, symbol)
		if errApi != nil {
			return Domain{}, businesses.ErrNotFound
		}
		domain := Domain{
			Symbol: symbolData.Symbol,
			Name:   symbolData.Name,
		}
		createSymbol, err := uc.Repo.CoinsCreateSymbol(ctx, domain)
		if err != nil {
			return Domain{}, err
		}
		return createSymbol, nil
	}

	return data, nil
}

func (uc *CoinUsecase) GetAllSymbol(ctx context.Context) ([]Domain, error) {
	data, err := uc.Repo.GetAllSymbol(ctx)
	if err != nil {
		return []Domain{}, err
	}
	return data, nil
}
