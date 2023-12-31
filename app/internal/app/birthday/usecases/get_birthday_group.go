package usecases

import (
	"fmt"
	"time"

	"github.com/victoraldir/birthday-api/app/internal/app/birthday/domain"
	"github.com/victoraldir/birthday-api/app/pkg/datetime"
	errorcode "github.com/victoraldir/birthday-api/app/pkg/error_code"
	"golang.org/x/exp/slog"
)

//go:generate mockgen -destination=../usecases/mocks/mockGetBirthdayUseCase.go -package=usecases github.com/victoraldir/birthday-api/app/internal/app/birthday/usecases GetBirthdayUseCase
type GetBirthdayUseCase interface {
	Execute(username string) (*GetBirthdayResponse, error)
}

type GetBirthdayResponse struct {
	Message   string              `json:"message"`
	ErrorCode int                 `json:"errorCode"`
	ErrorType errorcode.ErrorType `json:"errorType"`
	ErrorMsg  string              `json:"errorMsg"`
}

type getBirthDayUseCase struct {
	repository domain.BirthdayRepository
}

func NewGetBirthDayUseCase(repository domain.BirthdayRepository) GetBirthdayUseCase {
	return &getBirthDayUseCase{
		repository: repository,
	}
}

func (useCase *getBirthDayUseCase) Execute(username string) (*GetBirthdayResponse, error) {

	slog.Debug("Executing GetBirthdayUseCase with username: %s", username)

	birthday, err := useCase.repository.GetBirthday(username)
	if err != nil {
		slog.Error("Error executing repository GetBirthday: %s", err)
		return nil, err
	}

	if birthday == "" {
		return &GetBirthdayResponse{
			ErrorCode: 404,
			ErrorType: errorcode.NotFound,
			ErrorMsg:  "username not found",
		}, nil
	}

	date, err := time.Parse("2006-01-02", birthday)

	if err != nil {
		return nil, err
	}

	var msg string

	//Check if birthday is today
	if datetime.IsToday(date) {
		msg = fmt.Sprintf("Hello, %s! Happy birthday! V2", username)
		return &GetBirthdayResponse{
			Message: msg,
		}, nil
	}

	// Check how many days until birthday
	daysUntil, err := datetime.DaysUntil(date)
	if err != nil {
		return nil, err
	}

	return &GetBirthdayResponse{
		Message: fmt.Sprintf("Hello, %s! Your birthday is in %d day(s).", username, daysUntil),
	}, nil
}
