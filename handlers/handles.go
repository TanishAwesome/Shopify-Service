package handlers

import (
	"fmt"
	"net/http"

	"github.com/TanishAwesome/shopify-service/config"
	"github.com/TanishAwesome/shopify-service/shop"
	"github.com/TanishAwesome/shopify-service/webhook"
	goshopify "github.com/bold-commerce/go-shopify"
	"github.com/gin-gonic/gin"
)

var ShopClient *goshopify.Client

func OauthHandler(c *gin.Context) {
	shopName := c.Query("shop")
	state := "nonce"
	authUrl := ShopApp.AuthorizeUrl(shopName, state)
	c.Redirect(http.StatusFound, authUrl)
	c.Abort()
}

// Fetch a permanent access token in the callback
func OauthCallbackHandler(c *gin.Context) {
	// Check that the callback signature is valid
	if ok, err := ShopApp.VerifyAuthorizationURL(c.Request.URL); !ok {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	query := c.Request.URL.Query()
	shopName := query.Get("shop")
	code := query.Get("code")
	accessToken, err := ShopApp.GetAccessToken(shopName, code)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	// Create a Pointer to Client using the generated token
	ShopClient = ShopApp.NewClient(shopName, accessToken)
	config.Cfile.Access_Token = accessToken

	// Create webhooks after Oauth is complete
	// TODO: Don't make webhook if they are already in webhook list
	_, err = webhook.CreateWebhook(ShopClient, "orders/create", config.Cfile.Webhook_Addr)
	if err != nil {
		fmt.Println(err)
	}

	c.String(http.StatusOK, "The App has been authenticated, access token: "+accessToken)
}

func WebhookHandler(c *gin.Context) {
	//var resp WebhookResp
	json, err := c.GetRawData()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(json))
}

func ProductListHandler(c *gin.Context) {
	products, err := shop.GetProductsList(ShopClient)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	c.IndentedJSON(200, products)
}

// TODO: ADD More Handlers
