package usecases

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetBirthdayUseCase_Execute(t *testing.T) {
	t.Run("Should return error when repository return error", func(t *testing.T) {
		setup(t)

		expectedError := fmt.Errorf("error")

		mirthdayRepositoryMock.EXPECT().GetBirthday("username").Return("", expectedError)

		useCase := NewGetBirthDayUseCase(mirthdayRepositoryMock)

		_, err := useCase.Execute("username")

		assert.Error(t, err)
	})

	t.Run("Should return error when repository return empty birthday", func(t *testing.T) {
		setup(t)

		mirthdayRepositoryMock.EXPECT().GetBirthday("username").Return("", nil)

		useCase := NewGetBirthDayUseCase(mirthdayRepositoryMock)

		response, err := useCase.Execute("username")

		assert.NoError(t, err)
		assert.Equal(t, "username not found", response.ErrorMsg)
		assert.Equal(t, 404, response.ErrorCode)
	})

	t.Run("Should return error when repository return invalid birthday", func(t *testing.T) {
		setup(t)

		mirthdayRepositoryMock.EXPECT().GetBirthday("username").Return("invalid", nil)

		useCase := NewGetBirthDayUseCase(mirthdayRepositoryMock)

		_, err := useCase.Execute("username")

		assert.Error(t, err)
	})

	t.Run("Should return error when repository return birthday is in the future", func(t *testing.T) {
		setup(t)

		tomorrowStr := time.Now().AddDate(1, 0, 1).Format("2006-01-02")

		mirthdayRepositoryMock.EXPECT().GetBirthday("username").Return(tomorrowStr, nil)

		useCase := NewGetBirthDayUseCase(mirthdayRepositoryMock)

		_, err := useCase.Execute("username")

		assert.Error(t, err)

	})

	t.Run("Should return message when repository return valid birthday", func(t *testing.T) {
		setup(t)

		tomorrowStr := time.Now().AddDate(-1, 0, 1).Format("2006-01-02")

		mirthdayRepositoryMock.EXPECT().GetBirthday("username").Return(tomorrowStr, nil)

		useCase := NewGetBirthDayUseCase(mirthdayRepositoryMock)

		response, err := useCase.Execute("username")

		assert.NoError(t, err)
		assert.NotEmpty(t, response.Message)
	})

	t.Run("Should return message when repository returns today birthday", func(t *testing.T) {
		setup(t)

		todayStr := time.Now().AddDate(-1, 0, 0).Format("2006-01-02")

		mirthdayRepositoryMock.EXPECT().GetBirthday("username").Return(todayStr, nil)

		useCase := NewGetBirthDayUseCase(mirthdayRepositoryMock)

		response, err := useCase.Execute("username")

		assert.NoError(t, err)
		assert.Equal(t, "Hello, username! Happy birthday! V2", response.Message)
	})
}
