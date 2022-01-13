package shop

import goshopify "github.com/bold-commerce/go-shopify"

func (c *Client) GetProductsList() ([]goshopify.Product, error) {
	products, err := c.Product.List(nil)
	if err != nil {
		return nil, err
	}
	return products, nil
}

// TODO: ADD More Products Related functions
