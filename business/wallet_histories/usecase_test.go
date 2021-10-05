package wallet_histories_test

import (
	"aprian1337/thukul-service/business"
	"aprian1337/thukul-service/business/wallet_histories"
	_walletHistoriesMock "aprian1337/thukul-service/business/wallet_histories/mocks"
	"context"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

var walletHistoriesRepository _walletHistoriesMock.Repository

var walletHistoriesService wallet_histories.Usecase

var walletHistoriesDomain wallet_histories.Domain

func setup() {
	transactionId, _ := uuid.Parse("25765868-05d4-48e7-aed2-f4457a122ef7")
	walletHistoriesService = wallet_histories.NewWalletsUsecase(&walletHistoriesRepository, time.Second*10)
	walletHistoriesDomain = wallet_histories.Domain{
		ID:            1,
		WalletId:      1,
		TransactionId: &transactionId,
		Nominal:       1,
	}
}

func TestWalletHistoriesCreate(t *testing.T) {
	t.Run("Test Case 1 | WalletHistoriesCreate - Success", func(t *testing.T) {
		setup()
		walletHistoriesRepository.On("WalletHistoriesCreate", mock.Anything, mock.Anything).Return(nil).Once()
		err := walletHistoriesService.WalletHistoriesCreate(context.Background(), walletHistoriesDomain)

		assert.NoError(t, err)
	})

	t.Run("Test Case 2 | WalletHistoriesCreate - Error", func(t *testing.T) {
		setup()
		walletHistoriesRepository.On("WalletHistoriesCreate", mock.Anything, mock.Anything).Return(businesses.ErrForTest).Once()
		err := walletHistoriesService.WalletHistoriesCreate(context.Background(), walletHistoriesDomain)

		assert.Error(t, err)
	})

}
