package payments

import (
	businesses "aprian1337/thukul-service/business"
	"aprian1337/thukul-service/business/wallet_histories"
	"aprian1337/thukul-service/business/wallets"
	"context"
	"time"
)

type PaymentUsecase struct {
	WalletUsecase wallets.Usecase
	Timeout       time.Duration
}

func (uc *PaymentUsecase) SellCoin(ctx context.Context, domain Domain) error {
	panic("implement me")
}

func NewPaymentUsecase(walletUsecase wallets.Usecase, walletHistoryUsecase wallet_histories.Usecase, timeout time.Duration) *PaymentUsecase {
	return &PaymentUsecase{
		WalletUsecase: walletUsecase,
		Timeout:       timeout,
	}
}

func (uc *PaymentUsecase) TopUp(ctx context.Context, domain Domain) (wallets.Domain, error) {
	if domain.Nominal == 0 || domain.UserId == 0 {
		return wallets.Domain{}, businesses.ErrBadRequest
	}
	wallet, err := uc.WalletUsecase.GetByUserId(ctx, domain.UserId)
	if err != nil {
		return wallets.Domain{}, businesses.ErrUserIdNotFound
	}
	wallet.Total += domain.Nominal
	wallet.NominalTransaction = domain.Nominal
	wallet.Kind = "topup"
	_, err = uc.WalletUsecase.UpdateByUserId(ctx, wallet)
	if err != nil {
		return wallets.Domain{}, nil
	}
	return wallet, nil
}

func (uc *PaymentUsecase) BuyCoin(ctx context.Context, domain Domain) error {
	if domain.Coin == "" {
		return businesses.ErrCoinRequired
	}
	if domain.Qty == 0 {
		return businesses.ErrQtyRequired
	}
	return businesses.ErrUserIdRequired
}
