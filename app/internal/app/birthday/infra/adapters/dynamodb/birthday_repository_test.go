package dynamodb

import (
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/stretchr/testify/assert"
	"github.com/victoraldir/birthday-api/app/internal/app/birthday/domain"

	dynamodb_mock "github.com/victoraldir/birthday-api/app/internal/app/birthday/infra/adapters/dynamodb/mocks"
	"go.uber.org/mock/gomock"
)

var dynamodDBClientMock *dynamodb_mock.MockDynamodDBClient

func setup(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dynamodDBClientMock = dynamodb_mock.NewMockDynamodDBClient(ctrl)

}

func TestBirthdayRepo_PutBirthday(t *testing.T) {
	t.Run("Should return nil if birthday is successfully saved", func(t *testing.T) {
		setup(t)

		dynamodDBClientMock.EXPECT().PutItem(gomock.Any()).Return(nil, nil)

		birthdayRepo := NewBirthdayRepo(dynamodDBClientMock, "test")

		err := birthdayRepo.PutBirthday(domain.Birthday{
			Username:    "test",
			DateOfBirth: "1990-01-01",
		})

		assert.Nil(t, err)
	})

	t.Run("Should return error if birthday is not successfully saved", func(t *testing.T) {
		setup(t)

		expectedError := fmt.Errorf("error")

		dynamodDBClientMock.EXPECT().PutItem(gomock.Any()).Return(nil, expectedError)

		birthdayRepo := NewBirthdayRepo(dynamodDBClientMock, "test")

		err := birthdayRepo.PutBirthday(domain.Birthday{
			Username:    "test",
			DateOfBirth: "1990-01-01",
		})

		assert.NotNil(t, err)
	})

}

func TestBirthdayRepo_GetBirthday(t *testing.T) {
	t.Run("Should return birthday if it exists", func(t *testing.T) {
		setup(t)

		expectedUsername := "username"
		expectedBirthday := "1990-01-01"

		dynamodDBClientMock.EXPECT().GetItem(gomock.Any()).Return(&dynamodb.GetItemOutput{
			Item: map[string]*dynamodb.AttributeValue{
				"username": {
					S: &expectedUsername,
				},
				"dateOfBirth": {
					S: &expectedBirthday,
				},
			},
		}, nil)

		birthdayRepo := NewBirthdayRepo(dynamodDBClientMock, "test")

		birthday, err := birthdayRepo.GetBirthday("test")

		assert.Nil(t, err)
		assert.Equal(t, expectedBirthday, birthday)
	})

	t.Run("Should return empty string if birthday does not exist", func(t *testing.T) {
		setup(t)

		dynamodDBClientMock.EXPECT().GetItem(gomock.Any()).Return(&dynamodb.GetItemOutput{}, nil)

		birthdayRepo := NewBirthdayRepo(dynamodDBClientMock, "test")

		result, err := birthdayRepo.GetBirthday("test")

		assert.Nil(t, err)
		assert.Equal(t, "", result)
	})

	t.Run("Should return error if dynamodb returns error", func(t *testing.T) {
		setup(t)

		expectedError := fmt.Errorf("error")

		dynamodDBClientMock.EXPECT().GetItem(gomock.Any()).Return(nil, expectedError)

		birthdayRepo := NewBirthdayRepo(dynamodDBClientMock, "test")

		_, err := birthdayRepo.GetBirthday("test")

		assert.NotNil(t, err)
	})
}
