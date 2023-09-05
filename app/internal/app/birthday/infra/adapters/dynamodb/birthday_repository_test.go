package dynamodb

// import (
// 	"testing"

// 	"github.com/victoraldir/birthday-api/app/internal/app/birthday/domain"

// 	"github.com/aws/aws-sdk-go/aws"
// 	"github.com/aws/aws-sdk-go/aws/session"
// 	"github.com/aws/aws-sdk-go/service/dynamodb"
// 	"github.com/stretchr/testify/assert"
// )

// func TestBirthdayRepo_PutBirthday(t *testing.T) {
// 	t.Run("Should save birthday", func(t *testing.T) {
// 		// Arrange
// 		sess := session.Must(session.NewSessionWithOptions(session.Options{
// 			SharedConfigState: session.SharedConfigEnable,
// 			Config: aws.Config{
// 				Region: aws.String("us-east-1")},
// 		}))

// 		svc := dynamodb.New(sess)
// 		repo := NewBirthdayRepo(svc, "birthday")

// 		birthday := domain.Birthday{
// 			Username:    "test",
// 			DateOfBirth: "2020-01-01",
// 		}

// 		// Act
// 		err := repo.PutBirthday(birthday)

// 		// Assert
// 		assert.Nil(t, err)

// 	})
// }

// func TestBirthdayRepo_GetBirthday(t *testing.T) {
// 	t.Run("Should get birthday", func(t *testing.T) {
// 		// Arrange
// 		sess := session.Must(session.NewSessionWithOptions(session.Options{
// 			SharedConfigState: session.SharedConfigEnable,
// 			Config: aws.Config{
// 				Region: aws.String("us-east-1")},
// 		}))

// 		svc := dynamodb.New(sess)
// 		repo := NewBirthdayRepo(svc, "birthday")

// 		birthday := domain.Birthday{
// 			Username:    "test",
// 			DateOfBirth: "2020-01-01",
// 		}

// 		err := repo.PutBirthday(birthday)
// 		assert.Nil(t, err)

// 		// Act
// 		bday, err := repo.GetBirthday("test")

// 		// Assert
// 		assert.Nil(t, err)
// 		assert.Equal(t, "2020-01-01", bday)
// 	})
// }
