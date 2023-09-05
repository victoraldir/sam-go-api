package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBirthday_IsAlpha(t *testing.T) {
	t.Run("Should return true if username contains only letters", func(t *testing.T) {

		// Arrange
		birthday := Birthday{
			Username: "John",
		}

		// Act
		result := birthday.IsAlpha()

		// Assert
		assert.True(t, result)
	})

	t.Run("Should return false if username contains numbers", func(t *testing.T) {

		// Arrange
		birthday := Birthday{
			Username: "John123",
		}

		// Act
		result := birthday.IsAlpha()

		// Assert
		assert.False(t, result)
	})

	t.Run("Should return false if username is empty", func(t *testing.T) {

		// Arrange
		birthday := Birthday{
			Username: "",
		}

		// Act
		result := birthday.IsAlpha()

		// Assert
		assert.False(t, result)
	})
}
