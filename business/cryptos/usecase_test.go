package cryptos_test

import (
	businesses "aprian1337/thukul-service/business"
	"aprian1337/thukul-service/business/cryptos"
	_cryptoMock "aprian1337/thukul-service/business/cryptos/mocks"
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

var cryptoRepository _cryptoMock.Repository
var cryptoService cryptos.Usecase

var cryptoDomain cryptos.Domain
var listCryptoDomain []cryptos.Domain

func setup() {
	cryptoService = cryptos.NewCryptoUsecase(&cryptoRepository, time.Second*10)
	cryptoDomain = cryptos.Domain{
		ID:      1,
		UserId:  1,
		CoinId:  1,
		Symbol:  "BTC",
		Qty:     1,
		BuyQty:  1,
		SellQty: 1,
	}
	listCryptoDomain = append(listCryptoDomain, cryptoDomain)
}

func TestCryptosGetByUser(t *testing.T) {
	t.Run("Test Case 1 | CryptosGetByUser - Success", func(t *testing.T) {
		setup()
		cryptoRepository.On("CryptosGetByUser", mock.Anything, mock.Anything).Return(listCryptoDomain, nil).Once()
		crypto, err := cryptoService.CryptosGetByUser(context.Background(), cryptoDomain.UserId)

		assert.Nil(t, err)
		assert.Equal(t, crypto, listCryptoDomain)
	})

	t.Run("Test Case 2 | CryptosGetByUser - Error", func(t *testing.T) {
		setup()
		cryptoRepository.On("CryptosGetByUser", mock.Anything, mock.Anything).Return([]cryptos.Domain{}, businesses.ErrForTest).Once()
		crypto, err := cryptoService.CryptosGetByUser(context.Background(), cryptoDomain.UserId)

		assert.Error(t, err)
		assert.Equal(t, crypto, []cryptos.Domain{})
	})
}

func TestCryptosGetDetail(t *testing.T) {
	t.Run("Test Case 1 | CryptosGetDetail - Success", func(t *testing.T) {
		setup()
		cryptoRepository.On("CryptosGetDetail", mock.Anything, mock.Anything, mock.Anything).Return(cryptoDomain, nil).Once()
		crypto, err := cryptoService.CryptosGetDetail(context.Background(), cryptoDomain.UserId, cryptoDomain.CoinId)

		assert.Nil(t, err)
		assert.Equal(t, crypto, cryptoDomain)
	})

	t.Run("Test Case 2 | CryptosGetDetail - Error", func(t *testing.T) {
		setup()
		cryptoRepository.On("CryptosGetDetail", mock.Anything, mock.Anything, mock.Anything).Return(cryptos.Domain{}, businesses.ErrForTest).Once()
		crypto, err := cryptoService.CryptosGetDetail(context.Background(), cryptoDomain.UserId, cryptoDomain.CoinId)

		assert.Error(t, err)
		assert.Equal(t, crypto, cryptos.Domain{})
	})
}

func TestUpdateBuyCoin(t *testing.T) {
	t.Run("Test Case 1 | CryptosGetDetail - Success - Have no coin before", func(t *testing.T) {
		setup()
		cryptoRepository.On("CryptosGetDetail", mock.Anything, mock.Anything, mock.Anything).Return(cryptos.Domain{}, nil).Once()
		cryptoRepository.On("CryptosCreate", mock.Anything, mock.Anything).Return(cryptoDomain, nil).Once()
		crypto, err := cryptoService.UpdateBuyCoin(context.Background(), cryptoDomain)
		assert.Nil(t, err)
		assert.Equal(t, crypto, cryptoDomain)
	})

	t.Run("Test Case 2 | CryptosGetDetail - Success - Have coin before", func(t *testing.T) {
		setup()
		cryptoRepository.On("CryptosGetDetail", mock.Anything, mock.Anything, mock.Anything).Return(cryptoDomain, nil).Once()
		cryptoRepository.On("CryptosUpdate", mock.Anything, mock.Anything).Return(cryptoDomain, nil).Once()
		crypto, err := cryptoService.UpdateBuyCoin(context.Background(), cryptoDomain)
		assert.Nil(t, err)
		assert.Equal(t, crypto, cryptoDomain)
	})

	t.Run("Test Case 3 | CryptosGetDetail - Success - Have coin before but error", func(t *testing.T) {
		setup()
		cryptoRepository.On("CryptosGetDetail", mock.Anything, mock.Anything, mock.Anything).Return(cryptoDomain, nil).Once()
		cryptoRepository.On("CryptosUpdate", mock.Anything, mock.Anything).Return(cryptos.Domain{}, businesses.ErrForTest).Once()
		crypto, err := cryptoService.UpdateBuyCoin(context.Background(), cryptoDomain)
		assert.Error(t, err)
		assert.Equal(t, crypto, cryptos.Domain{})
	})

	t.Run("Test Case 4 | CryptosGetDetail - Success - Have coin before but error", func(t *testing.T) {
		setup()
		cryptoRepository.On("CryptosGetDetail", mock.Anything, mock.Anything, mock.Anything).Return(cryptos.Domain{}, nil).Once()
		cryptoRepository.On("CryptosCreate", mock.Anything, mock.Anything).Return(cryptos.Domain{}, businesses.ErrForTest).Once()
		crypto, err := cryptoService.UpdateBuyCoin(context.Background(), cryptoDomain)
		assert.Error(t, err)
		assert.Equal(t, crypto, cryptos.Domain{})
	})
}

func TestUpdateSellCoin(t *testing.T) {
	t.Run("Test Case 1 | UpdateSellCoin - Success", func(t *testing.T) {
		setup()
		cryptoRepository.On("CryptosGetDetail", mock.Anything, mock.Anything, mock.Anything).Return(cryptoDomain, nil).Once()
		cryptoRepository.On("CryptosUpdate", mock.Anything, mock.Anything).Return(cryptoDomain, nil).Once()
		crypto, err := cryptoService.UpdateSellCoin(context.Background(), cryptoDomain)
		assert.Nil(t, err)
		assert.Equal(t, crypto, cryptoDomain)
	})

	t.Run("Test Case 2 | UpdateSellCoin - Error", func(t *testing.T) {
		setup()
		cryptoRepository.On("CryptosGetDetail", mock.Anything, mock.Anything, mock.Anything).Return(cryptoDomain, nil).Once()
		cryptoRepository.On("CryptosUpdate", mock.Anything, mock.Anything).Return(cryptos.Domain{}, businesses.ErrForTest).Once()
		crypto, err := cryptoService.UpdateSellCoin(context.Background(), cryptoDomain)
		assert.Error(t, err)
		assert.Equal(t, crypto, cryptos.Domain{})
	})
}
