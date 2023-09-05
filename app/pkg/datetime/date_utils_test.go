package datetime

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestIsToday(t *testing.T) {
	t.Run("Should return true if date is today", func(t *testing.T) {
		// Arrange
		date, _ := ParseDate("1990-02-17")

		// Act
		result := IsToday(date)

		// Assert
		assert.False(t, result, "Date is not today")
	})

	t.Run("Should return false if date is not today", func(t *testing.T) {
		// Arrange
		date, _ := ParseDate("1990-02-17")

		// Act
		result := IsToday(date)

		// Assert
		assert.False(t, result, "Date is not today")
	})

	t.Run("Should return false if date is empty", func(t *testing.T) {
		// Arrange
		date, _ := ParseDate("")

		// Act
		result := IsToday(date)

		// Assert
		assert.False(t, result, "Date is not today")
	})
}

func TestDaysUntil(t *testing.T) {
	t.Run("Should return 0 if date is today", func(t *testing.T) {

		// Arrange
		dateStr := time.Now().Format("2006-01-02")
		date, _ := ParseDate(dateStr)
		date = date.AddDate(1990, 0, 0)

		// Act
		result, _ := DaysUntil(date)

		// Assert
		assert.EqualValues(t, 0, result, "Date is today")
	})

	t.Run("Should throw error if date is tomorrow", func(t *testing.T) {

		// Arrange
		dateStr := time.Now().AddDate(0, 0, 1).Format("2006-01-02")
		date, _ := ParseDate(dateStr)

		// Act
		_, err := DaysUntil(date)

		// Assert
		assert.NotNil(t, err, "Date is tomorrow")
	})

	t.Run("Should return 365 if date is yesterday", func(t *testing.T) {

		// Arrange
		dateStr := time.Now().AddDate(0, 0, -1).Format("2006-01-02")
		date, _ := ParseDate(dateStr)

		// Act
		result, _ := DaysUntil(date)

		// Assert
		assert.EqualValues(t, 365, result, "Date is yesterday")
	})

	t.Run("Should throw error if date is in the future", func(t *testing.T) {

		// Arrange
		dateStr := time.Now().AddDate(0, 0, 1).Format("2006-01-02")
		date, _ := ParseDate(dateStr)

		// Act
		_, err := DaysUntil(date)

		// Assert
		assert.NotNil(t, err, "Date is in the future")
	})
}

func TestParseDate(t *testing.T) {
	t.Run("Should return date if date is valid", func(t *testing.T) {

		// Arrange
		dateStr := "1990-02-17"

		// Act
		result, _ := ParseDate(dateStr)

		// Assert
		assert.NotNil(t, result, "Date is valid")
	})

	t.Run("Should throw error if date is invalid", func(t *testing.T) {

		// Arrange
		dateStrDataset := []string{
			"01-02-17",
			"1990-AB-17",
			"01-02-1990",
		}

		// Act
		for _, dateStr := range dateStrDataset {
			_, err := ParseDate(dateStr)

			// Assert
			assert.NotNil(t, err, "Date is invalid")
		}

	})
}

func TestIsBeforeToday(t *testing.T) {
	t.Run("Should return true if date is before today", func(t *testing.T) {

		// Arrange
		dateStr := time.Now().AddDate(0, 0, -1).Format("2006-01-02")

		// Act
		result := IsBeforeToday(dateStr)

		// Assert
		assert.True(t, result, "Date is before today")
	})

	t.Run("Should return false if date is today", func(t *testing.T) {

		// Arrange
		dateStr := time.Now().Format("2006-01-02")

		// Act
		result := IsBeforeToday(dateStr)

		// Assert
		assert.False(t, result, "Date is today")
	})

	t.Run("Should return false if date is after today", func(t *testing.T) {

		// Arrange
		dateStr := time.Now().AddDate(0, 0, 1).Format("2006-01-02")

		// Act
		result := IsBeforeToday(dateStr)

		// Assert
		assert.False(t, result, "Date is after today")
	})

	t.Run("Should return false if date is invalid", func(t *testing.T) {

		// Arrange
		dateStr := "1990-AB-17"

		// Act
		result := IsBeforeToday(dateStr)

		// Assert
		assert.False(t, result, "Date is invalid")
	})
}
