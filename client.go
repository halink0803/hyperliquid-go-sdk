// Package hyperliquid ...
package hyperliquid

import "net/http"

const baseURLConst = "https://api.hyperliquid.xyz"

// Client define API client
type Client struct {
	BaseURL    string
	HTTPClient *http.Client
}

func apiGetBaseURL() string {
	return baseURLConst
}

// NewClient ...
func NewClient() *Client {
	return &Client{
		HTTPClient: http.DefaultClient,
		BaseURL:    apiGetBaseURL(),
	}
}
