package domain

import "github.com/victoraldir/birthday-api/app/pkg/datetime"

type Birthday struct {
	Username    string `json:"username"`
	DateOfBirth string `json:"dateOfBirth"`
}

// Check if username contains only letters
func (b *Birthday) IsAlpha() bool {

	if len(b.Username) == 0 {
		return false
	}

	for _, char := range b.Username {
		if char < 'A' || char > 'z' {
			return false
		}
	}
	return true
}

// Check if is a date before the today date.
func (b *Birthday) IsBeforeToday() bool {
	return datetime.IsBeforeToday(b.DateOfBirth)
}
