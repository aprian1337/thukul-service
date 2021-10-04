package payments_test

import (
	businesses "aprian1337/thukul-service/business"
	"aprian1337/thukul-service/business/coinmarket"
	_coinmarketMock "aprian1337/thukul-service/business/coinmarket/mocks"
	"aprian1337/thukul-service/business/coins"
	_coinsMock "aprian1337/thukul-service/business/coins/mocks"
	"aprian1337/thukul-service/business/cryptos"
	_cryptosMock "aprian1337/thukul-service/business/cryptos/mocks"
	"aprian1337/thukul-service/business/payments"
	"aprian1337/thukul-service/business/smtp"
	_smtpMock "aprian1337/thukul-service/business/smtp/mocks"
	"aprian1337/thukul-service/business/transactions"
	_transactionsMock "aprian1337/thukul-service/business/transactions/mocks"
	"aprian1337/thukul-service/business/users"
	_usersMock "aprian1337/thukul-service/business/users/mocks"
	"aprian1337/thukul-service/business/wallet_histories"
	_walletHistoriesMock "aprian1337/thukul-service/business/wallet_histories/mocks"
	"aprian1337/thukul-service/business/wallets"
	_walletsMock "aprian1337/thukul-service/business/wallets/mocks"
	"context"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

var paymentService payments.Usecase
var userService _usersMock.Usecase
var cryptoService _cryptosMock.Usecase
var coinService _coinsMock.Usecase
var walletService _walletsMock.Usecase
var transactionService _transactionsMock.Usecase
var walletHistoriesService _walletHistoriesMock.Usecase
var smtpService _smtpMock.Usecase
var coinmarketRepo _coinmarketMock.Repository
var KeyString string
var KeyAdditional string
var BaseUrl string
var BasePort string

var userDomain users.Domain
var smtpDomain smtp.Domain
var cryptoDomain cryptos.Domain
var coinDomain coins.Domain
var coinmarketDomain coinmarket.Domain
var walletDomain wallets.Domain
var transactionDomain transactions.Domain
var transactionDomainSell transactions.Domain
var transactionDomainExp transactions.Domain
var walletHistoriesDomain wallet_histories.Domain
var paymentDomainTopup payments.Domain
var paymentDomainBuy payments.Domain
var paymentDomainSell payments.Domain
var price float64
var saldoSuccess float64
var saldoErr float64
var encode string
var encrypt string

func setup() {
	encode = "MjU3NjU4NjgtMDVkNC00OGU3LWFlZDItZjQ0NTdhMTIyZWY3"
	encrypt = "423565abbaaecc46ba7143272ae940503cfc07c6553206708171dbb761015d6686b98a8d1b563d7941c11a79ede3417a29f9eda2920d72b2a231ea507deb0c088876ce67523347962948764fea21ba7a1e5a5237688d5743a9323facc543f4336d569038"
	KeyAdditional = "secr3tsekalitem4n"
	KeyString = "5ff37e7594e33ce1dcedd17644dc542d"
	paymentService = payments.NewPaymentUsecase(&userService, &smtpService, &cryptoService, &coinService, &coinmarketRepo, &walletService, &walletHistoriesService, &transactionService, KeyString, KeyAdditional, BaseUrl, BasePort, time.Second*10)
	transactionId, _ := uuid.Parse("25765868-05d4-48e7-aed2-f4457a122ef7")
	walletHistoriesDomain = wallet_histories.Domain{
		ID:            1,
		WalletId:      1,
		TransactionId: &transactionId,
		Nominal:       9999,
	}
	userDomain = users.Domain{
		ID:       1,
		SalaryId: 1,
		Wallets: users.Wallets{
			Total: saldoSuccess,
		},
		Name:     "ABC",
		Password: "abc",
		IsAdmin:  0,
		Email:    "dddd@saz.com",
		Phone:    "",
		Gender:   "",
	}
	smtpDomain = smtp.Domain{
		MailTo:  []string{"dskaodjai@sdoadk.cada"},
		Subject: "dasijd@dsao.sa",
		Message: "jiodjasoidjas",
	}
	cryptoDomain = cryptos.Domain{
		ID:      1,
		UserId:  1,
		CoinId:  1,
		Symbol:  "BTC",
		Qty:     1000,
		BuyQty:  1,
		SellQty: 1,
	}
	coinDomain = coins.Domain{
		Id:     1,
		Symbol: "VYAS",
		Name:   "OAJSOIA",
	}
	coinmarketDomain = coinmarket.Domain{
		Symbol: "BTC",
		Name:   "Bitcoin",
	}

	walletDomain = wallets.Domain{
		Id:                 1,
		UserId:             1,
		Total:              213221,
		NominalTransaction: 11231,
		Kind:               "SAS",
		TransactionId:      &transactionId,
		CoinId:             1,
	}
	transactionDomain = transactions.Domain{
		UserId:            1,
		CoinId:            1,
		Qty:               111231,
		Price:             2231,
		Kind:              "buy",
		DatetimeRequest:   time.Date(2050, 2, 20, 2, 2, 2, 2, time.Local),
		DatetimeVerify:    nil,
		DatetimeCompleted: nil,
	}
	transactionDomainSell = transactions.Domain{
		UserId:            1,
		CoinId:            1,
		Qty:               1,
		Price:             231,
		Kind:              "sell",
		DatetimeRequest:   time.Date(2050, 2, 20, 2, 2, 2, 2, time.Local),
		DatetimeVerify:    nil,
		DatetimeCompleted: nil,
	}
	transactionDomainExp = transactions.Domain{
		UserId:            1,
		CoinId:            1,
		Qty:               111231,
		Price:             2131231,
		Kind:              "buy",
		DatetimeRequest:   time.Date(2000, 2, 20, 2, 2, 2, 2, time.Local),
		DatetimeVerify:    nil,
		DatetimeCompleted: nil,
	}
	paymentDomainTopup = payments.Domain{
		UserId:  1,
		Kind:    "topup",
		Coin:    "",
		Nominal: 111110,
	}
	price = 9019291
	saldoSuccess = 9999999
	saldoErr = 999
}

func TestTopUp(t *testing.T) {
	t.Run("Test Case 1 | Topup Success", func(t *testing.T) {
		setup()
		walletService.On("GetByUserId", mock.Anything, mock.Anything).Return(walletDomain, nil).Once()
		walletService.On("UpdateByUserId", mock.Anything, mock.Anything).Return(walletDomain, nil).Once()
		_, err := paymentService.TopUp(context.Background(), payments.Domain{
			UserId:  1,
			Kind:    "topup",
			Coin:    "BTC",
			Qty:     10,
			Price:   110,
			Nominal: 10,
		})

		assert.NotNil(t, walletDomain)
		assert.Nil(t, err)
	})

	t.Run("Test Case 2 | Topup - Error Nominal/Userid Required", func(t *testing.T) {
		setup()
		walletDomain, err := paymentService.TopUp(context.Background(), payments.Domain{
			Kind:  "topup",
			Coin:  "BTC",
			Qty:   10,
			Price: 110,
		})

		assert.Equal(t, wallets.Domain{}, walletDomain)
		assert.Error(t, err)
	})
}

func TestBuyCoin(t *testing.T) {
	t.Run("Test Case 1 | Buy Success", func(t *testing.T) {
		setup()
		userService.On("GetById", mock.Anything, mock.Anything).Return(userDomain, nil).Once()
		coinmarketRepo.On("GetPrice", mock.Anything, mock.Anything, mock.Anything).Return(price, nil).Once()
		coinService.On("GetBySymbol", mock.Anything, mock.Anything).Return(coinDomain, nil).Once()
		transactionService.On("TransactionsCreate", mock.Anything, mock.Anything).Return(transactionDomain, nil).Once()
		transactionService.On("TransactionsUpdaterVerify", mock.Anything, mock.Anything).Return(transactionDomain, nil).Once()
		smtpService.On("SendMailSMTP", mock.Anything, mock.Anything).Return(nil).Once()
		err := paymentService.BuyCoin(context.Background(), payments.Domain{
			UserId:  1,
			Kind:    "buy",
			Coin:    "BTC",
			Qty:     10,
			Price:   110,
			Nominal: 10,
		})

		assert.Nil(t, err)
	})

	t.Run("Test Case 2 | Buy - Error - Coin Required", func(t *testing.T) {
		setup()
		err := paymentService.BuyCoin(context.Background(), payments.Domain{
			Qty: 123,
		})
		assert.Error(t, err)
	})

	t.Run("Test Case 3 | Buy - Error - User Not Found", func(t *testing.T) {
		setup()
		userService.On("GetById", mock.Anything, mock.Anything).Return(users.Domain{}, businesses.ErrForTest).Once()
		err := paymentService.BuyCoin(context.Background(), payments.Domain{
			Coin: "BTC",
			Qty:  1231,
		})
		assert.Error(t, err)
	})

	t.Run("Test Case 4 | Buy - Coin Symbol Not Found", func(t *testing.T) {
		setup()
		coinService.On("GetBySymbol", mock.Anything, mock.Anything).Return(coins.Domain{}, businesses.ErrForTest).Once()
		userService.On("GetById", mock.Anything, mock.Anything).Return(userDomain, nil).Once()
		err := paymentService.BuyCoin(context.Background(), payments.Domain{
			Coin: "BTC",
			Qty:  1231,
		})
		assert.Error(t, err)
	})

	t.Run("Test Case 5 | Buy - Error - Price Err", func(t *testing.T) {
		setup()
		coinmarketRepo.On("GetPrice", mock.Anything, mock.Anything, mock.Anything).Return(price, businesses.ErrForTest).Once()
		coinService.On("GetBySymbol", mock.Anything, mock.Anything).Return(coinDomain, nil).Once()
		userService.On("GetById", mock.Anything, mock.Anything).Return(userDomain, nil).Once()
		err := paymentService.BuyCoin(context.Background(), payments.Domain{
			Coin: "BTC",
			Qty:  1231,
		})
		assert.Error(t, err)
	})

	t.Run("Test Case 6 | Buy - Error - Send Email Error", func(t *testing.T) {
		setup()
		coinmarketRepo.On("GetPrice", mock.Anything, mock.Anything, mock.Anything).Return(price, nil).Once()
		coinService.On("GetBySymbol", mock.Anything, mock.Anything).Return(coinDomain, nil).Once()
		userService.On("GetById", mock.Anything, mock.Anything).Return(userDomain, nil).Once()
		transactionService.On("TransactionsCreate", mock.Anything, mock.Anything).Return(transactionDomain, nil).Once()
		smtpService.On("SendMailSMTP", mock.Anything, mock.Anything).Return(businesses.ErrForTest).Once()
		err := paymentService.BuyCoin(context.Background(), payments.Domain{
			Coin: "BTC",
			Qty:  1231,
		})
		assert.Error(t, err)
	})
}

func TestSellCoin(t *testing.T) {
	t.Run("Test Case 1 | Sell Success", func(t *testing.T) {
		setup()
		userService.On("GetById", mock.Anything, mock.Anything).Return(userDomain, nil).Once()
		coinmarketRepo.On("GetPrice", mock.Anything, mock.Anything, mock.Anything).Return(price, nil).Once()
		cryptoService.On("CryptosGetDetail", mock.Anything, mock.Anything, mock.Anything).Return(cryptoDomain, nil)
		coinService.On("GetBySymbol", mock.Anything, mock.Anything).Return(coinDomain, nil).Once()
		transactionService.On("TransactionsCreate", mock.Anything, mock.Anything).Return(transactionDomain, nil).Once()
		transactionService.On("TransactionsUpdaterVerify", mock.Anything, mock.Anything).Return(transactionDomain, nil).Once()
		smtpService.On("SendMailSMTP", mock.Anything, mock.Anything).Return(nil).Once()
		err := paymentService.SellCoin(context.Background(), payments.Domain{
			UserId:  1,
			Kind:    "buy",
			Coin:    "BTC",
			Qty:     10,
			Price:   110,
			Nominal: 10,
		})

		assert.Nil(t, err)
	})

	t.Run("Test Case 2 | Sell - Error - Coin Required", func(t *testing.T) {
		setup()
		err := paymentService.SellCoin(context.Background(), payments.Domain{
			Qty: 123,
		})
		assert.Error(t, err)
	})

	t.Run("Test Case 3 | Sell - Error - Qty Required", func(t *testing.T) {
		setup()
		err := paymentService.SellCoin(context.Background(), payments.Domain{
			Coin: "BTC",
		})
		assert.Error(t, err)
	})

	t.Run("Test Case 4 | Sell - Error - Symbol Not Found", func(t *testing.T) {
		setup()
		coinService.On("GetBySymbol", mock.Anything, mock.Anything).Return(coins.Domain{}, businesses.ErrForTest).Once()
		err := paymentService.SellCoin(context.Background(), payments.Domain{
			Coin: "BTC",
			Qty:  1231,
		})
		assert.Error(t, err)
	})

	t.Run("Test Case 5 | Sell - Error - Haven't Crypto", func(t *testing.T) {
		setup()
		cryptoService.On("CryptosGetDetail", mock.Anything, mock.Anything, mock.Anything).Return(cryptoDomain, nil).Once()
		coinService.On("GetBySymbol", mock.Anything, mock.Anything).Return(coins.Domain{}, businesses.ErrForTest).Once()
		err := paymentService.SellCoin(context.Background(), payments.Domain{
			Coin: "BTC",
			Qty:  1231,
		})
		assert.Error(t, err)
	})

	t.Run("Test Case 5 | Sell - Error - Qty Required", func(t *testing.T) {
		setup()
		cryptoService.On("CryptosGetDetail", mock.Anything, mock.Anything, mock.Anything).Return(cryptoDomain, nil).Once()
		coinService.On("GetBySymbol", mock.Anything, mock.Anything).Return(coinDomain, nil).Once()
		coinmarketRepo.On("GetPrice", mock.Anything, mock.Anything).Return(price, businesses.ErrForTest).Once()
		err := paymentService.SellCoin(context.Background(), payments.Domain{
			Coin: "BTC",
			Qty:  1231,
		})
		assert.Error(t, err)
	})

	t.Run("Test Case 5 | Sell - Error - Price Err", func(t *testing.T) {
		setup()
		cryptoService.On("CryptosGetDetail", mock.Anything, mock.Anything, mock.Anything).Return(cryptoDomain, nil).Once()
		coinService.On("GetBySymbol", mock.Anything, mock.Anything).Return(coinDomain, nil).Once()
		coinmarketRepo.On("GetPrice", mock.Anything, mock.Anything).Return(price, nil).Once()
		transactionService.On("TransactionsCreate", mock.Anything, mock.Anything).Return(transactions.Domain{}, businesses.ErrForTest).Once()
		err := paymentService.SellCoin(context.Background(), payments.Domain{
			Coin: "BTC",
			Qty:  1231,
		})
		assert.Error(t, err)
	})

	t.Run("Test Case 6 | Sell - Error - Send Email Error", func(t *testing.T) {
		setup()
		cryptoService.On("CryptosGetDetail", mock.Anything, mock.Anything, mock.Anything).Return(cryptoDomain, nil).Once()
		coinService.On("GetBySymbol", mock.Anything, mock.Anything).Return(coinDomain, nil).Once()
		coinmarketRepo.On("GetPrice", mock.Anything, mock.Anything).Return(price, nil).Once()
		transactionService.On("TransactionsCreate", mock.Anything, mock.Anything).Return(transactionDomain, nil).Once()
		userService.On("GetById", mock.Anything, mock.Anything).Return(users.Domain{}, businesses.ErrForTest).Once()
		err := paymentService.SellCoin(context.Background(), payments.Domain{
			Coin: "BTC",
			Qty:  1231,
		})
		assert.Error(t, err)
	})

	t.Run("Test Case 6 | Sell - Error - Send Email Error", func(t *testing.T) {
		setup()
		cryptoService.On("CryptosGetDetail", mock.Anything, mock.Anything, mock.Anything).Return(cryptoDomain, nil).Once()
		coinService.On("GetBySymbol", mock.Anything, mock.Anything).Return(coinDomain, nil).Once()
		coinmarketRepo.On("GetPrice", mock.Anything, mock.Anything).Return(price, nil).Once()
		transactionService.On("TransactionsCreate", mock.Anything, mock.Anything).Return(transactionDomain, nil).Once()
		userService.On("GetById", mock.Anything, mock.Anything).Return(userDomain, nil).Once()
		smtpService.On("SendMailSMTP", mock.Anything, mock.Anything).Return(businesses.ErrForTest).Once()
		err := paymentService.SellCoin(context.Background(), payments.Domain{
			Coin: "BTC",
			Qty:  1231,
		})
		assert.Error(t, err)
	})
}

func TestConfirm(t *testing.T) {
	t.Run("Test Case 1 | Confirm Buy Success", func(t *testing.T) {
		transactionService.On("TransactionsById", mock.Anything, mock.Anything).Return(transactionDomain, nil).Once()
		transactionService.On("TransactionsUpdaterCompleted", mock.Anything, mock.Anything, mock.Anything).Return(transactionDomain, nil).Once()
		walletService.On("GetByUserId", mock.Anything, mock.Anything).Return(walletDomain, nil).Once()
		cryptoService.On("CryptosGetDetail", mock.Anything, mock.Anything, mock.Anything).Return(walletDomain, nil).Once()
		cryptoService.On("UpdateBuyCoin", mock.Anything, mock.Anything, mock.Anything).Return(cryptoDomain, nil).Once()
		walletService.On("UpdateByUserId", mock.Anything, mock.Anything).Return(walletDomain, nil).Once()
		walletHistoriesService.On("WalletHistoriesCreate", mock.Anything, mock.Anything).Return(nil).Once()
		data, err := paymentService.Confirm(context.Background(), encode, encrypt)
		assert.Equal(t, walletDomain, data)
		assert.NoError(t, err)
	})

	t.Run("Test Case 2 | Confirm Sell Success", func(t *testing.T) {
		transactionService.On("TransactionsById", mock.Anything, mock.Anything).Return(transactionDomainSell, nil).Once()
		transactionService.On("TransactionsUpdaterCompleted", mock.Anything, mock.Anything, mock.Anything).Return(transactionDomainSell, nil).Once()
		walletService.On("GetByUserId", mock.Anything, mock.Anything).Return(walletDomain, nil).Once()
		cryptoService.On("CryptosGetDetail", mock.Anything, mock.Anything, mock.Anything).Return(walletDomain, nil).Once()
		cryptoService.On("UpdateSellCoin", mock.Anything, mock.Anything, mock.Anything).Return(cryptoDomain, nil).Once()
		walletService.On("UpdateByUserId", mock.Anything, mock.Anything).Return(walletDomain, nil).Once()
		walletHistoriesService.On("WalletHistoriesCreate", mock.Anything, mock.Anything).Return(nil).Once()
		data, err := paymentService.Confirm(context.Background(), encode, encrypt)
		assert.Equal(t, walletDomain, data)
		assert.NoError(t, err)
	})

	t.Run("Test Case 3 | Error - Not Match", func(t *testing.T) {
		data, err := paymentService.Confirm(context.Background(), "encode", encrypt)
		assert.Equal(t, wallets.Domain{}, data)
		assert.Error(t, err)
	})

	t.Run("Test Case 4 | Error - Transaction Not Found", func(t *testing.T) {
		transactionService.On("TransactionsById", mock.Anything, mock.Anything).Return(transactions.Domain{}, businesses.ErrForTest).Once()
		data, err := paymentService.Confirm(context.Background(), encode, encrypt)
		assert.Equal(t, wallets.Domain{}, data)
		assert.Error(t, err)
	})

	t.Run("Test Case 5 | Error - Transaction Expired - Update Transaction Success", func(t *testing.T) {
		transactionService.On("TransactionsById", mock.Anything, mock.Anything).Return(transactionDomainExp, nil).Once()
		transactionService.On("TransactionsUpdaterCompleted", mock.Anything, mock.Anything, mock.Anything).Return(transactionDomainExp, nil).Once()
		data, err := paymentService.Confirm(context.Background(), encode, encrypt)
		assert.Equal(t, wallets.Domain{}, data)
		assert.Error(t, err)
	})

	t.Run("Test Case 6 | Error - Transaction Expired - Update Transaction Fail", func(t *testing.T) {
		transactionService.On("TransactionsById", mock.Anything, mock.Anything).Return(transactionDomainExp, nil).Once()
		transactionService.On("TransactionsUpdaterCompleted", mock.Anything, mock.Anything, mock.Anything).Return(transactions.Domain{}, businesses.ErrForTest).Once()
		data, err := paymentService.Confirm(context.Background(), encode, encrypt)
		assert.Equal(t, wallets.Domain{}, data)
		assert.Error(t, err)
	})
}
