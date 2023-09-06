package usecases

import (
	errorcode "github.com/victoraldir/birthday-api/app/pkg/error_code"

	"github.com/victoraldir/birthday-api/app/internal/app/birthday/domain"
)

type PutBirthdayCommand struct {
	Username    string `json:"username"`
	DateOfBirth string `json:"dateOfBirth"`
}

type PutBirthdayResponse struct {
	Message   string              `json:"message"`
	ErrorCode int                 `json:"errorCode"`
	ErrorType errorcode.ErrorType `json:"errorType"`
	ErrorMsg  string              `json:"errorMsg"`
}

//go:generate mockgen -destination=../usecases/mocks/mockPutBirthdayUseCase.go -package=usecases github.com/victoraldir/birthday-api/app/internal/app/birthday/usecases PutBirthdayUseCase
type PutBirthdayUseCase interface {
	Execute(command PutBirthdayCommand) (*PutBirthdayResponse, error)
}

type putBirthDayUseCase struct {
	repository domain.BirthdayRepository
}

func NewPutBirthDayUseCase(repository domain.BirthdayRepository) PutBirthdayUseCase {
	return &putBirthDayUseCase{
		repository: repository,
	}
}

func (useCase *putBirthDayUseCase) Execute(command PutBirthdayCommand) (*PutBirthdayResponse, error) {

	commandDomain := domain.Birthday{
		Username:    command.Username,
		DateOfBirth: command.DateOfBirth,
	}

	// Check if Username contains only letters
	if !commandDomain.IsAlpha() {
		return &PutBirthdayResponse{
			ErrorCode: 400,
			ErrorType: errorcode.InvalidUsername,
			ErrorMsg:  "username must contain only letters",
		}, nil
	}

	// Check if date is before today
	if !commandDomain.IsBeforeToday() {
		return &PutBirthdayResponse{
			ErrorCode: 400,
			ErrorType: errorcode.InvalidDateOfBirth,
			ErrorMsg:  "date of birth must be before today",
		}, nil
	}

	err := useCase.repository.PutBirthday(commandDomain)
	if err != nil {
		return nil, err
	}

	return &PutBirthdayResponse{
		Message: "Birthday saved!",
	}, nil
}
