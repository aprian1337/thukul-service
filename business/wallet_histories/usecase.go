package wallet_histories

import (
	"context"
	"time"
)

type WalletHistoryUseacse struct {
	Repo    Repository
	Timeout time.Duration
}

func NewWalletsUsecase(repo Repository, timeout time.Duration) *WalletHistoryUseacse {
	return &WalletHistoryUseacse{
		Repo:    repo,
		Timeout: timeout,
	}
}

func (uc *WalletHistoryUseacse) WalletHistoriesCreate(ctx context.Context, domain Domain) error {
	err := uc.Repo.WalletHistoriesCreate(ctx, domain)
	if err != nil {
		return err
	}
	return nil
}
