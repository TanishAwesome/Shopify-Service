package config

import (
	"fmt"
	"os"
)

type Config struct {
	ShopName string
	APIKey   string
	Password string
}

// Load all the Required config variables from Env Variables
func Load() (*Config, error) {
	cfile := &Config{}
	cfile.ShopName = os.Getenv("SHOP_NAME")
	if cfile.ShopName == "" {
		return nil, fmt.Errorf("SHOP_NAME is not set properly as env var")
	}
	cfile.APIKey = os.Getenv("API_KEY")
	if cfile.APIKey == "" {
		return nil, fmt.Errorf("API_KEY is not set properly as env var")
	}
	cfile.Password = os.Getenv("PASSWORD")
	if cfile.Password == "" {
		return nil, fmt.Errorf("PASSWORD is not set properly as env var")
	}
	return cfile, nil
}
