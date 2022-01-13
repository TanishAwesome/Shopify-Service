package handlers

import "github.com/gin-gonic/gin"

func ProductListHandler(c *gin.Context) {
	products, err := shopClient.GetProductsList()
	if err != nil {
		return
	}
	c.IndentedJSON(200, products)
}

// TODO: ADD More Handlers
