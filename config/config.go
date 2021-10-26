package config

import (
	"os"
	"time"

	"github.com/evalphobia/bigdatacloud-go/client"
)

const (
	envAPIKey   = "BIGDATACLOUD_APIKEY"
	envLanguage = "BIGDATACLOUD_LANGUAGE"

	defaultLanguage = "en"
)

// Config contains parameters for BigDataCloud.
type Config struct {
	APIKey   string
	Lanugage string

	Debug   bool
	Timeout time.Duration
}

// Client generates client.Client from Config data.
func (c Config) Client() (*client.Client, error) {
	cli := client.New()
	cli.SetOption(client.Option{
		Debug:   c.Debug,
		Timeout: c.Timeout,
	})
	cli.SetAPIKey(c.getAPIKey())
	return cli, nil
}

// getAPIKey returns API Key from Config data or Environment variables.
func (c Config) getAPIKey() string {
	apiKey := os.Getenv(envAPIKey)
	if apiKey != "" {
		return apiKey
	}
	return c.APIKey
}

// GetLanguage returns lanugage from Config data or Environment variables.
func (c Config) GetLanguage() string {
	lang := os.Getenv(envLanguage)
	if lang != "" {
		return lang
	}
	if c.Lanugage != "" {
		return c.Lanugage
	}
	return defaultLanguage
}
