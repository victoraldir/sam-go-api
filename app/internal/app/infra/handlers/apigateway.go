package handlers

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/victoraldir/birthday-api/app/internal/app/birthday/usecases"

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

func (h *APIGatewayV2Handler) HelloHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var greeting string
	sourceIP := request.RequestContext.Identity.SourceIP

	if sourceIP == "" {
		greeting = "Hello, world!\n"
	} else {
		greeting = fmt.Sprintf("Hello, %s!. V4\n", sourceIP)
	}

	return events.APIGatewayProxyResponse{
		Body:       greeting,
		StatusCode: 200,
	}, nil
}

func (h *APIGatewayV2Handler) PutBirthdayHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	userName := request.PathParameters["username"]
	if userName == "" {
		return events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "username is required",
		}, nil
	}

	var putBirthdayCommand usecases.PutBirthdayCommand
	err := json.Unmarshal([]byte(request.Body), &putBirthdayCommand)

	if err != nil {
		log.Println("Error unmarshalling request body", err)
		return events.APIGatewayProxyResponse{}, err
	}

	if putBirthdayCommand.DateOfBirth == "" {
		return events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "dateOfBirth is required",
		}, nil
	}

	putBirthdayCommand.Username = userName

	response, err := h.PutBirthdayUseCase.Execute(putBirthdayCommand)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	if response.ErrorType != "" {
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

	userName := request.PathParameters["username"]

	getBirthdayResponseUc, err := h.GetBirthdayUseCase.Execute(userName)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	if getBirthdayResponseUc.ErrorType != "" {
		return events.APIGatewayProxyResponse{
			Body:       getBirthdayResponseUc.ErrorMsg,
			StatusCode: getBirthdayResponseUc.ErrorCode,
		}, nil
	}

	getBirthdayResponse := GetBirthdayResponse{
		Message: getBirthdayResponseUc.Message,
	}

	body, err := json.Marshal(getBirthdayResponse)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	return events.APIGatewayProxyResponse{
		Body:       string(body),
		Headers:    map[string]string{"Content-Type": "application/json"},
		StatusCode: 200,
	}, nil
}
