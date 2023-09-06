package dynamodb

import (
	"github.com/victoraldir/birthday-api/app/internal/app/birthday/domain"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

//go:generate mockgen -destination=../dynamodb/mocks/mockDynamodDBClient.go -package=dynamodb github.com/victoraldir/birthday-api/app/internal/app/birthday/infra/adapters/dynamodb DynamodDBClient
type DynamodDBClient interface {
	PutItem(input *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error)
	GetItem(input *dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error)
}

type birthdayRepo struct {
	client    DynamodDBClient
	tableName string
}

// NewBirthdayRepository returns a new instance of a DynamoDB birthday repository.
func NewBirthdayRepo(client DynamodDBClient, tableName string) domain.BirthdayRepository {
	return birthdayRepo{
		client:    client,
		tableName: tableName,
	}
}

func (r birthdayRepo) PutBirthday(birthday domain.Birthday) error {

	bday, _ := dynamodbattribute.MarshalMap(birthday)

	input := &dynamodb.PutItemInput{
		Item:      bday,
		TableName: aws.String(r.tableName),
	}

	_, err := r.client.PutItem(input)
	if err != nil {
		return err
	}

	return nil
}

func (r birthdayRepo) GetBirthday(username string) (string, error) {
	input := &dynamodb.GetItemInput{
		TableName: aws.String(r.tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"username": {
				S: aws.String(username),
			},
		},
	}

	result, err := r.client.GetItem(input)
	if err != nil {
		return "", err
	}

	birthday := domain.Birthday{}
	dynamodbattribute.UnmarshalMap(result.Item, &birthday)

	return birthday.DateOfBirth, nil
}
