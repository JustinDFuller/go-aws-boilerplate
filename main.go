package main

import (
	"github.com/aws/aws-lambda-go/lambda"
)

type HeaderResponse struct {
  AccessControlAllowOrigin string `json:"Access-Control-Allow-Origin"`
  AccessControlAllowCredentials bool `json:"Access-Control-Allow-Credentials"`
  ContentType string `json:"Content-Type"`
}

type Response struct {
  IsBase64Encoded bool `json:"isBase64Encoded"`
  StatusCode int `json:"statusCode"`
  Headers HeaderResponse `json:"headers"`
  Body string `json:"body"`
}

func HandleRequest() (Response, error) {
  var response Response;
  response.IsBase64Encoded = false;
  response.StatusCode = 200;
  response.Headers = HeaderResponse{
    AccessControlAllowOrigin : "*",
    AccessControlAllowCredentials : true,
    ContentType: "application/json",
  }
  response.Body = "Hello world"
  return response, nil
}

func main() {
	lambda.Start(HandleRequest)
}