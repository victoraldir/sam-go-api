package handlers

import (
	"encoding/json"

	"github.com/victoraldir/birthday-api/app/internal/app/birthday/usecases"
	"golang.org/x/exp/slog"

	"github.com/aws/aws-lambda-go/events"
)

type APIGatewayV2Handler struct {
	PutBirthdayUseCase usecases.PutBirthdayUseCase
	GetBirthdayUseCase usecases.GetBirthdayUseCase
}

type GetBirthdayResponse struct {
	Message string `json:"message"`
}

func NewAPIGatewayV2Handler(putBirthdayUseCase usecases.PutBirthdayUseCase, getBirthdayUseCase usecases.GetBirthdayUseCase) APIGatewayV2Handler {
	return APIGatewayV2Handler{
		PutBirthdayUseCase: putBirthdayUseCase,
		GetBirthdayUseCase: getBirthdayUseCase,
	}
}

func (h *APIGatewayV2Handler) PutBirthdayHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	slog.Debug("Received request %s", request)

	userName := request.PathParameters["username"]
	if userName == "" {
		slog.Warn("username is required")
		return events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "username is required",
		}, nil
	}

	var putBirthdayCommand usecases.PutBirthdayCommand
	json.Unmarshal([]byte(request.Body), &putBirthdayCommand)

	if putBirthdayCommand.DateOfBirth == "" {
		slog.Warn("dateOfBirth is required")
		return events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "dateOfBirth is required",
		}, nil
	}

	putBirthdayCommand.Username = userName

	response, err := h.PutBirthdayUseCase.Execute(putBirthdayCommand)
	if err != nil {
		slog.Error("Error executing PutBirthdayUseCase: %s", err)
		return events.APIGatewayProxyResponse{}, err
	}

	if response.ErrorType != "" {
		slog.Warn("Error executing PutBirthdayUseCase: %s", response.ErrorMsg)
		return events.APIGatewayProxyResponse{
			Body:       response.ErrorMsg,
			StatusCode: response.ErrorCode,
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 204,
	}, nil
}

func (h *APIGatewayV2Handler) GetBirthdayHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	slog.Debug("Received request %s", request)

	userName := request.PathParameters["username"]

	if userName == "" {
		slog.Warn("username is required")
		return events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "username is required",
		}, nil
	}

	getBirthdayResponseUc, err := h.GetBirthdayUseCase.Execute(userName)
	if err != nil {
		slog.Error("Error executing GetBirthdayUseCase: %s", err)
		return events.APIGatewayProxyResponse{}, err
	}

	if getBirthdayResponseUc.ErrorType != "" {
		slog.Warn("Error executing GetBirthdayUseCase: %s", getBirthdayResponseUc.ErrorMsg)
		return events.APIGatewayProxyResponse{
			Body:       getBirthdayResponseUc.ErrorMsg,
			StatusCode: getBirthdayResponseUc.ErrorCode,
		}, nil
	}

	getBirthdayResponse := GetBirthdayResponse{
		Message: getBirthdayResponseUc.Message,
	}

	body, _ := json.Marshal(getBirthdayResponse)

	return events.APIGatewayProxyResponse{
		Body:       string(body),
		Headers:    map[string]string{"Content-Type": "application/json"},
		StatusCode: 200,
	}, nil
}
