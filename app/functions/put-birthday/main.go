package main

import (
	"log/slog"

	"github.com/victoraldir/birthday-api/app/internal/app"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	slog.Info("Starting get-birthday function")
	handler := app.NewAPIGatewayV2Handler()
	lambda.Start(handler.PutBirthdayHandler)
}
