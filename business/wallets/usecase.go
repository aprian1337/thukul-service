package wallets

import (
	"aprian1337/thukul-service/business/wallet_histories"
	"context"
	"time"
)

type WalletUsecase struct {
	Repo                 Repository
	WalletHistoryUsecase wallet_histories.Usecase
	Timeout              time.Duration
}

func (uc *WalletUsecase) Create(ctx context.Context, domain Domain) error {
	err := uc.Repo.Create(ctx, domain)
	if err != nil {
		return err
	}
	return nil
}

func NewWalletsUsecase(repo Repository, walletHistoryUsecase wallet_histories.Usecase, timeout time.Duration) *WalletUsecase {
	return &WalletUsecase{
		Repo:                 repo,
		WalletHistoryUsecase: walletHistoryUsecase,
		Timeout:              timeout,
	}
}

func (uc *WalletUsecase) GetByUserId(ctx context.Context, userId int) (Domain, error) {
	wallet, err := uc.Repo.GetByUserId(ctx, userId)
	if err != nil {
		return Domain{}, err
	}
	return wallet, nil
}
func (uc *WalletUsecase) UpdateByUserId(ctx context.Context, domain Domain) (Domain, error) {
	data, err := uc.Repo.UpdateByUserId(ctx, domain)
	if err != nil {
		return Domain{}, err
	}
	data.Kind = domain.Kind
	data.NominalTransaction = domain.NominalTransaction
	data.CoinId = domain.CoinId
	data.TransactionId = domain.TransactionId
	if data.Kind == "topup" {
		err := uc.WalletHistoryUsecase.WalletHistoriesCreate(ctx, data.ToHistoryDomain())
		if err != nil {
			return Domain{}, err
		}
	}

	return data, nil
}
