package shop

import (
	"fmt"

	goshopify "github.com/bold-commerce/go-shopify"
)

func GetProductsList(c *goshopify.Client) ([]goshopify.Product, error) {
	fmt.Println("Client is: ")
	fmt.Println(c)
	if c == nil {
		fmt.Println("Yes Empty!")
		return nil, fmt.Errorf("client is empty, connect to the store first using Oauth")
	}
	products, err := c.Product.List(nil)
	if err != nil {
		return nil, err
	}
	return products, nil
}

// TODO: ADD More Products Related functions
