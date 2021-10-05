package activities_test

import (
	businesses "aprian1337/thukul-service/business"
	"aprian1337/thukul-service/business/activities"
	_mockActivitiesRepository "aprian1337/thukul-service/business/activities/mocks"
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

var activitiyRepository _mockActivitiesRepository.Repository
var activitiyService activities.Usecase
var activitiyDomain activities.Domain
var listActivityDomain []activities.Domain
var rowsAffectedSuccess int64
var rowsAffectedError int64
var totalSuccess int64
var totalErr int64

func setup() {
	activitiyService = activities.NewActivityUsecase(&activitiyRepository, time.Second*10)
	activitiyDomain = activities.Domain{
		ID:       1,
		PocketId: 1,
		Name:     "Beli",
		Type:     "income",
		Nominal:  1000,
		Note:     "Abc",
		Date:     "2021-01-01",
	}
	totalSuccess = 1
	totalErr = 0
	rowsAffectedSuccess = 1
	rowsAffectedError = 0
	listActivityDomain = append(listActivityDomain, activitiyDomain)
}

func TestGetList(t *testing.T) {
	t.Run("Test Case 1 | GetList - Success", func(t *testing.T) {
		setup()
		activitiyRepository.On("ActivitiesGetAll",
			mock.Anything, mock.AnythingOfType("int")).Return(listActivityDomain, nil).Once()
		data, err := activitiyService.GetList(context.Background(), activitiyDomain.PocketId)

		assert.NoError(t, err)
		assert.NotNil(t, data)

		activitiyRepository.AssertExpectations(t)
	})

	t.Run("Test Case 2 | GetList - Error", func(t *testing.T) {
		setup()
		activitiyRepository.On("ActivitiesGetAll",
			mock.Anything, mock.AnythingOfType("int")).Return([]activities.Domain{}, businesses.ErrForTest).Once()
		data, err := activitiyService.GetList(context.Background(), activitiyDomain.PocketId)

		assert.Error(t, err)
		assert.Equal(t, data, []activities.Domain{})

		activitiyRepository.AssertExpectations(t)
	})
}

func TestGetById(t *testing.T) {
	t.Run("Test Case 1 | GetById - Success", func(t *testing.T) {
		setup()
		activitiyRepository.On("ActivitiesGetById",
			mock.Anything, mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(activitiyDomain, nil).Once()
		data, err := activitiyService.GetById(context.Background(), activitiyDomain.PocketId, activitiyDomain.ID)

		assert.NoError(t, err)
		assert.NotNil(t, data)

		activitiyRepository.AssertExpectations(t)
	})

	t.Run("Test Case 2 | GetById - Error", func(t *testing.T) {
		setup()
		activitiyRepository.On("ActivitiesGetById",
			mock.Anything, mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(activities.Domain{}, businesses.ErrForTest).Once()
		data, err := activitiyService.GetById(context.Background(), activitiyDomain.PocketId, activitiyDomain.ID)

		assert.Error(t, err)
		assert.Equal(t, data, activities.Domain{})

		activitiyRepository.AssertExpectations(t)
	})
}

func TestCreate(t *testing.T) {
	t.Run("Test Case 1 | Create - Success", func(t *testing.T) {
		setup()
		activitiyRepository.On("ActivitiesCreate",
			mock.Anything, mock.AnythingOfType("activities.Domain"), mock.AnythingOfType("int")).Return(activitiyDomain, nil).Once()
		data, err := activitiyService.Create(context.Background(), activitiyDomain, activitiyDomain.PocketId)

		assert.NoError(t, err)
		assert.NotNil(t, data)

		activitiyRepository.AssertExpectations(t)
	})

	t.Run("Test Case 2 | GetById - Error", func(t *testing.T) {
		setup()
		activitiyRepository.On("ActivitiesCreate",
			mock.Anything, mock.AnythingOfType("activities.Domain"), mock.AnythingOfType("int")).Return(activities.Domain{}, businesses.ErrForTest).Once()
		data, err := activitiyService.Create(context.Background(), activitiyDomain, activitiyDomain.PocketId)

		assert.Error(t, err)
		assert.Equal(t, data, activities.Domain{})

		activitiyRepository.AssertExpectations(t)
	})

	t.Run("Test Case 3 | GetById - Name / Nominal / Date Required", func(t *testing.T) {
		setup()
		data := activities.Domain{
			PocketId: 2,
			Name:     "",
			Type:     "",
			Date:     "",
		}
		data, err := activitiyService.Create(context.Background(), data, data.PocketId)

		assert.Error(t, err)
		assert.Equal(t, data, activities.Domain{})

		activitiyRepository.AssertExpectations(t)
	})

	t.Run("Test Case 4 | GetById - Type Must Income / Expense", func(t *testing.T) {
		setup()
		data := activities.Domain{
			PocketId: 2,
			Name:     "asda",
			Type:     "dasdsa",
			Date:     "2020-01-01",
			Nominal:  1111,
		}
		data, err := activitiyService.Create(context.Background(), data, data.PocketId)

		assert.Error(t, err)
		assert.Equal(t, data, activities.Domain{})

		activitiyRepository.AssertExpectations(t)
	})
	t.Run("Test Case 5 | GetById - Date Data Type", func(t *testing.T) {
		setup()
		data := activities.Domain{
			PocketId: 2,
			Name:     "asda",
			Type:     "expense",
			Date:     "20z20-01-01",
			Nominal:  1111,
		}
		data, err := activitiyService.Create(context.Background(), data, data.PocketId)

		assert.Error(t, err)
		assert.Equal(t, data, activities.Domain{})

		activitiyRepository.AssertExpectations(t)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Test Case 1 | Update - Success", func(t *testing.T) {
		setup()
		activitiyRepository.On("ActivitiesUpdate",
			mock.Anything, mock.AnythingOfType("activities.Domain"), mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(activitiyDomain, nil).Once()
		data, err := activitiyService.Update(context.Background(), activitiyDomain, activitiyDomain.PocketId, activitiyDomain.ID)

		assert.NoError(t, err)
		assert.NotNil(t, data)

		activitiyRepository.AssertExpectations(t)
	})

	t.Run("Test Case 2 | Update - Fail - Date Not Formed", func(t *testing.T) {
		setup()
		activitiyDomain.Date = "kzlkzl"
		data, err := activitiyService.Update(context.Background(), activitiyDomain, activitiyDomain.PocketId, activitiyDomain.ID)

		assert.Error(t, err)
		assert.Equal(t, activities.Domain{}, data)

		activitiyRepository.AssertExpectations(t)
	})
}

func TestDelete(t *testing.T) {
	t.Run("Test Case 1 | Delete - Success", func(t *testing.T) {
		setup()
		activitiyRepository.On("ActivitiesDelete",
			mock.Anything, mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(rowsAffectedSuccess, nil).Once()
		err := activitiyService.Delete(context.Background(), activitiyDomain.ID, activitiyDomain.PocketId)

		assert.NoError(t, err)

		activitiyRepository.AssertExpectations(t)
	})

	t.Run("Test Case 2 | Update - Delete - Data Not Found", func(t *testing.T) {
		setup()
		activitiyRepository.On("ActivitiesDelete",
			mock.Anything, mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(rowsAffectedError, nil).Once()
		err := activitiyService.Delete(context.Background(), activitiyDomain.ID, activitiyDomain.PocketId)

		assert.Error(t, err)

		activitiyRepository.AssertExpectations(t)
	})
}

func TestGetTotal(t *testing.T) {
	t.Run("Test Case 1 | GetTotal - Success", func(t *testing.T) {
		setup()
		activitiyRepository.On("ActivitiesGetTotal",
			mock.Anything, mock.AnythingOfType("int"), mock.AnythingOfType("int"), mock.AnythingOfType("string")).Return(totalSuccess, nil).Once()
		total, err := activitiyService.GetTotal(context.Background(), activitiyDomain.ID, activitiyDomain.PocketId, "expense")

		assert.NoError(t, err)
		assert.Equal(t, total, totalSuccess)
		activitiyRepository.AssertExpectations(t)
	})

	t.Run("Test Case 2 | Update - Delete - Data Not Found", func(t *testing.T) {
		setup()
		activitiyRepository.On("ActivitiesGetTotal",
			mock.Anything, mock.AnythingOfType("int"), mock.AnythingOfType("int"), mock.AnythingOfType("string")).Return(totalErr, businesses.ErrForTest).Once()
		total, err := activitiyService.GetTotal(context.Background(), activitiyDomain.ID, activitiyDomain.PocketId, "expense")
		assert.Equal(t, total, totalErr)
		assert.Error(t, err)
		activitiyRepository.AssertExpectations(t)
	})
}
