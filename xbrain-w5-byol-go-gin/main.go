package main

import (
	"byol-go-gin/server"

	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	adapter := ginadapter.NewV2(server.New())
	lambda.Start(adapter.ProxyWithContext)
}
