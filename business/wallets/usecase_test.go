package wallets_test

import (
	"aprian1337/thukul-service/business"
	"aprian1337/thukul-service/business/wallet_histories"
	_walletHistoriesMock "aprian1337/thukul-service/business/wallet_histories/mocks"
	"aprian1337/thukul-service/business/wallets"
	_walletMock "aprian1337/thukul-service/business/wallets/mocks"
	"context"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

var walletRepository _walletMock.Repository

var walletService wallets.Usecase
var walletHistoryService _walletHistoriesMock.Usecase

var walletDomain wallets.Domain
var walletHistoriesDomain wallet_histories.Domain

func setup() {
	transactionId, _ := uuid.Parse("25765868-05d4-48e7-aed2-f4457a122ef7")
	walletService = wallets.NewWalletsUsecase(&walletRepository, &walletHistoryService, time.Second*10)
	walletDomain = wallets.Domain{
		Id:                 1,
		UserId:             1,
		Total:              1,
		NominalTransaction: 1,
		Kind:               "",
		TransactionId:      &transactionId,
		CoinId:             1,
	}
}

func TestGetByUserId(t *testing.T) {
	t.Run("Test Case 1 | GetByUserId - Success", func(t *testing.T) {
		setup()
		walletRepository.On("GetByUserId", mock.Anything, mock.Anything).Return(walletDomain, nil).Once()
		data, err := walletService.GetByUserId(context.Background(), walletDomain.UserId)

		assert.NoError(t, err)
		assert.Equal(t, data, walletDomain)
	})

	t.Run("Test Case 2 | GetByUserId - Error", func(t *testing.T) {
		setup()
		walletRepository.On("GetByUserId", mock.Anything, mock.Anything).Return(wallets.Domain{}, businesses.ErrForTest).Once()
		data, err := walletService.GetByUserId(context.Background(), walletDomain.UserId)

		assert.Equal(t, data, wallets.Domain{})
		assert.Error(t, err)
	})

}

func TestCreate(t *testing.T) {
	t.Run("Test Case 1 | Create - Success", func(t *testing.T) {
		setup()
		walletRepository.On("Create", mock.Anything, mock.Anything).Return(nil).Once()
		err := walletService.Create(context.Background(), walletDomain)

		assert.NoError(t, err)
	})

	t.Run("Test Case 1 | Create - Success", func(t *testing.T) {
		setup()
		walletRepository.On("Create", mock.Anything, mock.Anything).Return(businesses.ErrForTest).Once()
		err := walletService.Create(context.Background(), walletDomain)

		assert.Error(t, err)
	})

}

func TestUpdateByUserId(t *testing.T) {
	t.Run("Test Case 1 | UpdateByUserId - Success", func(t *testing.T) {
		setup()
		walletRepository.On("UpdateByUserId", mock.Anything, mock.Anything).Return(walletDomain, nil).Once()
		data, err := walletService.UpdateByUserId(context.Background(), walletDomain)

		assert.NoError(t, err)
		assert.Equal(t, data, walletDomain)
	})

	t.Run("Test Case 2 | UpdateByUserId - Error", func(t *testing.T) {
		setup()
		walletRepository.On("UpdateByUserId", mock.Anything, mock.Anything).Return(wallets.Domain{}, businesses.ErrForTest).Once()
		data, err := walletService.UpdateByUserId(context.Background(), walletDomain)

		assert.Error(t, err)
		assert.Equal(t, data, wallets.Domain{})
	})

	t.Run("Test Case 3 | Topup - Error", func(t *testing.T) {
		setup()
		walletRepository.On("UpdateByUserId", mock.Anything, mock.Anything).Return(walletDomain, nil).Once()
		walletHistoryService.On("WalletHistoriesCreate", mock.Anything, mock.Anything).Return(businesses.ErrForTest)
		walletDomain.Kind = "topup"
		data, err := walletService.UpdateByUserId(context.Background(), walletDomain)

		assert.Error(t, err)
		assert.Equal(t, data, wallets.Domain{})
	})

}
