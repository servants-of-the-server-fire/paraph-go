package paraph

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const (
	DefaultServerURL = "https://paraph.dev/api/v1"
	EnvAPIKey        = "PARAPH_API_KEY"
	EnvBaseURL       = "PARAPH_BASE_URL"
)

const (
	authHeader   = "Authorization"
	bearerScheme = "Bearer "
)

type Option func(*options)

type options struct {
	baseURL    string
	httpClient *http.Client
	timeout    time.Duration
	userAgent  string
}

func WithBaseURL(url string) Option        { return func(o *options) { o.baseURL = url } }
func WithHTTPClient(c *http.Client) Option { return func(o *options) { o.httpClient = c } }
func WithTimeout(d time.Duration) Option   { return func(o *options) { o.timeout = d } }
func WithUserAgent(ua string) Option       { return func(o *options) { o.userAgent = ua } }

// Client wraps the generated APIClient with ergonomic defaults.
type Client struct {
	raw       *APIClient
	Account   AccountAPI
	Templates TemplatesAPI
	Requests  RequestsAPI
	Signers   SignersAPI
	Webhooks  WebhooksAPI
}

// NewClient creates a client authenticated with the given API key.
//
//	client := paraph.NewClient("your-api-key")
//	resp, _, err := client.Account.GetAccount(ctx).Execute()
func NewClient(apiKey string, opts ...Option) *Client {
	o := options{baseURL: DefaultServerURL}
	for _, opt := range opts {
		opt(&o)
	}

	cfg := NewConfiguration()
	cfg.AddDefaultHeader(authHeader, bearerScheme+apiKey)
	cfg.Servers = ServerConfigurations{{URL: o.baseURL}}

	if o.userAgent != "" {
		cfg.UserAgent = o.userAgent
	}
	if o.httpClient != nil {
		cfg.HTTPClient = o.httpClient
	} else if o.timeout > 0 {
		cfg.HTTPClient = &http.Client{Timeout: o.timeout}
	}

	raw := NewAPIClient(cfg)
	return &Client{
		raw:       raw,
		Account:   raw.AccountAPI,
		Templates: raw.TemplatesAPI,
		Requests:  raw.RequestsAPI,
		Signers:   raw.SignersAPI,
		Webhooks:  raw.WebhooksAPI,
	}
}

// NewClientFromEnv reads PARAPH_API_KEY (required) and PARAPH_BASE_URL
// (optional, defaults to DefaultServerURL) from environment variables.
//
//	client, err := paraph.NewClientFromEnv()
func NewClientFromEnv(opts ...Option) (*Client, error) {
	apiKey := os.Getenv(EnvAPIKey)
	if apiKey == "" {
		return nil, fmt.Errorf("paraph: %s environment variable is not set", EnvAPIKey)
	}
	if base := os.Getenv(EnvBaseURL); base != "" {
		opts = append([]Option{WithBaseURL(base)}, opts...)
	}
	return NewClient(apiKey, opts...), nil
}
