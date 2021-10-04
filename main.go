package main

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"time"
)

type Person struct {
	Name string `json:"name"`
}

type Error struct {
	Msg       string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
}

var people = []Person{{Name: "Hello"}, {Name: "World"}}

func main() {
	lambda.Start(HandleRequest)
}

func HandleRequest(request events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	m, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err.Error())
		error := Error{
			Msg:       err.Error(),
			Timestamp: time.Now(),
		}
		m, _ := json.Marshal(error)
		return events.APIGatewayV2HTTPResponse{StatusCode: 500, Body: string(m)}, nil
	}
	if request.RequestContext.HTTP.Method == "POST" {
		return events.APIGatewayV2HTTPResponse{StatusCode: 200, Body: "TODO: POST response"}, nil
	}
	return events.APIGatewayV2HTTPResponse{StatusCode: 200, Body: string(m), Headers: map[string]string{
		"Content-Type": "application/json",
	}}, nil
}
