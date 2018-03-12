@echo off 

set lambda=github.com\aws\aws-lambda-go
set file=%USERPROFILE%\Go\src\%lambda%

echo "Checking if aws-lambda-go exists at %file%"
if EXIST "%file%" (
	echo "aws-lambda-go package already installed."
) else (
	echo "Installing aws-lambda-go"
	go get "%lambda%\lambda"
)

CALL ./cleanup.bat

echo "Creating new dist folder"
mkdir dist

echo "Building main.go as dist\main"
set GOOS=linux
set GOARCH=amd64
go build -o dist\main main.go

echo "Creating zip file"
"%USERPROFILE%\Go\bin\build-lambda-zip.exe" -o dist\main.zip dist\main

echo "Removing main executable file"
rm dist\main