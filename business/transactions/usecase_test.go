package transactions_test

import (
	"aprian1337/thukul-service/business"
	"aprian1337/thukul-service/business/transactions"
	_transactionMock "aprian1337/thukul-service/business/transactions/mocks"
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

var transactionRepository _transactionMock.Repository

var transactionService transactions.Usecase

var transactionDomain transactions.Domain

func setup() {
	transactionService = transactions.NewTransactionUsecase(&transactionRepository, time.Second*10)
	transactionDomain = transactions.Domain{
		UserId: 1,
		CoinId: 1,
		Qty:    1,
		Price:  1,
	}
}

func TestTransactionsById(t *testing.T) {
	t.Run("Test Case 1 | TransactionsById - Success", func(t *testing.T) {
		setup()
		transactionRepository.On("TransactionsById", mock.Anything, mock.Anything).Return(transactionDomain, nil).Once()
		data, err := transactionService.TransactionsById(context.Background(), "1")

		assert.Nil(t, err)
		assert.Equal(t, data, transactionDomain)
	})

	t.Run("Test Case 2 | TransactionsById - Error", func(t *testing.T) {
		setup()
		transactionRepository.On("TransactionsById", mock.Anything, mock.Anything).Return(transactions.Domain{}, businesses.ErrForTest).Once()
		data, err := transactionService.TransactionsById(context.Background(), "1")

		assert.Error(t, err)
		assert.Equal(t, data, transactions.Domain{})
	})

}

func TestTransactionsCreate(t *testing.T) {
	t.Run("Test Case 1 | TransactionsCreate - Success", func(t *testing.T) {
		setup()
		transactionRepository.On("TransactionsCreate", mock.Anything, mock.Anything).Return(transactionDomain, nil).Once()
		pocket, err := transactionService.TransactionsCreate(context.Background(), transactionDomain)

		assert.NoError(t, err)
		assert.Equal(t, pocket, transactionDomain)
	})

	t.Run("Test Case 2 | TransactionsCreate - Error", func(t *testing.T) {
		setup()
		transactionRepository.On("TransactionsCreate", mock.Anything, mock.Anything).Return(transactions.Domain{}, businesses.ErrForTest).Once()
		pocket, err := transactionService.TransactionsCreate(context.Background(), transactionDomain)

		assert.Error(t, err)
		assert.Equal(t, pocket, transactions.Domain{})
	})
}

func TestTransactionsUpdaterVerify(t *testing.T) {
	t.Run("Test Case 1 | TransactionsUpdaterVerify - Success", func(t *testing.T) {
		setup()
		transactionRepository.On("TransactionsUpdaterVerify",
			mock.Anything, mock.Anything, mock.Anything).Return(transactionDomain, nil).Once()
		data, err := transactionService.TransactionsUpdaterVerify(context.Background(), transactionDomain.Id)

		assert.NoError(t, err)
		assert.NotNil(t, data)
		transactionRepository.AssertExpectations(t)
	})

	t.Run("Test Case 2 | TransactionsUpdaterVerify - Error", func(t *testing.T) {
		setup()
		transactionRepository.On("TransactionsUpdaterVerify",
			mock.Anything, mock.Anything, mock.Anything).Return(transactionDomain, nil).Once()
		data, err := transactionService.TransactionsUpdaterVerify(context.Background(), transactionDomain.Id)

		assert.NoError(t, err)
		assert.NotNil(t, data)
		transactionRepository.AssertExpectations(t)
	})
}

func TestTransactionsUpdaterCompleted(t *testing.T) {
	t.Run("Test Case 1 | TransactionsUpdaterCompleted - Success", func(t *testing.T) {
		setup()
		transactionRepository.On("TransactionsUpdaterCompleted",
			mock.Anything, mock.Anything, mock.Anything).Return(transactionDomain, nil).Once()
		data, err := transactionService.TransactionsUpdaterCompleted(context.Background(), transactionDomain.Id, 2)

		assert.NoError(t, err)
		assert.NotNil(t, data)
		transactionRepository.AssertExpectations(t)
	})

	t.Run("Test Case 2 | TransactionsUpdaterCompleted - Update Error", func(t *testing.T) {
		setup()
		transactionRepository.On("TransactionsUpdaterCompleted",
			mock.Anything, mock.Anything, mock.Anything).Return(transactions.Domain{}, businesses.ErrForTest).Once()
		data, err := transactionService.TransactionsUpdaterCompleted(context.Background(), transactionDomain.Id, 2)

		assert.Error(t, err)
		assert.Equal(t, data, transactions.Domain{})
		transactionRepository.AssertExpectations(t)
	})

	t.Run("Test Case 3 | TransactionsUpdaterCompleted - Error - Status Must 1 Or 2", func(t *testing.T) {
		setup()
		data, err := transactionService.TransactionsUpdaterCompleted(context.Background(), transactionDomain.Id, 55)

		assert.Error(t, err)
		assert.Equal(t, data, transactions.Domain{})
		transactionRepository.AssertExpectations(t)
	})
}
