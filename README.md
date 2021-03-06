# go-aws-boilerplate
Boilerplate for creating a Golang AWS Lambda function that works with AWS API Gateway.

## What
This boilerplate provides a working template for AWS Lambda + AWS API Gateway. It is written in Go.

It provides a template that shapes the response so that it works correctly with API Gateway. Which expects a specific response shape.

## Scripts

### Build

Will create a main.exe in a dist folder. It looks for a main.go file in the root of the project.

```shell
./build
```

### Cleanup

Will remove the dist folder.

```shell
./cleanup
```