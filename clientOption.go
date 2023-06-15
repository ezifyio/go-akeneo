package goakeneo

import (
	"net/url"
	"time"

	"go.uber.org/ratelimit"
)

// Option is client option function
type Option func(*Client)

// WithBaseURL sets the base URL of the Akeneo API
func WithBaseURL(u string) Option {
	return func(c *Client) {
		c.baseURL, _ = url.Parse(u)
	}
}

// WithRateLimit sets the rate limit of the Akeneo API
func WithRateLimit(limit int, t time.Duration) Option {
	return func(c *Client) {
		c.limiter = ratelimit.New(limit, ratelimit.WithoutSlack, ratelimit.Per(t))
	}
}

// WithPimVersion sets the version of the Akeneo PIM
func WithPimVersion(v int) Option {
	return func(c *Client) {
		c.osVersion = v
	}
}
