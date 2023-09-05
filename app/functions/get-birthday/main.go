package main

import (
	"github.com/victoraldir/birthday-api/app/internal/app"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	handler := app.NewAPIGatewayV2Handler()
	lambda.Start(handler.GetBirthdayHandler)
}
