package awsLambdaImplementation

import (
	"github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest() (string, error) {
  return "Hello world", nil
}

func main() {
	lambda.Start(HandleRequest)
}