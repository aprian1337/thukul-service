package salaries_test

import (
	"aprian1337/thukul-service/business/salaries"
	_mockSalariesRepository "aprian1337/thukul-service/business/salaries/mocks"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

var salaryRepository _mockSalariesRepository.Repository
var salaryService salaries.Usecase
var salaryDomain salaries.Domain
var salariesDomain []salaries.Domain
var rowsSuccessAffected int64
var rowsErrAffected int64

func setup() {
	salaryService = salaries.NewSalaryUsecase(&salaryRepository, time.Second*10)
	salaryDomain = salaries.Domain{
		ID:      1,
		Minimal: 1000,
		Maximal: 100000,
	}
	rowsSuccessAffected = 1
	rowsErrAffected = 0
	salariesDomain = append(salariesDomain, salaryDomain)
}

func TestGetById(t *testing.T) {
	t.Run("Test Case 1 | Get Salary By Id - Success", func(t *testing.T) {
		setup()
		salaryRepository.On("SalariesGetById",
			mock.Anything, mock.AnythingOfType("uint")).Return(salaryDomain, nil).Once()
		data, err := salaryService.GetById(context.Background(), salaryDomain.ID)

		assert.NoError(t, err)
		assert.NotNil(t, data)

		salaryRepository.AssertExpectations(t)
	})

	t.Run("Test Case 2 | Get Salary By Id - Error", func(t *testing.T) {
		setup()
		salaryRepository.On("SalariesGetById",
			mock.Anything, mock.AnythingOfType("uint")).Return(salaries.Domain{}, errors.New("error")).Once()
		data, err := salaryService.GetById(context.Background(), salaryDomain.ID)

		assert.Error(t, err)
		assert.NotNil(t, data)

		salaryRepository.AssertExpectations(t)
	})
}

func TestGetList(t *testing.T) {
	t.Run("Test Case 1 | Get All (Success)", func(t *testing.T) {
		setup()
		salaryRepository.On("SalariesGetList",
			mock.Anything).Return(salariesDomain, nil).Once()
		data, err := salaryService.GetList(context.Background())

		assert.NoError(t, err)
		assert.NotNil(t, data)

		salaryRepository.AssertExpectations(t)
	})

	t.Run("Test Case 2 | Get All (Error)", func(t *testing.T) {
		setup()
		salaryRepository.On("SalariesGetList",
			mock.Anything).Return([]salaries.Domain{}, errors.New("error")).Once()
		data, err := salaryService.GetList(context.Background())

		assert.Error(t, err)
		assert.NotNil(t, data)

		salaryRepository.AssertExpectations(t)
	})
}

func TestCreate(t *testing.T) {
	setup()
	t.Run("Test Case 1 | Create (Success)", func(t *testing.T) {
		salaryRepository.On("SalariesCreate",
			mock.Anything, mock.AnythingOfType("salaries.Domain")).Return(salaryDomain, nil).Once()
		data, err := salaryService.Create(context.Background(), salaryDomain)

		assert.Nil(t, err)
		assert.NotNil(t, data)

		salaryRepository.AssertExpectations(t)
	})

	t.Run("Test Case 2 | Create (Error)", func(t *testing.T) {
		salaryRepository.On("SalariesCreate",
			mock.Anything, mock.AnythingOfType("salaries.Domain")).Return(salaries.Domain{}, errors.New("error")).Once()
		data, err := salaryService.Create(context.Background(), salaries.Domain{})

		assert.Error(t, err)
		assert.Equal(t, data, salaries.Domain{})

		salaryRepository.AssertExpectations(t)
	})

}

func TestDelete(t *testing.T) {
	setup()
	t.Run("Test Case 1 | Delete (Success)", func(t *testing.T) {
		salaryRepository.On("SalariesDelete",
			mock.Anything, mock.AnythingOfType("uint")).Return(rowsSuccessAffected, nil).Once()
		err := salaryService.Delete(context.Background(), salaryDomain.ID)

		assert.Nil(t, err)

		salaryRepository.AssertExpectations(t)
	})

	t.Run("Test Case 2 | Delete (Error)", func(t *testing.T) {
		salaryRepository.On("SalariesDelete",
			mock.Anything, mock.AnythingOfType("uint")).Return(rowsErrAffected, errors.New("error")).Once()
		err := salaryService.Delete(context.Background(), salaryDomain.ID)

		assert.NotNil(t, err)

		salaryRepository.AssertExpectations(t)
	})

}

func TestUpdate(t *testing.T) {
	setup()
	t.Run("Test Case 1 | Update (Success)", func(t *testing.T) {
		salaryRepository.On("SalariesUpdate",
			mock.Anything, mock.AnythingOfType("salaries.Domain")).Return(salaryDomain, nil).Once()
		data, err := salaryService.Update(context.Background(), salaryDomain.ID, salaryDomain)

		assert.NotNil(t, data)
		assert.Equal(t, data.ID, salaryDomain.ID)
		assert.Nil(t, err)

		salaryRepository.AssertExpectations(t)
	})

	t.Run("Test Case 2 | Update (Error)", func(t *testing.T) {
		salaryRepository.On("SalariesUpdate",
			mock.Anything, mock.AnythingOfType("salaries.Domain")).Return(salaries.Domain{}, errors.New("error")).Once()
		data, err := salaryService.Update(context.Background(), 2, salaryDomain)

		assert.NotNil(t, err)
		assert.Error(t, err)
		assert.Equal(t, data, salaries.Domain{})

		salaryRepository.AssertExpectations(t)
	})

}
