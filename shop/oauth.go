package shop

import (
	"github.com/TanishAwesome/shopify-service/config"
	goshopify "github.com/bold-commerce/go-shopify"
)

func NewApp(c *config.Config) *goshopify.App {
	App := goshopify.App{
		ApiKey:      c.APIKey,
		ApiSecret:   c.Password,
		RedirectUrl: "https://gi56bayi4g.execute-api.ap-south-1.amazonaws.com/shopify-test-api/callback",
		Scope:       "read_products,read_orders",
	}
	return &App
}
