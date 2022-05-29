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
	"aprian1337/thukul-service/helpers"
	"aprian1337/thukul-service/helpers/constants"
	"context"
	"fmt"
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
	KeyString            string
	KeyAdditional        string
	BaseUrl              string
	BasePort             string
	Timeout              time.Duration
}

func NewPaymentUsecase(usersUsecase users.Usecase, smtpUsecase smtp.Usecase, cryptoUsecase cryptos.Usecase, coinUsecase coins.Usecase, coinMarketRepo coinmarket.Repository, walletsUsecase wallets.Usecase, walletsHistoryUsecase wallet_histories.Usecase, transactionsUsecase transactions.Usecase, keyString string, keyAdditional string, baseUrl string, basePort string, timeoutContext time.Duration) *PaymentUsecase {
	return &PaymentUsecase{
		UsersUsecase:         usersUsecase,
		SmtpEmailUsecase:     smtpUsecase,
		CryptoUsecase:        cryptoUsecase,
		CoinUsecase:          coinUsecase,
		CoinMarketRepo:       coinMarketRepo,
		KeyString:            keyString,
		WalletUsecase:        walletsUsecase,
		BasePort:             basePort,
		BaseUrl:              baseUrl,
		TransactionUsecase:   transactionsUsecase,
		KeyAdditional:        keyAdditional,
		WalletHistoryUsecase: walletsHistoryUsecase,
		Timeout:              timeoutContext,
	}
}

func (uc *PaymentUsecase) SellCoin(ctx context.Context, domain Domain) error {
	if domain.Coin == "" {
		return businesses.ErrCoinRequired
	}

	if helpers.IsZero(domain.Qty) {
		return businesses.ErrQtyRequired
	}

	coin, err := uc.CoinUsecase.GetBySymbol(ctx, domain.Coin)
	if err != nil {
		return businesses.ErrCoinNotFound
	}

	crypto, err := uc.CryptoUsecase.CryptosGetDetail(ctx, domain.UserId, coin.Id)
	if err != nil {
		if err.Error() == "record not found" {
			return fmt.Errorf("you don't have this coin")
		}
		return err
	}

	if crypto.Qty < domain.Qty {
		return businesses.ErrCoinNotEnough
	}

	price, err := uc.CoinMarketRepo.GetPrice(ctx, domain.Coin, domain.Qty)
	if err != nil {
		return err
	}

	domain.Price = price
	transaction, err := uc.TransactionUsecase.TransactionsCreate(ctx, domain.ToTransactionDomain(coin.Id, "sell"))
	if err != nil {
		return err
	}

	user, err := uc.UsersUsecase.GetById(ctx, domain.UserId)
	if err != nil {
		return err
	}

	tomorrow := transaction.DatetimeRequest.Add(time.Hour * 24).Local()
	url := uc.BaseUrl + ":" + uc.BasePort + "/api/v1/payments/confirm/" + helpers.HashTransactionToSlug(transaction.Id, uc.KeyString, uc.KeyAdditional)
	bodyEmail := `
		<h2>Hello ` + user.Name + `!</h2><br/>
		You will selling a <b>` + ` ` + helpers.FloatToString(domain.Qty) + ` ` + coin.Symbol + ` (` + coin.Name + `)` + ` for Rp` + helpers.FloatToString(price) + `</b>, please confirm by clicking the link below to sell<br/><br/>
	` + url + `<br/><br/>
	This link will expired at ` + tomorrow.Format(constants.TimeFormat)
	err = uc.SmtpEmailUsecase.SendMailSMTP(ctx, user.ToSmtpDomain("Thukul Confirmation - SELL", bodyEmail))
	if err != nil {
		return err
	}
	_, err = uc.TransactionUsecase.TransactionsUpdaterVerify(ctx, transaction.Id)
	return nil
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

	if helpers.IsZero(domain.Qty) {
		return businesses.ErrQtyRequired
	}
	user, err := uc.UsersUsecase.GetById(ctx, domain.UserId)
	if err != nil {
		return err
	}

	coin, err := uc.CoinUsecase.GetBySymbol(ctx, domain.Coin)
	if err != nil {
		return businesses.ErrCoinNotFound
	}

	price, err := uc.CoinMarketRepo.GetPrice(ctx, domain.Coin, domain.Qty)
	if err != nil {
		return err
	}
	diff := user.Wallets.Total - price
	if diff < 0 {
		return businesses.ErrWalletNotEnough
	}
	domain.Price = price
	transaction, err := uc.TransactionUsecase.TransactionsCreate(ctx, domain.ToTransactionDomain(coin.Id, "buy"))
	if err != nil {
		return err
	}
	tomorrow := transaction.DatetimeRequest.Add(time.Hour * 24).Local()
	url := uc.BaseUrl + ":" + uc.BasePort + "/api/v1/payments/confirm/" + helpers.HashTransactionToSlug(transaction.Id, uc.KeyString, uc.KeyAdditional)
	bodyEmail := `
		<h2>Hello ` + user.Name + `!</h2><br/>
		You will buying a <b>` + ` ` + helpers.FloatToString(domain.Qty) + ` ` + coin.Symbol + ` (` + coin.Name + `)` + ` for Rp` + helpers.FloatToString(price) + `</b>, please confirm by clicking the link below to purchase<br/><br/>
	` + url + `<br/><br/>
	This link will expired at ` + tomorrow.Format(constants.TimeFormat)
	err = uc.SmtpEmailUsecase.SendMailSMTP(ctx, user.ToSmtpDomain("Thukul Confirmation - BUY", bodyEmail))
	if err != nil {
		return err
	}
	_, err = uc.TransactionUsecase.TransactionsUpdaterVerify(ctx, transaction.Id)
	return nil
}

func (uc *PaymentUsecase) Confirm(ctx context.Context, encode string, encrypt string) (wallets.Domain, error) {
	match, _ := helpers.DecodeTransactionFromSlug(encode, encrypt, uc.KeyString, uc.KeyAdditional)
	if match == "" {
		return wallets.Domain{}, businesses.ErrInvalidPayload
	}

	transaction, err := uc.TransactionUsecase.TransactionsById(ctx, match)
	if err != nil {
		return wallets.Domain{}, err
	}

	if transaction.DatetimeCompleted != nil {
		return wallets.Domain{}, businesses.ErrHasBeenVerified
	}
	expired := transaction.DatetimeRequest.Add(time.Hour * 24).Local()
	now := time.Now()
	if now.After(expired) {
		_, err := uc.TransactionUsecase.TransactionsUpdaterCompleted(ctx, transaction.Id, -1)
		if err != nil {
			return wallets.Domain{}, err
		}
		return wallets.Domain{}, businesses.ErrExpiredConfirm
	}
	_, err = uc.TransactionUsecase.TransactionsUpdaterCompleted(ctx, transaction.Id, 2)
	if err != nil {
		return wallets.Domain{}, err
	}
	wallet, err := uc.WalletUsecase.GetByUserId(ctx, transaction.UserId)
	if err != nil {
		return wallets.Domain{}, nil
	}

	crypto, err := uc.CryptoUsecase.CryptosGetDetail(ctx, transaction.UserId, transaction.CoinId)

	if transaction.Kind == "buy" {
		if wallet.Total-transaction.Price < 0 {
			return wallets.Domain{}, businesses.ErrWalletNotEnoughVerify
		}
		wallet.Total -= transaction.Price
		wallet.NominalTransaction = transaction.Price

		_, err = uc.CryptoUsecase.UpdateBuyCoin(ctx, cryptos.Domain{
			UserId: transaction.UserId,
			CoinId: transaction.CoinId,
			BuyQty: transaction.Qty,
		})
		wallet.Kind = "buy"

	} else {
		if transaction.Qty > crypto.Qty {
			return wallets.Domain{}, businesses.ErrCoinNotEnough
		}
		wallet.Total += transaction.Price
		wallet.NominalTransaction = transaction.Price

		_, err = uc.CryptoUsecase.UpdateSellCoin(ctx, cryptos.Domain{
			UserId:  transaction.UserId,
			CoinId:  transaction.CoinId,
			SellQty: transaction.Qty,
		})
		wallet.Kind = "sell"

	}
	if err != nil {
		return wallets.Domain{}, err
	}
	wallet, err = uc.WalletUsecase.UpdateByUserId(ctx, wallet)
	if err != nil {
		return wallets.Domain{}, err
	}
	err = uc.WalletHistoryUsecase.WalletHistoriesCreate(ctx, ToWalletHistoriesDomain(wallet.Id, transaction.Id, transaction.Price))
	return wallet, nil
}
