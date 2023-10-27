// Package fga includes client libraries to interact with openfga authorization
package fga

import (
	ofgaclient "github.com/openfga/go-sdk/client"
	"github.com/openfga/go-sdk/credentials"
	"go.uber.org/zap"
)

// Client is an event bus client with some configuration
type Client struct {
	c      *ofgaclient.OpenFgaClient
	config ofgaclient.ClientConfiguration
	logger *zap.SugaredLogger
}

// Option is a functional configuration option for openFGA client
type Option func(c *Client)

// NewClient returns a wrapped OpenFGA API client ensuring all calls are made
// to the provided authorization model (id) and returns what is necessary.
func NewClient(host string, opts ...Option) (*Client, error) {
	if host == "" {
		return nil, ErrFGAMissingHost
	}

	// The api host is the only required field when setting up a new FGA client connection
	client := Client{
		config: ofgaclient.ClientConfiguration{
			ApiHost: host,
		},
	}

	for _, opt := range opts {
		opt(&client)
	}

	fgaClient, err := ofgaclient.NewSdkClient(&client.config)
	if err != nil {
		return nil, err
	}

	client.c = fgaClient

	return &client, err
}

// WithScheme sets the open fga scheme, defaults to "https"
func WithScheme(scheme string) Option {
	return func(c *Client) {
		c.config.ApiScheme = scheme
	}
}

// WithStoreID sets the store IDs, not needed when calling `CreateStore` or `ListStores`
func WithStoreID(storeID string) Option {
	return func(c *Client) {
		c.config.StoreId = storeID
	}
}

// WithAuthorizationModelID sets the authorization model ID
func WithAuthorizationModelID(authModelID string) Option {
	return func(c *Client) {
		c.config.AuthorizationModelId = &authModelID
	}
}

// WithToken sets the client credentials
func WithToken(token string) Option {
	return func(c *Client) {
		c.config.Credentials = &credentials.Credentials{
			Method: credentials.CredentialsMethodApiToken,
			Config: &credentials.Config{
				ApiToken: token,
			},
		}
	}
}

// WithLogger sets logger
func WithLogger(l *zap.SugaredLogger) Option {
	return func(c *Client) {
		c.logger = l
	}
}
