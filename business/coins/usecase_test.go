package coins_test

import (
	businesses "aprian1337/thukul-service/business"
	"aprian1337/thukul-service/business/coinmarket"
	_mockCoinmarket "aprian1337/thukul-service/business/coinmarket/mocks"
	"aprian1337/thukul-service/business/coins"
	_mockCoins "aprian1337/thukul-service/business/coins/mocks"
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

var coinsRepository _mockCoins.Repository
var coinmarketRepository _mockCoinmarket.Repository

var rowsAffectedSuccess int64
var rowsAffectedError int64
var coinService coins.Usecase
var coinDomain coins.Domain
var coinmarketDomain coinmarket.Domain
var listCoinDomain []coins.Domain

func setup() {
	coinService = coins.NewCoinUsecase(&coinsRepository, &coinmarketRepository, time.Second*10)
	coinDomain = coins.Domain{
		Id:     1,
		Symbol: "BTC",
		Name:   "Bitcoin",
	}
	coinmarketDomain = coinmarket.Domain{
		Symbol: "BTC",
		Name:   "Bitcoin",
	}
	rowsAffectedSuccess = 1
	rowsAffectedError = 0
	listCoinDomain = append(listCoinDomain, coinDomain)
}

func TestGetBySymbol(t *testing.T) {
	t.Run("Test Case 1 | Get By Symbol - Success", func(t *testing.T) {
		setup()
		coinmarketRepository.On("GetBySymbol", mock.Anything, mock.AnythingOfType("string")).Return(coinmarketDomain, nil).Once()
		coinsRepository.On("CoinsGetSymbol", mock.Anything, mock.AnythingOfType("string")).Return(coinDomain, rowsAffectedSuccess, nil).Once()
		symbol, err := coinService.GetBySymbol(context.Background(), coinDomain.Symbol)

		assert.Equal(t, symbol.Symbol, coinmarketDomain.Symbol)
		assert.Nil(t, err)

		coinsRepository.AssertExpectations(t)
	})

	t.Run("Test Case 2 | Get By Symbol - Error", func(t *testing.T) {
		coinsRepository.On("CoinsGetSymbol", mock.Anything, mock.Anything).Return(coins.Domain{}, rowsAffectedError, businesses.ErrForTest).Once()
		coin, err := coinService.GetBySymbol(context.Background(), "Z")

		assert.Equal(t, coin, coins.Domain{})
		assert.Error(t, err)

		coinsRepository.AssertExpectations(t)
	})

	t.Run("Test Case 3 | Get By Symbol - Error - Found Coinmarket API", func(t *testing.T) {
		coinsRepository.On("CoinsGetSymbol", mock.Anything, mock.Anything).Return(coinDomain, rowsAffectedError, nil).Once()
		coinsRepository.On("CoinsCreateSymbol", mock.Anything, mock.Anything).Return(coinDomain, nil).Once()
		coinmarketRepository.On("GetBySymbol", mock.Anything, mock.AnythingOfType("string")).Return(coinmarketDomain, nil).Once()
		_, err := coinService.GetBySymbol(context.Background(), "ZTC")

		assert.Nil(t, err)

		coinsRepository.AssertExpectations(t)
	})

}

func TestGetAllSymbol(t *testing.T){
	t.Run("Test Case 1 | Success", func(t *testing.T) {
		coinsRepository.On("GetAllSymbol", mock.Anything).Return(listCoinDomain, nil).Once()
		data, err := coinService.GetAllSymbol(context.Background())
		assert.Equal(t, data, listCoinDomain)
		assert.NoError(t, err)
	})

	t.Run("Test Case 1 | Error", func(t *testing.T) {
		coinsRepository.On("GetAllSymbol", mock.Anything).Return([]coins.Domain{}, businesses.ErrForTest).Once()
		data, err := coinService.GetAllSymbol(context.Background())
		assert.Equal(t, data, []coins.Domain{})
		assert.Error(t, err)
	})
}
