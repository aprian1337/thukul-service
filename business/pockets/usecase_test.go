package pockets_test

import (
	"aprian1337/thukul-service/business"
	"aprian1337/thukul-service/business/activities"
	_activityUsecase "aprian1337/thukul-service/business/activities/mocks"
	"aprian1337/thukul-service/business/pockets"
	_pocketRepository "aprian1337/thukul-service/business/pockets/mocks"
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

var pocketsRepository _pocketRepository.Repository
var activityService _activityUsecase.Usecase

var pocketService pockets.Usecase

var pocketDomain pockets.Domain
var activityDomain activities.Domain
var listPocketDomain []pockets.Domain

var rowsAffectedSuccess int64
var rowsAffectedError int64

func setup() {
	rowsAffectedError = 0
	rowsAffectedSuccess = 1
	pocketService = pockets.NewPocketUsecase(&pocketsRepository, &activityService, time.Second*10)
	pocketDomain = pockets.Domain{
		ID:     1,
		UserId: 1,
		Name:   "Kantong",
	}
	listPocketDomain = append(listPocketDomain, pocketDomain)

	activityDomain = activities.Domain{
		ID:       1,
		PocketId: 1,
		Name:     "zizi",
		Type:     "expense",
		Nominal:  100,
	}
}

func TestGetList(t *testing.T) {
	t.Run("Test Case 1 | GetList - Success", func(t *testing.T) {
		setup()
		pocketsRepository.On("PocketsGetList", mock.Anything, mock.Anything).Return(listPocketDomain, nil).Once()
		data, err := pocketService.GetList(context.Background(), "1")

		assert.Nil(t, err)
		assert.Equal(t, data, listPocketDomain)
	})

	t.Run("Test Case 2 | GetList - Error - Invalid Id", func(t *testing.T) {
		setup()
		_, err := pocketService.GetList(context.Background(), "zz")

		assert.Error(t, err)
	})

	t.Run("Test Case 3 | GetList - Error Get Pocket List", func(t *testing.T) {
		setup()
		pocketsRepository.On("PocketsGetList", mock.Anything, mock.Anything).Return([]pockets.Domain{}, businesses.ErrForTest).Once()
		pocket, err := pocketService.GetList(context.Background(), "1")

		assert.Error(t, err)
		assert.Equal(t, pocket, []pockets.Domain{})
	})
}

func TestGetById(t *testing.T) {
	t.Run("Test Case 1 | GetById - Success", func(t *testing.T) {
		setup()
		pocketsRepository.On("PocketsGetById", mock.Anything, mock.Anything, mock.Anything).Return(pocketDomain, nil).Once()
		pocket, err := pocketService.GetById(context.Background(), pocketDomain.UserId, pocketDomain.ID)

		assert.NoError(t, err)
		assert.Equal(t, pocket, pocketDomain)
	})

	t.Run("Test Case 2 | GetById - Error", func(t *testing.T) {
		setup()
		pocketsRepository.On("PocketsGetById", mock.Anything, mock.Anything, mock.Anything).Return(pockets.Domain{}, businesses.ErrForTest).Once()
		pocket, err := pocketService.GetById(context.Background(), pocketDomain.UserId, pocketDomain.ID)

		assert.Error(t, err)
		assert.Equal(t, pocket, pockets.Domain{})
	})
}

func TestDelete(t *testing.T) {
	t.Run("Test Case 1 | Delete - Success", func(t *testing.T) {
		setup()
		pocketsRepository.On("PocketsDelete",
			mock.Anything, mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(rowsAffectedSuccess, nil).Once()
		err := pocketService.Delete(context.Background(), pocketDomain.UserId, pocketDomain.ID)

		assert.NoError(t, err)
		assert.Nil(t, err)
		pocketsRepository.AssertExpectations(t)
	})

	t.Run("Test Case 2 | Delete - Error", func(t *testing.T) {
		setup()
		pocketsRepository.On("PocketsDelete",
			mock.Anything, mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(rowsAffectedSuccess, businesses.ErrForTest).Once()
		err := pocketService.Delete(context.Background(), pocketDomain.UserId, pocketDomain.ID)

		assert.Error(t, err)
		pocketsRepository.AssertExpectations(t)
	})

	t.Run("Test Case 2 | Delete - Error - No Found Data", func(t *testing.T) {
		setup()
		pocketsRepository.On("PocketsDelete",
			mock.Anything, mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(rowsAffectedError, nil).Once()
		err := pocketService.Delete(context.Background(), pocketDomain.UserId, pocketDomain.ID)

		assert.Error(t, err)
		pocketsRepository.AssertExpectations(t)
	})
}

func TestCreate(t *testing.T) {
	t.Run("Test Case 1 | Create - Success", func(t *testing.T) {
		setup()
		pocketsRepository.On("PocketsCreate", mock.Anything, mock.Anything).Return(pocketDomain, nil).Once()
		data, err := pocketService.Create(context.Background(), pocketDomain)

		assert.NoError(t, err)
		assert.Equal(t, data, pocketDomain)
		pocketsRepository.AssertExpectations(t)
	})
}

func TestGetTotalByActivities(t *testing.T) {
	t.Run("Test Case 1 | GetTotalByActivities - Success", func(t *testing.T) {
		setup()
		pocketsRepository.On("PocketsGetById", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(pocketDomain, nil).Once()
		activityService.On("GetTotal", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(rowsAffectedSuccess, nil).Once()
		_, err := pocketService.GetTotalByActivities(context.Background(), 1, 1, "income")

		assert.NoError(t, err)
		pocketsRepository.AssertExpectations(t)
	})

	t.Run("Test Case 2 | GetTotalByActivities - Error - Required Income, Expense, Profit", func(t *testing.T) {
		setup()
		_, err := pocketService.GetTotalByActivities(context.Background(), 1, 1, "zz")

		assert.Error(t, err)
		pocketsRepository.AssertExpectations(t)
	})

	t.Run("Test Case 3 | GetTotalByActivities - Error - UserId or Pocket Not Found", func(t *testing.T) {
		setup()
		pocketsRepository.On("PocketsGetById", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(pockets.Domain{}, businesses.ErrForTest).Once()
		_, err := pocketService.GetTotalByActivities(context.Background(), 1, 1, "expense")

		assert.Error(t, err)
		pocketsRepository.AssertExpectations(t)
	})
}
