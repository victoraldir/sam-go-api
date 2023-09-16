package handlers

import (
	"fmt"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
	"github.com/victoraldir/birthday-api/app/internal/app/birthday/usecases"
	usecases_mock "github.com/victoraldir/birthday-api/app/internal/app/birthday/usecases/mocks"
	"go.uber.org/mock/gomock"
)

var putBirthdayUseCaseMock *usecases_mock.MockPutBirthdayUseCase
var getBirthdayUseCaseMock *usecases_mock.MockGetBirthdayUseCase

func setup(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	putBirthdayUseCaseMock = usecases_mock.NewMockPutBirthdayUseCase(ctrl)
	getBirthdayUseCaseMock = usecases_mock.NewMockGetBirthdayUseCase(ctrl)
}

func TestAPIGatewayV2Handler_PutBirthdayHandler(t *testing.T) {

	t.Run("Should return 400 if username is empty", func(t *testing.T) {

		setup(t)

		// Arrange
		apiGatewayV2Handler := NewAPIGatewayV2Handler(putBirthdayUseCaseMock, getBirthdayUseCaseMock)

		// Act
		response, err := apiGatewayV2Handler.PutBirthdayHandler(events.APIGatewayProxyRequest{
			PathParameters: map[string]string{
				"username": "",
			},
		})

		// Assert
		assert.Nil(t, err)
		assert.Equal(t, 400, response.StatusCode)
	})

	t.Run("Should return 400 if dateOfBirth is empty", func(t *testing.T) {

		setup(t)

		// Arrange
		apiGatewayV2Handler := NewAPIGatewayV2Handler(putBirthdayUseCaseMock, getBirthdayUseCaseMock)

		// Act
		response, err := apiGatewayV2Handler.PutBirthdayHandler(events.APIGatewayProxyRequest{
			PathParameters: map[string]string{
				"username": "username",
			},
			Body: `{"username": "username"}`,
		})

		// Assert
		assert.Nil(t, err)
		assert.Equal(t, 400, response.StatusCode)
	})

	t.Run("Should return 400 if username contains numbers", func(t *testing.T) {

		setup(t)

		// Arrange
		putBirthdayUseCaseMock.EXPECT().Execute(gomock.Any()).Return(&usecases.PutBirthdayResponse{
			ErrorCode: 400,
			ErrorType: "InvalidUsername",
			ErrorMsg:  "username must contain only letters",
		}, nil)
		apiGatewayV2Handler := NewAPIGatewayV2Handler(putBirthdayUseCaseMock, getBirthdayUseCaseMock)

		// Act
		response, err := apiGatewayV2Handler.PutBirthdayHandler(events.APIGatewayProxyRequest{
			PathParameters: map[string]string{
				"username": "username1",
			},
			Body: `{"username": "username1", "dateOfBirth": "2020-01-01"}`,
		})

		// Assert
		assert.Nil(t, err)
		assert.Equal(t, 400, response.StatusCode)
	})

	t.Run("Should return 400 if dateOfBirth is after today", func(t *testing.T) {

		setup(t)

		// Arrange
		putBirthdayUseCaseMock.EXPECT().Execute(gomock.Any()).Return(&usecases.PutBirthdayResponse{
			ErrorCode: 400,
			ErrorType: "InvalidDateOfBirth",
			ErrorMsg:  "date of birth must be before today",
		}, nil)
		apiGatewayV2Handler := NewAPIGatewayV2Handler(putBirthdayUseCaseMock, getBirthdayUseCaseMock)

		// Act
		response, err := apiGatewayV2Handler.PutBirthdayHandler(events.APIGatewayProxyRequest{
			PathParameters: map[string]string{
				"username": "username",
			},
			Body: `{"username": "username", "dateOfBirth": "2021-01-01"}`,
		})

		// Assert
		assert.Nil(t, err)
		assert.Equal(t, 400, response.StatusCode)
	})

	t.Run("Should return 204 if everything is ok", func(t *testing.T) {

		setup(t)

		// Arrange
		putBirthdayUseCaseMock.EXPECT().Execute(gomock.Any()).Return(&usecases.PutBirthdayResponse{
			Message: "Hello, username! Happy birthday!",
		}, nil)
		apiGatewayV2Handler := NewAPIGatewayV2Handler(putBirthdayUseCaseMock, getBirthdayUseCaseMock)

		// Act
		response, err := apiGatewayV2Handler.PutBirthdayHandler(events.APIGatewayProxyRequest{
			PathParameters: map[string]string{
				"username": "username",
			},
			Body: `{"username": "username", "dateOfBirth": "2020-01-01"}`,
		})

		// Assert
		assert.Nil(t, err)
		assert.Equal(t, 204, response.StatusCode)
	})

	t.Run("Should return 500 if something goes wrong", func(t *testing.T) {

		setup(t)

		// Arrange

		expectedError := fmt.Errorf("some error")

		putBirthdayUseCaseMock.EXPECT().Execute(gomock.Any()).Return(nil, expectedError)
		apiGatewayV2Handler := NewAPIGatewayV2Handler(putBirthdayUseCaseMock, getBirthdayUseCaseMock)

		// Act
		response, err := apiGatewayV2Handler.PutBirthdayHandler(events.APIGatewayProxyRequest{
			PathParameters: map[string]string{
				"username": "username",
			},
			Body: `{"username": "username", "dateOfBirth": "2020-01-01"}`,
		})

		// Assert
		assert.NotNil(t, err)
		assert.NotNil(t, response)
	})
}

func TestAPIGatewayV2Handler_GetBirthdayHandler(t *testing.T) {

	t.Run("Should return 400 if username is empty", func(t *testing.T) {

		setup(t)

		// Arrange
		getBirthdayUseCaseMock.EXPECT().Execute(gomock.Any()).Return(nil, nil)
		apiGatewayV2Handler := NewAPIGatewayV2Handler(putBirthdayUseCaseMock, getBirthdayUseCaseMock)

		// Act
		response, err := apiGatewayV2Handler.GetBirthdayHandler(events.APIGatewayProxyRequest{
			PathParameters: map[string]string{
				"username": "",
			},
		})

		// Assert
		assert.Nil(t, err)
		assert.Equal(t, 400, response.StatusCode)
	})

	t.Run("Should return 404 if username is not found", func(t *testing.T) {

		setup(t)

		// Arrange
		getBirthdayUseCaseMock.EXPECT().Execute(gomock.Any()).Return(&usecases.GetBirthdayResponse{
			ErrorCode: 404,
			ErrorType: "NotFound",
			ErrorMsg:  "username not found",
		}, nil)
		apiGatewayV2Handler := NewAPIGatewayV2Handler(putBirthdayUseCaseMock, getBirthdayUseCaseMock)

		// Act
		response, err := apiGatewayV2Handler.GetBirthdayHandler(events.APIGatewayProxyRequest{
			PathParameters: map[string]string{
				"username": "username",
			},
		})

		// Assert
		assert.Nil(t, err)
		assert.Equal(t, 404, response.StatusCode)
	})

	t.Run("Should return 200 if everything is ok", func(t *testing.T) {

		setup(t)

		// Arrange
		getBirthdayUseCaseMock.EXPECT().Execute(gomock.Any()).Return(&usecases.GetBirthdayResponse{
			Message: "Hello, username! Your birthday is in 1 day(s)",
		}, nil)
		apiGatewayV2Handler := NewAPIGatewayV2Handler(putBirthdayUseCaseMock, getBirthdayUseCaseMock)

		// Act
		response, err := apiGatewayV2Handler.GetBirthdayHandler(events.APIGatewayProxyRequest{
			PathParameters: map[string]string{
				"username": "username",
			},
		})

		// Assert
		assert.Nil(t, err)
		assert.Equal(t, 200, response.StatusCode)
	})

	t.Run("Should return 500 if something goes wrong", func(t *testing.T) {

		setup(t)

		// Arrange

		expectedError := fmt.Errorf("some error")

		getBirthdayUseCaseMock.EXPECT().Execute(gomock.Any()).Return(nil, expectedError)
		apiGatewayV2Handler := NewAPIGatewayV2Handler(putBirthdayUseCaseMock, getBirthdayUseCaseMock)

		// Act
		response, err := apiGatewayV2Handler.GetBirthdayHandler(events.APIGatewayProxyRequest{
			PathParameters: map[string]string{
				"username": "username",
			},
		})

		// Assert
		assert.NotNil(t, err)
		assert.NotNil(t, response)
	})

}
