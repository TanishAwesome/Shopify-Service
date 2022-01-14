package webhook

import (
	"encoding/json"

	goshopify "github.com/bold-commerce/go-shopify"
)

func CreateWebhook(c *goshopify.Client, topic string, addr string) (*goshopify.Webhook, error) {
	resp, err := c.Webhook.Create(goshopify.Webhook{
		Address: addr,
		Topic:   topic,
		Format:  "json",
	})
	if err != nil {
		return nil, err
	}
	return resp, nil
}
func ListWebhooks(c *goshopify.Client) (string, error) {
	resp, err := c.Webhook.List(nil)
	if err != nil {
		return "", err
	}
	webhjson, _ := json.Marshal(resp)

	return string(webhjson), nil
}
