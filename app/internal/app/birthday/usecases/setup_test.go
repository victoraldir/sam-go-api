package usecases

import (
	"testing"

	domain "github.com/victoraldir/birthday-api/app/internal/app/birthday/domain/mocks"
	"go.uber.org/mock/gomock"
)

var mirthdayRepositoryMock *domain.MockBirthdayRepository

func setup(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mirthdayRepositoryMock = domain.NewMockBirthdayRepository(ctrl)
}
