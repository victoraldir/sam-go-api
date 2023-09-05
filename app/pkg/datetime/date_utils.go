package datetime

import (
	"fmt"
	"time"
)

// Create a function that parses a string date YYYY-MM-DD to a date
func ParseDate(date string) (time.Time, error) {

	// Parse string date to date
	dateTime, err := time.Parse("2006-01-02", date)

	if err != nil {
		return time.Time{}, err
	}

	return dateTime, nil
}

// Create a function that checks if the string date YYYY-MM-DD is today
func IsToday(date time.Time) bool {

	if date.Day() == time.Now().Day() && date.Month() == time.Now().Month() {
		return true
	}

	return false
}

// Create a function that checks how many days until the birthday YYYY-MM-DD
func DaysUntil(date time.Time) (int, error) {
	currentDate := time.Now()

	// Check if date is in the future
	if date.After(currentDate) {
		return 0, fmt.Errorf("date %s is in the future", date)
	}

	// Check if birthday is today
	if date.Day() == currentDate.Day() && date.Month() == currentDate.Month() {
		return 0, nil
	}

	// Create a new date with the same year as the current date
	birthdayThisYear := time.Date(currentDate.Year(), date.Month(), date.Day(), 0, 0, 0, 0, time.UTC)

	// Check if birthday is in the past this year
	if birthdayThisYear.Before(currentDate) {
		// Add one year to birthday
		birthdayThisYear = birthdayThisYear.AddDate(1, 0, 0)
	}

	// Calculate days until birthday
	daysUntil := birthdayThisYear.Sub(currentDate).Hours() / 24

	// Round up the number of days to the nearest integer
	if daysUntil < 0 {
		return int(daysUntil) - 1, nil
	}
	return int(daysUntil) + 1, nil
}

// Create a function that checks if the string date YYYY-MM-DD is before today
func IsBeforeToday(dateStr string) bool {

	// Parse string date to date
	dateTime, err := ParseDate(dateStr)

	if err != nil {
		return false
	}

	// Check if date is today
	if IsToday(dateTime) {
		return false
	}

	return dateTime.Before(time.Now())
}
