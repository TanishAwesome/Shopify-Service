package config

import (
	"fmt"
	"os"
)

type Config struct {
	ShopName     string
	APIKey       string
	Password     string
	Access_Token string
	Webhook_Addr string
}

var Cfile *Config

// Load all the Required config variables from Env Variables
func Load() (*Config, error) {
	Cfile = &Config{}
	Cfile.ShopName = os.Getenv("SHOP_NAME")
	if Cfile.ShopName == "" {
		return nil, fmt.Errorf("SHOP_NAME is not set properly as env var")
	}
	Cfile.APIKey = os.Getenv("API_KEY")
	if Cfile.APIKey == "" {
		return nil, fmt.Errorf("API_KEY is not set properly as env var")
	}
	Cfile.Password = os.Getenv("PASSWORD")
	if Cfile.Password == "" {
		return nil, fmt.Errorf("PASSWORD is not set properly as env var")
	}
	Cfile.Webhook_Addr = os.Getenv("WEBHOOK_ADDR")
	if Cfile.Password == "" {
		return nil, fmt.Errorf("WEBHOOK_ADDR is not set properly as env var")
	}
	return Cfile, nil
}
