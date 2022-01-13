package handlers

import (
	"context"
	"log"

	"github.com/TanishAwesome/shopify-service/config"
	"github.com/TanishAwesome/shopify-service/shop"
	"github.com/aws/aws-lambda-go/events"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
)

var ginLambda *ginadapter.GinLambda
var shopClient *shop.Client

func Initiliaze(c *config.Config) {
	shopClient = shop.NewClient(c)
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if ginLambda == nil {
		// stdout and stderr are sent to AWS CloudWatch Logs
		log.Printf("Gin cold start")
		r := gin.Default()
		r.GET("/productlist", ProductListHandler)

		ginLambda = ginadapter.New(r)
	}
	// If no name is provided in the HTTP request body, throw an error
	return ginLambda.ProxyWithContext(ctx, req)
}
