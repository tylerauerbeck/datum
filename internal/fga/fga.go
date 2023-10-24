// Package fga includes client libraries to interact with openfga authorization
package fga

import (
	"strings"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	ofgaclient "github.com/openfga/go-sdk/client"
	"github.com/openfga/go-sdk/credentials"
	"go.uber.org/zap"
)

const (
	bearerPrefix = "Bearer "
)

// Client is an event bus client with some configuration
type Client struct {
	o      *ofgaclient.OpenFgaClient
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

	client.o = fgaClient

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

// Authz handles supporting authorization checks
type Authz struct {
	logger *zap.SugaredLogger
	client *Client
}

func New(l *zap.SugaredLogger, c *Client) Authz {
	return Authz{
		logger: l,
		client: c,
	}
}

// Middleware produces echo middleware to handle authorization checks
func (a *Authz) Middleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			actor := c.Get("user").(*jwt.Token)
			if actor == nil {
				a.logger.Infof("unauthorized access")

				return echo.ErrUnauthorized
			}

			authHeader := strings.TrimSpace(c.Request().Header.Get(echo.HeaderAuthorization))

			if len(authHeader) <= len(bearerPrefix) {
				a.logger.Infof("unauthorized access")

				return echo.ErrUnauthorized
			}

			if !strings.EqualFold(authHeader[:len(bearerPrefix)], bearerPrefix) {
				a.logger.Infof("unauthorized access")

				return echo.ErrUnauthorized
			}

			// testing stuff
			tuple := Tuple{
				Subject:  "user:sfunkhouser",
				Relation: "member",
				Object:   "organization:datum",
			}

			allowed, err := a.client.CheckTuple(c.Request().Context(), tuple)
			if err != nil {
				a.logger.Errorf("error checking access, %v", err)

				return echo.ErrUnauthorized
			}

			if !allowed {
				a.logger.Infof("unauthorized access")

				return echo.ErrUnauthorized
			}

			a.logger.Infof("access granted")

			return next(c)
		}
	}
}
