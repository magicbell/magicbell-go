package clientconfig

import (
	"time"

	"github.com/magicbell/magicbell-go/pkg/project-client/internal/clients/rest/hooks"
)

// Config holds all configuration parameters for the SDK client.
// It manages base URL, timeout, authentication credentials, and custom hooks.
type Config struct {
	BaseUrl     *string
	Timeout     *time.Duration
	AccessToken *string
	HookParams  map[string]string
	hook        hooks.Hook
}

// NewConfig creates a new Config instance with default values.
// Sets the base URL to the default environment and timeout to 10 seconds.
func NewConfig() Config {
	baseUrl := DEFAULT_ENVIRONMENT
	timeout := time.Second * 10
	newConfig := Config{
		BaseUrl:    &baseUrl,
		Timeout:    &timeout,
		HookParams: make(map[string]string),
	}

	return newConfig
}

func (c *Config) SetBaseUrl(baseUrl string) {
	c.BaseUrl = &baseUrl
}

func (c *Config) GetBaseUrl() string {
	return *c.BaseUrl
}

func (c *Config) SetTimeout(timeout time.Duration) {
	c.Timeout = &timeout
}

func (c *Config) GetTimeout() time.Duration {
	return *c.Timeout
}

func (c *Config) SetAccessToken(accessToken string) {
	c.AccessToken = &accessToken
}

func (c *Config) GetAccessToken() string {
	return *c.AccessToken
}
