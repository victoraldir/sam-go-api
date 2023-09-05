package app

import (
	"log"
	"os"

	"github.com/victoraldir/birthday-api/app/internal/app/infra/handlers"

	"github.com/victoraldir/birthday-api/app/internal/app/birthday/usecases"

	adapter_dynamodb "github.com/victoraldir/birthday-api/app/internal/app/birthday/infra/adapters/dynamodb"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var table_name string
var env string

func NewAPIGatewayV2Handler() handlers.APIGatewayV2Handler {

	// DynamoDB Service
	svc := createDynamodbClient()

	// DynamoDB Repository
	repo := adapter_dynamodb.NewBirthdayRepo(svc, "birthday")

	// Use Cases
	putBirthdayUseCase := usecases.NewPutBirthDayUseCase(repo)
	getBirthdayUseCase := usecases.NewGetBirthDayUseCase(repo)

	return handlers.APIGatewayV2Handler{
		PutBirthdayUseCase: putBirthdayUseCase,
		GetBirthdayUseCase: getBirthdayUseCase,
	}
}

func createDynamodbClient() *dynamodb.DynamoDB {

	loadEnv()

	if env == "local" {
		return createLocalDynamodbClient()
	}

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
		Config: aws.Config{
			Region: aws.String("us-east-1")},
	}))

	return dynamodb.New(sess)
}

func createLocalDynamodbClient() *dynamodb.DynamoDB {

	log.Println("Creating local DynamoDB client")

	// Set dummy credentials
	os.Setenv("AWS_ACCESS_KEY_ID", "dummykey")
	os.Setenv("AWS_SECRET_ACCESS_KEY", "dummysecret")

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
		Config: aws.Config{
			Region:   aws.String("us-east-1"),
			Endpoint: aws.String("http://dynamodb:8000")},
	}))

	return dynamodb.New(sess)
}

// Load environment variables
func loadEnv() {

	table_name = os.Getenv("TABLE_NAME")

	if table_name == "" {
		panic("TABLE_NAME environment variable is required")
	}

	env = os.Getenv("ENV")

	if env == "" {
		panic("ENV environment variable is required")
	}

}
