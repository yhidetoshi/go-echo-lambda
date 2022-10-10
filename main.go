package main

import (
	"context"
	"log"
	"yhidetoshi/go-echo-lambda/api/healthcheck"
	"yhidetoshi/go-echo-lambda/handler/auth"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	echoadapter "github.com/awslabs/aws-lambda-go-api-proxy/echo"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var echoLambda *echoadapter.EchoLambda

func init() {
	log.Printf("echo cold start")

	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(auth.BasicAuth())
	e.GET("/api/healthcheck", healthcheck.Healthcheck)

	echoLambda = echoadapter.New(e)
}

func main() {
	lambda.Start(Handler)
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return echoLambda.ProxyWithContext(ctx, req)
}
