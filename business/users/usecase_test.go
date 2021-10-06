package users_test

import (
	"aprian1337/thukul-service/app/middlewares"
	businesses "aprian1337/thukul-service/business"
	"aprian1337/thukul-service/business/users"
	_usersMockRepository "aprian1337/thukul-service/business/users/mocks"
	_walletMock "aprian1337/thukul-service/business/wallets/mocks"
	"aprian1337/thukul-service/business/wallets"
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

var userRepository _usersMockRepository.Repository
var userService users.Usecase
var walletService _walletMock.Usecase

var userDomain users.Domain
var listUserDomain []users.Domain
var walletDomain wallets.Domain
var listWalletDomain []wallets.Domain
var token string

func setup() {
	userService = users.NewUserUsecase(&userRepository, &walletService, time.Second*10, &middlewares.ConfigJWT{})
	userDomain = users.Domain{
		ID:       1,
		SalaryId: 1,
		Name:     "Nama",
		Password: "$2a$12$SaZVHXYiMiygeHoVV33cY..FwaM/oFNO9EVTnscfnKOlKvIWtnRRS",
		IsAdmin:  0,
		Email:    "user@aprian1337.io",
		Phone:    "0812",
		Salary: users.Salary{
			ID:      1,
			Minimal: 100,
			Maximal: 1000,
		},
		Wallets: users.Wallets{
			ID:    1,
			Total: 100,
		},
		Gender:   "M",
		Birthday: "2021-01-01",
		Address:  "Indonesia",
		Company:  "UMM",
		IsValid:  1,
	}
	token = "token"
	listUserDomain = append(listUserDomain, userDomain)

	walletDomain = wallets.Domain{
		Id:                 1,
		UserId:             1,
		Total:              1000,
		NominalTransaction: 100,
	}
}

func TestGetByIdWithWallet(t *testing.T) {
	t.Run("Test Case 1 | GetByIdWithWallet - Success", func(t *testing.T) {
		setup()
		userRepository.On("UsersGetByIdWithWallet",
			mock.Anything, mock.AnythingOfType("int")).Return(userDomain, nil).Once()
		data, err := userService.GetByIdWithWallet(context.Background(), int(userDomain.ID))

		assert.NoError(t, err)
		assert.NotNil(t, data)

		userRepository.AssertExpectations(t)
	})

	t.Run("Test Case 2 | GetByIdWithWallet - Fail", func(t *testing.T) {
		setup()
		userRepository.On("UsersGetByIdWithWallet",
			mock.Anything, mock.AnythingOfType("int")).Return(users.Domain{}, businesses.ErrForTest).Once()
		data, err := userService.GetByIdWithWallet(context.Background(), int(userDomain.ID+1))

		assert.Error(t, err)
		assert.Equal(t, data, users.Domain{})

		userRepository.AssertExpectations(t)
	})
}

func TestGetAll(t *testing.T) {
	t.Run("Test Case 1 | GetAll - Success", func(t *testing.T) {
		setup()
		userRepository.On("UsersGetAll",
			mock.Anything).Return(listUserDomain, nil).Once()
		data, err := userService.GetAll(context.Background())

		assert.NoError(t, err)
		assert.NotNil(t, data)
		assert.Equal(t, len(data), len(listUserDomain))
		userRepository.AssertExpectations(t)
	})

	t.Run("Test Case 2| UsersGetAll - Error", func(t *testing.T) {
		setup()
		userRepository.On("UsersGetAll",
			mock.Anything).Return([]users.Domain{}, businesses.ErrForTest).Once()
		data, err := userService.GetAll(context.Background())

		assert.Error(t, err)
		assert.Equal(t, data, []users.Domain{})
		userRepository.AssertExpectations(t)
	})
}

func TestGetById(t *testing.T) {
	t.Run("Test Case 1 | GetById - Success", func(t *testing.T) {
		setup()
		userRepository.On("UsersGetById",
			mock.Anything, mock.AnythingOfType("int")).Return(userDomain, nil).Once()
		data, err := userService.GetById(context.Background(), int(userDomain.ID))

		assert.NoError(t, err)
		assert.NotNil(t, data)

		userRepository.AssertExpectations(t)
	})

	t.Run("Test Case 2 | GetById - Error - User Id 0", func(t *testing.T) {
		setup()
		userDomain.ID = 0
		userRepository.On("UsersGetById",
			mock.Anything, mock.AnythingOfType("int")).Return(userDomain, nil).Once()
		data, err := userService.GetById(context.Background(), 7)

		assert.Error(t, err)
		assert.Equal(t, data, users.Domain{})

		userRepository.AssertExpectations(t)
	})

	t.Run("Test Case 3 | GetById - Error", func(t *testing.T) {
		setup()
		userRepository.On("UsersGetById",
			mock.Anything, mock.AnythingOfType("int")).Return(users.Domain{}, nil).Once()
		data, err := userService.GetById(context.Background(), int(userDomain.ID))

		assert.Error(t, err)
		assert.Equal(t, data, users.Domain{})

		userRepository.AssertExpectations(t)
	})
}

func TestLogin(t *testing.T) {
	t.Run("Test Case 1 | Login - Success", func(t *testing.T) {
		setup()
		userRepository.On("UsersGetByEmail",
			mock.Anything, mock.AnythingOfType("string")).Return(userDomain, nil).Once()
		data, token, err := userService.Login(context.Background(), userDomain.Email, "123")

		assert.NotNil(t, token)
		assert.NoError(t, err)
		assert.Equal(t, data, userDomain)

		userRepository.AssertExpectations(t)
	})

	t.Run("Test Case 2 | Login - Error (Username/Pass Not Found)", func(t *testing.T) {
		setup()
		data, token, err := userService.Login(context.Background(), userDomain.Email, "")

		assert.Equal(t, users.Domain{}, data)
		assert.Error(t, err)
		assert.Equal(t, token, "")

		userRepository.AssertExpectations(t)
	})

	t.Run("Test Case 3 | Login - Error (Wrong Auth)", func(t *testing.T) {
		setup()
		userRepository.On("UsersGetByEmail",
			mock.Anything, mock.AnythingOfType("string")).Return(users.Domain{}, businesses.ErrForTest).Once()
		data, token, err := userService.Login(context.Background(), userDomain.Email, "1234")

		assert.Equal(t, users.Domain{}, data)
		assert.Error(t, err)
		assert.Equal(t, token, "")

		userRepository.AssertExpectations(t)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Test Case 1 | Update - Success", func(t *testing.T) {
		setup()
		userRepository.On("UsersGetByEmail",
			mock.Anything, mock.AnythingOfType("string")).Return(userDomain, nil).Once()
		userRepository.On("UsersUpdate",
			mock.Anything, mock.AnythingOfType("*users.Domain")).Return(userDomain, nil).Once()
		data, err := userService.Update(context.Background(), &userDomain, userDomain.ID)

		assert.NotNil(t, data)
		assert.NoError(t, err)

		userRepository.AssertExpectations(t)
	})

	t.Run("Test Case 2 | Update - Error (Update Fail)", func(t *testing.T) {
		setup()
		userRepository.On("UsersGetByEmail",
			mock.Anything, mock.AnythingOfType("string")).Return(userDomain, nil).Once()
		userRepository.On("UsersUpdate",
			mock.Anything, mock.AnythingOfType("*users.Domain")).Return(users.Domain{}, businesses.ErrForTest).Once()
		data, err := userService.Update(context.Background(), &userDomain, userDomain.ID)

		assert.Equal(t, users.Domain{}, data)
		assert.Error(t, err)

		userRepository.AssertExpectations(t)
	})

}
func TestDelete(t *testing.T) {
	t.Run("Test Case 1 | Delete - Success", func(t *testing.T) {
		setup()
		userRepository.On("UsersDelete",
			mock.Anything, mock.AnythingOfType("uint")).Return(nil).Once()
		err := userService.Delete(context.Background(), userDomain.ID)

		assert.Nil(t, err)

		userRepository.AssertExpectations(t)
	})

	t.Run("Test Case 2 | Delete - Error", func(t *testing.T) {
		setup()
		userRepository.On("UsersDelete",
			mock.Anything, mock.AnythingOfType("uint")).Return(businesses.ErrForTest).Once()
		err := userService.Delete(context.Background(), userDomain.ID)

		assert.Equal(t, err, businesses.ErrForTest)
		assert.Error(t, err)

		userRepository.AssertExpectations(t)
	})

}

func TestCreate(t *testing.T) {
	t.Run("Test Case 1 | Create - Success", func(t *testing.T) {
		setup()
		userRepository.On("UsersGetByEmail",
			mock.Anything, mock.AnythingOfType("string")).Return(users.Domain{}, nil).Once()
		userRepository.On("UsersCreate",
			mock.Anything, mock.Anything).Return(userDomain, nil).Once()
		walletService.On("Create", mock.Anything, mock.Anything).Return(nil)

		data, err := userService.Create(context.Background(), &users.Domain{
			SalaryId: 1,
			Name:     "Nama",
			Password: "123",
			Email:    "user@aprian1337.io",
			Phone:    "0812",
			Gender:   "M",
			Birthday: "2021-01-01",
			Address:  "Indonesia",
			Company:  "UMM",
		})

		assert.Nil(t, err)
		assert.Equal(t, data, userDomain)
		userRepository.AssertExpectations(t)
	})

	t.Run("Test Case 2 | Error - Email Required", func(t *testing.T) {
		setup()
		data, err := userService.Create(context.Background(), &users.Domain{
			SalaryId: 1,
			Name:     "Nama",
			Password: "123",
			Phone:    "0812",
			Gender:   "M",
			Birthday: "2021-01-01",
			Address:  "Indonesia",
			Company:  "UMM",
		})

		assert.Error(t, err)
		assert.Equal(t, data, users.Domain{})
		userRepository.AssertExpectations(t)
	})

	t.Run("Test Case 3 | Error - Is Not Valid Email", func(t *testing.T) {
		setup()
		data, err := userService.Create(context.Background(), &users.Domain{
			SalaryId: 1,
			Name:     "Nama",
			Password: "123",
			Email:    "useraprian1337.io",
			Phone:    "0812",
			Gender:   "M",
			Birthday: "2021-01-01",
			Address:  "Indonesia",
			Company:  "UMM",
		})

		assert.Error(t, err)
		assert.Equal(t, data, users.Domain{})
		userRepository.AssertExpectations(t)
	})

	t.Run("Test Case 3 | Error - Email Has Been Used", func(t *testing.T) {
		setup()
		userRepository.On("UsersGetByEmail",
			mock.Anything, mock.AnythingOfType("string")).Return(userDomain, nil).Once()
		data, err := userService.Create(context.Background(), &users.Domain{
			SalaryId: 1,
			Name:     "Nama",
			Email:    "user@aprian1337.io",
			Phone:    "0812",
			Gender:   "M",
			Birthday: "2021-01-01",
			Address:  "Indonesia",
			Company:  "UMM",
		})

		assert.Error(t, err)
		assert.Equal(t, data, users.Domain{})
		userRepository.AssertExpectations(t)
	})

	t.Run("Test Case 3 | Error - Password Required", func(t *testing.T) {
		setup()
		userRepository.On("UsersGetByEmail",
			mock.Anything, mock.AnythingOfType("string")).Return(users.Domain{}, nil).Once()
		data, err := userService.Create(context.Background(), &users.Domain{
			SalaryId: 1,
			Name:     "Nama",
			Email:    "user@aprian1337.io",
			Phone:    "0812",
			Gender:   "M",
			Birthday: "2021-01-01",
			Address:  "Indonesia",
			Company:  "UMM",
		})

		assert.Error(t, err)
		assert.Equal(t, data, users.Domain{})
		userRepository.AssertExpectations(t)
	})

	t.Run("Test Case 3 | Error - Date Not Valid", func(t *testing.T) {
		setup()
		userRepository.On("UsersGetByEmail",
			mock.Anything, mock.AnythingOfType("string")).Return(users.Domain{}, nil).Once()
		data, err := userService.Create(context.Background(), &users.Domain{
			SalaryId: 1,
			Name:     "Nama",
			Email:    "user@aprian1337.io",
			Phone:    "0812",
			Gender:   "M",
			Password: "secret#1x",
			Birthday: "2021-01-022zz1",
			Address:  "Indonesia",
			Company:  "UMM",
		})

		assert.Error(t, err)
		assert.Equal(t, data, users.Domain{})
		userRepository.AssertExpectations(t)
	})
}
