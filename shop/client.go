package shop

import (
	"github.com/TanishAwesome/shopify-service/config"
	goshopify "github.com/bold-commerce/go-shopify"
)

type Client struct {
	*goshopify.Client
}

func NewClient(c *config.Config) *Client {
	return &Client{
		Client: goshopify.NewClient(
			goshopify.App{
				ApiKey:   c.APIKey,
				Password: c.Password,
			},
			c.ShopName,
			"", // Private Apps don't need a token
		),
	}
}
