package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	Name string `json:"what is your name?"`
	Age  int    `json:"what is your age?"`
}

type MyResponse struct {
	Message string `json:"answer"`
}

func main() {
	lambda.Start(HandleLambdaEvent)
}

func HandleLambdaEvent(event MyEvent) (MyResponse, error) {
	return MyResponse{Message: fmt.Sprintf("%s is %d years old!\n", event.Name, event.Age)}, nil
}
