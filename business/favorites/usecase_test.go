package favorites_test

import (
	"aprian1337/thukul-service/business"
	"aprian1337/thukul-service/business/coins"
	_coinUsecase "aprian1337/thukul-service/business/coins/mocks"
	"aprian1337/thukul-service/business/favorites"
	_favRepository "aprian1337/thukul-service/business/favorites/mocks"
	"aprian1337/thukul-service/business/users"
	_userUsecase "aprian1337/thukul-service/business/users/mocks"
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

var favoritesRepository _favRepository.Repository
var favService favorites.Usecase
var coinService _coinUsecase.Usecase
var userService _userUsecase.Usecase

var coinDomain coins.Domain
var userDomain users.Domain
var favDomain favorites.Domain
var listFavDomain []favorites.Domain

var rowsAffectedSuccess int64
var rowsAffectedError int64

func setup() {
	favService = favorites.NewFavoriteUsecase(&favoritesRepository, &userService, &coinService, time.Second*10)
	favDomain = favorites.Domain{
		ID:     1,
		UserId: 1,
		CoinId: 1,
		Symbol: "BTC",
	}
	coinDomain = coins.Domain{
		Id:     1,
		Symbol: "BTC",
		Name:   "Bitcoin",
	}
	userDomain = users.Domain{
		ID:       1,
		SalaryId: 1,
		Name:     "ASD",
		Password: "123",
	}
	rowsAffectedSuccess = 1
	rowsAffectedError = 0
	listFavDomain = append(listFavDomain, favDomain)
}

func TestGetList(t *testing.T) {
	t.Run("Test Case 1 | GetList - Success", func(t *testing.T) {
		setup()
		favoritesRepository.On("FavoritesGetList", mock.Anything, mock.Anything).Return(listFavDomain, nil).Once()
		crypto, err := favService.GetList(context.Background(), favDomain.UserId)

		assert.Nil(t, err)
		assert.Equal(t, crypto, listFavDomain)
	})

	t.Run("Test Case 2 | GetList - Error", func(t *testing.T) {
		setup()
		favoritesRepository.On("FavoritesGetList", mock.Anything, mock.Anything).Return([]favorites.Domain{}, businesses.ErrForTest).Once()
		crypto, err := favService.GetList(context.Background(), favDomain.UserId)

		assert.Error(t, err)
		assert.Equal(t, crypto, []favorites.Domain{})
	})
}

func TestGetById(t *testing.T) {
	t.Run("Test Case 1 | GetById - Success", func(t *testing.T) {
		setup()
		favoritesRepository.On("FavoritesGetById", mock.Anything, mock.Anything, mock.Anything).Return(favDomain, nil).Once()
		crypto, err := favService.GetById(context.Background(), favDomain.UserId, favDomain.CoinId)

		assert.Nil(t, err)
		assert.Equal(t, crypto, favDomain)
	})

	t.Run("Test Case 2 | CryptosGetDetail - Error", func(t *testing.T) {
		setup()
		favoritesRepository.On("FavoritesGetById", mock.Anything, mock.Anything, mock.Anything).Return(favorites.Domain{}, businesses.ErrForTest).Once()
		crypto, err := favService.GetById(context.Background(), favDomain.UserId, favDomain.ID)

		assert.Error(t, err)
		assert.Equal(t, crypto, favorites.Domain{})
	})
}

func TestDelete(t *testing.T) {
	t.Run("Test Case 1 | Delete - Success", func(t *testing.T) {
		setup()
		favoritesRepository.On("FavoritesDelete",
			mock.Anything, mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(rowsAffectedSuccess, nil).Once()
		err := favService.Delete(context.Background(), favDomain.UserId, favDomain.ID)

		assert.NoError(t, err)

		favoritesRepository.AssertExpectations(t)
	})

	t.Run("Test Case 2 | Delete - Error", func(t *testing.T) {
		setup()
		favoritesRepository.On("FavoritesDelete",
			mock.Anything, mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(rowsAffectedError, businesses.ErrForTest).Once()
		err := favService.Delete(context.Background(), favDomain.UserId, favDomain.ID)

		assert.Error(t, err)

		favoritesRepository.AssertExpectations(t)
	})

	t.Run("Test Case 2 | Delete - Nothing to Destroy", func(t *testing.T) {
		setup()
		favoritesRepository.On("FavoritesDelete",
			mock.Anything, mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(rowsAffectedError, nil).Once()
		err := favService.Delete(context.Background(), favDomain.UserId, favDomain.ID)

		assert.Error(t, err)

		favoritesRepository.AssertExpectations(t)
	})
}

func TestCreate(t *testing.T) {
	t.Run("Test Case 1 | Create - Success", func(t *testing.T) {
		setup()
		coinService.On("GetBySymbol",
			mock.Anything, mock.Anything).Return(coinDomain, nil).Once()
		userService.On("GetById",
			mock.Anything, mock.Anything).Return(userDomain, nil).Once()
		favoritesRepository.On("FavoritesCheck",
			mock.Anything, mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(rowsAffectedError, nil).Once()
		favoritesRepository.On("FavoritesCreate",
			mock.Anything, mock.Anything).Return(favDomain, nil).Once()
		data, err := favService.Create(context.Background(), favDomain, favDomain.UserId)

		assert.NoError(t, err)
		assert.Equal(t, data.CoinId, favDomain.CoinId)
		favoritesRepository.AssertExpectations(t)
	})

}
