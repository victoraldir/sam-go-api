package dynamodb

import (
	"log"

	"github.com/victoraldir/birthday-api/app/internal/app/birthday/domain"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type birthdayRepo struct {
	client    *dynamodb.DynamoDB
	tableName string
}

// NewBirthdayRepository returns a new instance of a DynamoDB birthday repository.
func NewBirthdayRepo(client *dynamodb.DynamoDB, tableName string) domain.BirthdayRepository {
	return birthdayRepo{
		client:    client,
		tableName: tableName,
	}
}

func (r birthdayRepo) PutBirthday(birthday domain.Birthday) error {

	bday, err := dynamodbattribute.MarshalMap(birthday)
	if err != nil {
		log.Fatalf("Got error marshalling new movie item: %s", err)
	}

	input := &dynamodb.PutItemInput{
		Item:      bday,
		TableName: aws.String(r.tableName),
	}

	_, err = r.client.PutItem(input)
	if err != nil {
		log.Fatalf("Got error calling PutItem: %s", err)
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
		log.Fatalf("Got error calling GetItem: %s", err)
	}

	birthday := domain.Birthday{}
	err = dynamodbattribute.UnmarshalMap(result.Item, &birthday)
	if err != nil {
		log.Fatalf("Got error unmarshalling: %s", err)
	}

	return birthday.DateOfBirth, nil
}
