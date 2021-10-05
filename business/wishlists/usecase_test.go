package wishlists_test

import (
	"aprian1337/thukul-service/business"
	"aprian1337/thukul-service/business/users"
	_usersMock "aprian1337/thukul-service/business/users/mocks"
	"aprian1337/thukul-service/business/wishlists"
	_wishlistMock "aprian1337/thukul-service/business/wishlists/mocks"
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

var wishlistsRepository _wishlistMock.Repository
var userService _usersMock.Usecase

var wishlistService wishlists.Usecase

var wishlistDomain wishlists.Domain
var userDomain users.Domain
var listWishlistDomain []wishlists.Domain

var rowsAffectedSuccess int64
var rowsAffectedError int64

func setup() {
	rowsAffectedError = 0
	rowsAffectedSuccess = 1
	wishlistService = wishlists.NewWishlistUsecase(&wishlistsRepository, &userService, time.Second*10)
	wishlistDomain = wishlists.Domain{
		ID:     1,
		UserId: 1,
		Name:   "Kantong",
	}
	listWishlistDomain = append(listWishlistDomain, wishlistDomain)

	userDomain = users.Domain{
		ID:       1,
		SalaryId: 1,
		Name:     "Nama",
		IsAdmin:  0,
		Email:    "user@aprian1337.io",
		Phone:    "0812",
		Birthday: "2021-01-01",
		Address:  "Indonesia",
		Company:  "UMM",
		IsValid:  1,
	}
}

func TestGetList(t *testing.T) {
	t.Run("Test Case 1 | GetList - Success", func(t *testing.T) {
		setup()
		wishlistsRepository.On("WishlistsGetList", mock.Anything, mock.Anything).Return(listWishlistDomain, nil).Once()
		data, err := wishlistService.GetList(context.Background(), 1)

		assert.Nil(t, err)
		assert.Equal(t, data, listWishlistDomain)
	})

	t.Run("Test Case 2 | GetList - Error", func(t *testing.T) {
		setup()
		wishlistsRepository.On("WishlistsGetList", mock.Anything, mock.Anything).Return([]wishlists.Domain{}, businesses.ErrForTest).Once()
		data, err := wishlistService.GetList(context.Background(), 1)

		assert.Error(t, err)
		assert.Equal(t, data, []wishlists.Domain{})
	})

}

func TestGetById(t *testing.T) {
	t.Run("Test Case 1 | GetById - Success", func(t *testing.T) {
		setup()
		wishlistsRepository.On("WishlistsGetById", mock.Anything, mock.Anything, mock.Anything).Return(wishlistDomain, nil).Once()
		wishlist, err := wishlistService.GetById(context.Background(), wishlistDomain.UserId, wishlistDomain.ID)

		assert.NoError(t, err)
		assert.Equal(t, wishlist, wishlistDomain)
	})

	t.Run("Test Case 2 | GetById - Error", func(t *testing.T) {
		setup()
		wishlistsRepository.On("WishlistsGetById", mock.Anything, mock.Anything, mock.Anything).Return(wishlists.Domain{}, businesses.ErrForTest).Once()
		wishlist, err := wishlistService.GetById(context.Background(), wishlistDomain.UserId, wishlistDomain.ID)

		assert.Error(t, err)
		assert.Equal(t, wishlist, wishlists.Domain{})
	})
}

func TestDelete(t *testing.T) {
	t.Run("Test Case 1 | Delete - Success", func(t *testing.T) {
		setup()
		wishlistsRepository.On("WishlistsDelete",
			mock.Anything, mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(rowsAffectedSuccess, nil).Once()
		err := wishlistService.Delete(context.Background(), wishlistDomain.UserId, wishlistDomain.ID)

		assert.NoError(t, err)
		assert.Nil(t, err)
		wishlistsRepository.AssertExpectations(t)
	})

	t.Run("Test Case 2 | Delete - Error", func(t *testing.T) {
		setup()
		wishlistsRepository.On("WishlistsDelete",
			mock.Anything, mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(rowsAffectedSuccess, businesses.ErrForTest).Once()
		err := wishlistService.Delete(context.Background(), wishlistDomain.UserId, wishlistDomain.ID)

		assert.Error(t, err)
		wishlistsRepository.AssertExpectations(t)
	})

	t.Run("Test Case 3 | Delete - Error - No Found Data", func(t *testing.T) {
		setup()
		wishlistsRepository.On("WishlistsDelete",
			mock.Anything, mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(rowsAffectedError, nil).Once()
		err := wishlistService.Delete(context.Background(), wishlistDomain.UserId, wishlistDomain.ID)

		assert.Error(t, err)
		wishlistsRepository.AssertExpectations(t)
	})
}

func TestCreate(t *testing.T) {
	t.Run("Test Case 1 | Create - Success", func(t *testing.T) {
		setup()
		userService.On("GetById", mock.Anything, mock.Anything).Return(userDomain, nil).Once()
		wishlistsRepository.On("WishlistsCreate", mock.Anything, mock.Anything, mock.Anything).Return(wishlistDomain, nil).Once()
		data, err := wishlistService.Create(context.Background(), wishlistDomain, wishlistDomain.UserId)

		assert.NoError(t, err)
		assert.Equal(t, data, wishlistDomain)
		wishlistsRepository.AssertExpectations(t)
	})

	t.Run("Test Case 2 | Create - Error - UserId Not Found", func(t *testing.T) {
		setup()
		userService.On("GetById", mock.Anything, mock.Anything, mock.Anything).Return(users.Domain{}, businesses.ErrForTest).Once()
		data, err := wishlistService.Create(context.Background(), wishlistDomain, wishlistDomain.UserId)

		assert.Error(t, err)
		assert.Equal(t, data, wishlists.Domain{})
		wishlistsRepository.AssertExpectations(t)
	})

	t.Run("Test Case 3 | Create - Error - Fail Create Wishlists", func(t *testing.T) {
		setup()
		userService.On("GetById", mock.Anything, mock.Anything).Return(userDomain, nil).Once()
		wishlistsRepository.On("WishlistsCreate", mock.Anything, mock.Anything, mock.Anything).Return(wishlists.Domain{}, businesses.ErrForTest).Once()
		data, err := wishlistService.Create(context.Background(), wishlistDomain, wishlistDomain.UserId)

		assert.Error(t, err)
		assert.Equal(t, data, wishlists.Domain{})
		wishlistsRepository.AssertExpectations(t)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Test Case 1 | Update - Success", func(t *testing.T) {
		setup()
		wishlistsRepository.On("WishlistsUpdate", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(wishlistDomain, nil).Once()
		data, err := wishlistService.Update(context.Background(), wishlistDomain, wishlistDomain.UserId, wishlistDomain.ID)

		assert.NoError(t, err)
		assert.Equal(t, data, wishlistDomain)
		wishlistsRepository.AssertExpectations(t)
	})

	t.Run("Test Case 2 | Update - Error", func(t *testing.T) {
		setup()
		wishlistsRepository.On("WishlistsUpdate", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(wishlists.Domain{}, businesses.ErrForTest).Once()
		data, err := wishlistService.Update(context.Background(), wishlistDomain, wishlistDomain.UserId, wishlistDomain.ID)

		assert.Error(t, err)
		assert.Equal(t, data, wishlists.Domain{})
		wishlistsRepository.AssertExpectations(t)
	})
}
