package jamf

import (
	"net/http"
	"time"

	"github.com/pkg/errors"
)

const (
	uriAuthToken  = "/api/v1/auth/token"
	uriOAuthToken = "/api/oauth/token"
)

// Client ... stores an object to talk with Jamf API
type Client struct {
	username, password, url, clientId, clientSecret string
	token                                           *string
	tokenExpiration                                 *time.Time

	// The Http Client that is used to make requests
	HttpClient       *http.Client
	HttpRetryTimeout time.Duration

	// Option to specify extra headers like User-Agent
	ExtraHeader map[string]string
}

type Option func(*Client)

// WithHttpClient sets the http client to use for requests
func WithHttpClient(client *http.Client) Option {
	return func(c *Client) {
		c.HttpClient = client
	}
}

// WithHttpRetryTimeout sets the timeout for retrying requests
func WithHttpRetryTimeout(timeout time.Duration) Option {
	return func(c *Client) {
		c.HttpRetryTimeout = timeout
	}
}

// WithExtraHeader sets the extra headers to use for requests
func WithExtraHeader(header map[string]string) Option {
	return func(c *Client) {
		c.ExtraHeader = header
	}
}

// WithOAuth sets the client ID and Secret to use OAuth for authentication
func WithOAuth(clientId, clientSecret string) Option {
	return func(c *Client) {
		c.clientId = clientId
		c.clientSecret = clientSecret
	}
}

// WithBasicAuth sets the username and password to use Basic Auth for authentication
func WithBasicAuth(username, password string) Option {
	return func(c *Client) {
		c.username = username
		c.password = password
	}
}

// NewClient ... returns a new jamf.Client which can be used to access the API
func NewClient(url string, options ...Option) (*Client, error) {
	c := &Client{
		url:              url,
		token:            nil,
		HttpClient:       http.DefaultClient,
		HttpRetryTimeout: 60 * time.Second,
		ExtraHeader:      make(map[string]string),
	}

	for _, option := range options {
		option(c)
	}

	if err := c.refreshAuthToken(); err != nil {
		return c, errors.Wrap(err, "Error getting bearer auth token")
	}

	return c, nil
}

func (c *Client) refreshAuthToken() error {
	if c.tokenExpiration != nil {
		if c.tokenExpiration.After(time.Now()) {
			return nil
		}
	}

	c.token = nil

	if c.clientId != "" && c.clientSecret != "" {
		tokenResp, err := c.CreateOAuthToken()
		if err != nil {
			return err
		}

		c.token = tokenResp.Token
		expiration := time.Now().Add(time.Duration(*tokenResp.ExpiresIn) * time.Second)
		c.tokenExpiration = &expiration
	} else {
		tokenResp, err := c.CreateAccessToken()
		if err != nil {
			return err
		}

		c.token = tokenResp.Token
		c.tokenExpiration = tokenResp.Expires
	}

	return nil
}
