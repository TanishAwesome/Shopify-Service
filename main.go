package main

import (
	"fmt"
	"os"

	"github.com/TanishAwesome/shopify-service/config"
	"github.com/TanishAwesome/shopify-service/handlers"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	conf, err := config.Load()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	handlers.Initiliaze(conf)
	lambda.Start(handlers.Handler)
}
