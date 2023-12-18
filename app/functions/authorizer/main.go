package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(handleRequest)
}

// Help function to generate an IAM policy
func generatePolicy(principalId, effect, resource string) events.APIGatewayCustomAuthorizerResponse {
	authResponse := events.APIGatewayCustomAuthorizerResponse{PrincipalID: principalId}

	if effect != "" && resource != "" {

		resourceFirstPart := fmt.Sprintf("%s/*", strings.Split(resource, "/")[0])

		authResponse.PolicyDocument = events.APIGatewayCustomAuthorizerPolicy{
			Version: "2012-10-17",
			Statement: []events.IAMPolicyStatement{
				{
					Action:   []string{"execute-api:Invoke"},
					Effect:   effect,
					Resource: []string{resourceFirstPart},
				},
			},
		}
	}

	// Optional output with custom properties of the String, Number or Boolean type.
	authResponse.Context = map[string]interface{}{
		"stringKey":  "stringval",
		"numberKey":  123,
		"booleanKey": true,
	}
	return authResponse
}

// Implement Lambda Authorizer handler
func handleRequest(ctx context.Context, event events.APIGatewayCustomAuthorizerRequest) (events.APIGatewayCustomAuthorizerResponse, error) {

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	slog.Info("event.AuthorizationToken", "token", event.AuthorizationToken)

	token := event.AuthorizationToken
	switch strings.ToLower(token) {
	case "allow":

		generatedPolicy := generatePolicy("user", "Allow", event.MethodArn)

		slog.Info("generatedPolicy", "generatedPolicy", generatedPolicy)

		return generatedPolicy, nil
	case "deny":

		generatedPolicy := generatePolicy("user", "Deny", event.MethodArn)

		slog.Info("generatedPolicy", "generatedPolicy", generatedPolicy)

		return generatedPolicy, nil
	case "unauthorized":
		return events.APIGatewayCustomAuthorizerResponse{}, errors.New("Unauthorized") // Return a 401 Unauthorized response
	default:
		return events.APIGatewayCustomAuthorizerResponse{}, errors.New("error: Invalid token")
	}
}
