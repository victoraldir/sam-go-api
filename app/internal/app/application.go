package app

import (
	"log"
	"log/slog"
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
var logLevel string

const (
	awsApiKey          = "dummykey"
	awsSecretAccessKey = "dummysecret"
	region             = "us-east-1"
	localDynamodbAddr  = "http://dynamodb:8000"
)

func NewAPIGatewayV2Handler() handlers.APIGatewayV2Handler {

	// Load environment variables
	loadEnv()

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

	if env == "local" {
		return createLocalDynamodbClient()
	}

	log.Println("Creating AWS DynamoDB client")
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
		Config: aws.Config{
			Region: aws.String(region)},
	}))

	return dynamodb.New(sess)
}

func createLocalDynamodbClient() *dynamodb.DynamoDB {

	slog.Info("Creating local DynamoDB client")

	// Set dummy credentials
	os.Setenv("AWS_ACCESS_KEY_ID", awsApiKey)
	os.Setenv("AWS_SECRET_ACCESS_KEY", awsSecretAccessKey)

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
		Config: aws.Config{
			Region:   aws.String(region),
			Endpoint: aws.String(localDynamodbAddr)},
	}))

	slog.Debug("Connecting to local DynamoDB on %s with AWS_ACCESS_KEY_ID: %s, AWS_SECRET_ACCESS_KEY: %s, and REGIO: %s",
		localDynamodbAddr,
		awsApiKey,
		awsSecretAccessKey,
		region)

	return dynamodb.New(sess)
}

// Load environment variables
func loadEnv() {

	slog.Info("Loading environment variables")
	table_name = os.Getenv("TABLE_NAME")

	if table_name == "" {
		panic("TABLE_NAME environment variable is required")
	}

	env = os.Getenv("ENV")

	if env == "" {
		panic("ENV environment variable is required")
	}

	logLevel = os.Getenv("LOG_LEVEL")

	if logLevel == "" {
		logLevel = "info"
	}

	slog.Debug("Environment variables loaded: TABLE_NAME: %s, ENV: %s", table_name, env)
}
