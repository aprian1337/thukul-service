package wallets

import (
	"context"
	"time"
)

type WalletUsecase struct {
	Repo    Repository
	Timeout time.Duration
}

func (uc *WalletUsecase) Create(ctx context.Context, domain Domain) error {
	err := uc.Repo.Create(ctx, domain)
	if err != nil {
		return err
	}
	return nil
}

func NewWalletsUsecase(repo Repository, timeout time.Duration) *WalletUsecase {
	return &WalletUsecase{
		Repo:    repo,
		Timeout: timeout,
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

	return data, nil
}
