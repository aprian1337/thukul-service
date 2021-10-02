package payments

import (
	businesses "aprian1337/thukul-service/business"
	"aprian1337/thukul-service/business/coinmarket"
	"aprian1337/thukul-service/business/coins"
	"aprian1337/thukul-service/business/cryptos"
	"aprian1337/thukul-service/business/smtp"
	"aprian1337/thukul-service/business/transactions"
	"aprian1337/thukul-service/business/users"
	"aprian1337/thukul-service/business/wallet_histories"
	"aprian1337/thukul-service/business/wallets"
	"context"
	"time"
)

type PaymentUsecase struct {
	UsersUsecase         users.Usecase
	CryptoUsecase        cryptos.Usecase
	CoinUsecase          coins.Usecase
	WalletUsecase        wallets.Usecase
	TransactionUsecase   transactions.Usecase
	WalletHistoryUsecase wallet_histories.Usecase
	SmtpEmailUsecase     smtp.Usecase
	CoinMarketRepo       coinmarket.Repository
	Timeout              time.Duration
}

func NewPaymentUsecase(usersUsecase users.Usecase, smtpUsecase smtp.Usecase, cryptoUsecase cryptos.Usecase, coinUsecase coins.Usecase, coinMarketRepo coinmarket.Repository, walletsUsecase wallets.Usecase, walletsHistoryUsecase wallet_histories.Usecase, transactionsUsecase transactions.Usecase, timeoutContext time.Duration) *PaymentUsecase {
	return &PaymentUsecase{
		UsersUsecase:         usersUsecase,
		SmtpEmailUsecase:     smtpUsecase,
		CryptoUsecase:        cryptoUsecase,
		CoinUsecase:          coinUsecase,
		CoinMarketRepo:       coinMarketRepo,
		WalletUsecase:        walletsUsecase,
		TransactionUsecase:   transactionsUsecase,
		WalletHistoryUsecase: walletsHistoryUsecase,
		Timeout:              timeoutContext,
	}
}

func (uc *PaymentUsecase) SellCoin(ctx context.Context, domain Domain) error {
	//if domain.Coin == "" {
	//	return businesses.ErrCoinRequired
	//}
	//if domain.Qty == 0 {
	//	return businesses.ErrQtyRequired
	//}
	//wallet, err := uc.WalletUsecase.GetByUserId(ctx, domain.UserId)
	//if err != nil {
	//	return err
	//}
	//price, err := uc.CoinMarketRepo.GetPrice(ctx, domain.Coin, domain.Qty)
	//if err != nil {
	//	return businesses.ErrQtyRequired
	//}
	//diff := wallet.Total - price
	//if diff < 0 {
	//	return businesses.ErrWalletNotEnough
	//}
	//coin, err := uc.CoinUsecase.GetBySymbol(ctx, domain.Coin)
	//if err != nil {
	//	return businesses.ErrCoinNotFound
	//}
	//_, err = uc.TransactionUsecase.ActivitiesCreate(ctx, domain.ToTransactionDomain(coin.Id))
	//if err != nil {
	//	return businesses.ErrCoinNotFound
	//}
	//return nil
	panic("PANIC")
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
	user, err := uc.UsersUsecase.GetById(ctx, domain.UserId)
	if err != nil {
		return err
	}
	price, err := uc.CoinMarketRepo.GetPrice(ctx, domain.Coin, domain.Qty)
	if err != nil {
		return businesses.ErrQtyRequired
	}
	diff := user.Wallets.Total - price
	if diff < 0 {
		return businesses.ErrWalletNotEnough
	}
	coin, err := uc.CoinUsecase.GetBySymbol(ctx, domain.Coin)
	if err != nil {
		return businesses.ErrCoinNotFound
	}
	_, err = uc.TransactionUsecase.TransactionsCreate(ctx, domain.ToTransactionDomain(coin.Id))
	if err != nil {
		return businesses.ErrCoinNotFound
	}
	//helpers.SendMail()
	err = uc.SmtpEmailUsecase.SendMailSMTP(ctx, user.ToSmtpDomain("Email Confirmation", "Hello World!"))
	if err != nil {
		return err
	}
	return nil
}
