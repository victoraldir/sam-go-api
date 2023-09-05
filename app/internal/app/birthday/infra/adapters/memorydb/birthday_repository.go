package memorydb

import "github.com/victoraldir/birthday-api/app/internal/app/birthday/domain"

type BirthdayRepository struct {
	birthdays map[string]domain.Birthday
}

func NewBirthdayRepository() *BirthdayRepository {
	return &BirthdayRepository{
		birthdays: make(map[string]domain.Birthday, 100),
	}
}

func (repository *BirthdayRepository) PutBirthday(birthday domain.Birthday) error {
	repository.birthdays[birthday.Username] = birthday
	return nil
}

func (repository *BirthdayRepository) GetBirthday(username string) (string, error) {
	birthday, ok := repository.birthdays[username]
	if !ok {
		return "", nil
	}

	return birthday.DateOfBirth, nil
}
