package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"

	"github.com/natnad59/aws-rest-api/resolvers"
)

var ginLambda *ginadapter.GinLambda

func init() {
	router := gin.Default()
	// router.Use(utility.CorsMiddleware())

	gin.SetMode(gin.ReleaseMode)

	router.GET("/users/:id", resolvers.GetUser)
	router.GET("/users", resolvers.GetUsers)
	router.POST("/users", resolvers.PostUser)
	router.DELETE("/users/:id", resolvers.DeleteUser)

	ginLambda = ginadapter.New(router)
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return ginLambda.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(Handler)
}
