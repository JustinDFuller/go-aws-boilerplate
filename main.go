package main

import (
  "encoding/json"
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

var response = Response{
  IsBase64Encoded: false,
  StatusCode: 200,
  Headers: HeaderResponse{
    AccessControlAllowOrigin : "*",
    AccessControlAllowCredentials : true,
    ContentType: "application/json",
  },
};

func HandleRequest() (Response, error) {
  stringifiedBody, _ := json.Marshal("Hello world")
  response.Body = string(stringifiedBody)
  return response, nil
}

func main() {
	lambda.Start(HandleRequest)
}