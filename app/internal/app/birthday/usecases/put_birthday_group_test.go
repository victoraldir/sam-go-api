package usecases

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestPutBirthdayUseCase(t *testing.T) {

	t.Run("Should return error when repository return error", func(t *testing.T) {
		setup(t)

		expectedError := fmt.Errorf("error")

		mirthdayRepositoryMock.EXPECT().PutBirthday(gomock.Any()).Return(expectedError)

		useCase := NewPutBirthDayUseCase(mirthdayRepositoryMock)

		_, err := useCase.Execute(PutBirthdayCommand{
			Username:    "username",
			DateOfBirth: "2006-01-02",
		})

		assert.Error(t, err)
		assert.Equal(t, expectedError.Error(), err.Error())
	})

	t.Run("Should return error response when username contains numbers", func(t *testing.T) {

		useCase := NewPutBirthDayUseCase(mirthdayRepositoryMock)

		response, err := useCase.Execute(PutBirthdayCommand{
			Username:    "username123",
			DateOfBirth: "2006-01-02",
		})

		assert.Nil(t, err)
		assert.Equal(t, "username must contain only letters", response.ErrorMsg)
		assert.Equal(t, 400, response.ErrorCode)

	})

	t.Run("Should return error response when username contains special characters", func(t *testing.T) {

		useCase := NewPutBirthDayUseCase(mirthdayRepositoryMock)

		response, err := useCase.Execute(PutBirthdayCommand{
			Username:    "username@",
			DateOfBirth: "2006-01-02",
		})

		assert.Nil(t, err)
		assert.Equal(t, "username must contain only letters", response.ErrorMsg)
		assert.Equal(t, 400, response.ErrorCode)

	})

	t.Run("Should return error response when username contains spaces", func(t *testing.T) {

		useCase := NewPutBirthDayUseCase(mirthdayRepositoryMock)

		response, err := useCase.Execute(PutBirthdayCommand{
			Username:    "user name",
			DateOfBirth: "2006-01-02",
		})

		assert.Nil(t, err)
		assert.Equal(t, "username must contain only letters", response.ErrorMsg)
		assert.Equal(t, 400, response.ErrorCode)

	})

	t.Run("Should return error response when date of birth is in the future", func(t *testing.T) {
		useCase := NewPutBirthDayUseCase(mirthdayRepositoryMock)

		dataFuture := time.Now().AddDate(1, 0, 0).Format("2006-01-02")

		response, err := useCase.Execute(PutBirthdayCommand{
			Username:    "username",
			DateOfBirth: dataFuture,
		})

		assert.Nil(t, err)

		assert.Equal(t, "date of birth must be before today", response.ErrorMsg)
		assert.Equal(t, 400, response.ErrorCode)
	})

	t.Run("Should return error response when date of birth is invalid", func(t *testing.T) {

		useCase := NewPutBirthDayUseCase(mirthdayRepositoryMock)

		response, err := useCase.Execute(PutBirthdayCommand{
			Username:    "username",
			DateOfBirth: "12-12-12",
		})

		assert.Nil(t, err)

		assert.Equal(t, "date of birth must be before today", response.ErrorMsg)
		assert.Equal(t, 400, response.ErrorCode)
	})

	t.Run("Should return error response when date of birth is in the future", func(t *testing.T) {
		useCase := NewPutBirthDayUseCase(mirthdayRepositoryMock)

		todayNextYear := time.Now().AddDate(1, 0, 0).Format("2006-01-02")
		tomorrow := time.Now().AddDate(0, 0, 1).Format("2006-01-02")

		dobList := []string{
			todayNextYear,
			tomorrow,
		}

		for _, dob := range dobList {
			response, err := useCase.Execute(PutBirthdayCommand{
				Username:    "username",
				DateOfBirth: dob,
			})

			assert.Nil(t, err)

			assert.Equal(t, "date of birth must be before today", response.ErrorMsg)
			assert.Equal(t, 400, response.ErrorCode)
		}

	})

	t.Run("Should return success response when date of birth is valid", func(t *testing.T) {
		setup(t)

		dobLastYearTomorrow := time.Now().AddDate(-1, 0, 1).Format("2006-01-02")

		mirthdayRepositoryMock.EXPECT().PutBirthday(gomock.Any()).Return(nil).AnyTimes()

		useCase := NewPutBirthDayUseCase(mirthdayRepositoryMock)

		dobList := []string{
			"2006-01-02",
			"1993-09-16",
			dobLastYearTomorrow,
		}

		for _, dob := range dobList {
			response, err := useCase.Execute(PutBirthdayCommand{
				Username:    "username",
				DateOfBirth: dob,
			})

			assert.Nil(t, err)

			assert.Equal(t, "Birthday saved!", response.Message)
			assert.Equal(t, 0, response.ErrorCode)
		}

	})

}
