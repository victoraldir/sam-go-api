package domain

//go:generate mockgen -destination=../domain/mocks/mockBirthdayRepository.go -package=domain github.com/victoraldir/birthday-api/app/internal/app/birthday/domain BirthdayRepository
type BirthdayRepository interface {
	PutBirthday(birthday Birthday) error
	GetBirthday(username string) (string, error)
}
