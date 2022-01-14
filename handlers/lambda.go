package handlers

import (
	"context"
	"log"

	"github.com/TanishAwesome/shopify-service/config"
	"github.com/TanishAwesome/shopify-service/shop"
	"github.com/aws/aws-lambda-go/events"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	goshopify "github.com/bold-commerce/go-shopify"
	"github.com/gin-gonic/gin"
)

var ShopApp *goshopify.App

func Initiliaze(c *config.Config) {
	ShopApp = shop.NewApp(c)
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var ginLambda *ginadapter.GinLambda
	if ginLambda == nil {
		// Logs sent to AWS CloudWatch Logs
		log.Printf("Gin Started!")
		r := gin.Default()
		r.GET("/oauth", OauthHandler)
		r.GET("/callback", OauthCallbackHandler)
		r.GET("/productlist", ProductListHandler)
		r.POST("/webhooks", WebhookHandler)

		ginLambda = ginadapter.New(r)
	}
	return ginLambda.ProxyWithContext(ctx, req)
}
