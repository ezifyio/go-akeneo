package goakeneo

import (
	"github.com/go-resty/resty/v2"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/pkg/errors"
	"go.uber.org/ratelimit"
)

// Connector is the struct to use to store the Akeneo connection information
type Connector struct {
	ClientID string `json:"client_id" mapstructure:"client_id"`
	Secret   string `json:"secret" mapstructure:"secret"`
	UserName string `json:"username" mapstructure:"username"`
	Password string `json:"password" mapstructure:"password"`
}

// NewClient creates a new Akeneo client
func (c Connector) NewClient(opts ...Option) (*Client, error) {
	return NewClient(c, opts...)
}

// Client is the main struct to use to interact with the Akeneo API
type Client struct {
	connector    Connector
	baseURL      *url.URL
	httpClient   *http.Client
	token        string            // token is the access token
	refreshToken string            // refreshToken is the refresh token
	tokenExp     time.Time         // tokenExp is the token expiration time,5 minutes before the actual expiration
	osVersion    int               // osVersion is the version of the OS
	limiter      ratelimit.Limiter // limiter, default 5 requests per second
	Auth         AuthService
	Product      ProductService
	Family       FamilyService
	Attribute    AttributeService
	Category     CategoryService
	Channel      ChannelService
	Locale       LocaleService
	Media        MediaFileService
}

func (c *Client) validate() error {
	if c.baseURL == nil {
		return errors.New("baseURL is nil")
	}
	switch {
	case c.connector.ClientID == "":
		return errors.New("clientID is empty")
	case c.connector.Secret == "":
		return errors.New("secret is empty")
	case c.connector.UserName == "":
		return errors.New("username is empty")
	case c.connector.Password == "":
		return errors.New("password is empty")
	default:
	}
	if _, ok := pimVersionMap[c.osVersion]; !ok {
		return errors.Errorf("invalid osVersion %d", c.osVersion)
	}
	return nil
}

func (c *Client) init() error {
	if c.limiter == nil {
		c.limiter = ratelimit.New(defaultRateLimit, ratelimit.WithoutSlack, ratelimit.Per(time.Second))
	}
	if err := c.Auth.GrantByPassword(); err != nil {
		return err
	}
	return nil
}

// NewClient creates a new Akeneo client
func NewClient(con Connector, opts ...Option) (*Client, error) {

	c := &Client{
		httpClient: &http.Client{
			Timeout: defaultHTTPTimeout,
			Transport: &http.Transport{
				MaxIdleConns: 10,
			},
		},
		connector: con,
		osVersion: defaultVersion,
	}
	for _, opt := range opts {
		opt(c)
	}
	if err := c.validate(); err != nil {
		return nil, err
	}
	// Set services
	c.Auth = &authOp{c}
	c.Product = &productOp{c}
	c.Family = &familyOp{c}
	c.Attribute = &attributeOp{c}
	c.Category = &categoryOp{c}
	c.Channel = &channelOp{c}
	c.Locale = &localeOp{c}
	c.Media = &mediaOp{c}
	if err := c.init(); err != nil {
		return nil, err
	}
	return c, nil
}

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

// WithVersion sets the version of the Akeneo API
func WithVersion(v int) Option {
	return func(c *Client) {
		c.osVersion = v
	}
}

// createAndDo create a request and get the headers
func (c *Client) createAndDo(method, relPath string, opts, data, result any) error {
	if err := c.Auth.AutoRefreshToken(); err != nil {
		return err
	}
	rel, err := url.Parse(relPath)
	if err != nil {
		return err
	}
	// Make the full url based on the relative path
	u := c.baseURL.ResolveReference(rel)

	var errResp ErrorResponse
	request := resty.NewWithClient(c.httpClient).R().
		SetHeader("Content-Type", defaultContentType).
		SetHeader("Accept", defaultAccept).
		SetHeader("User-Agent", defaultUserAgent).
		SetAuthToken(c.token).
		SetResult(result).
		SetError(&errResp)
	if opts != nil {
		if v, ok := opts.(url.Values); ok {
			request.SetQueryParamsFromValues(v)
		} else {
			// check if opts is a struct or a pointer to a struct
			t := reflect.TypeOf(opts)
			if t.Kind() == reflect.Ptr && t.Elem().Kind() == reflect.Struct || t.Kind() == reflect.Struct {
				v, err := structToURLValues(opts)
				if err != nil {
					return errors.Wrap(err, "unable to convert struct to url values")
				}
				request.SetQueryParamsFromValues(v)
			} else {
				return errors.New("opts must be a struct or a pointer to a struct or a url.Values")
			}
		}
	}
	if data != nil {
		request.SetBody(data)
	}
	// rate limit
	c.limiter.Take()
	resp, err := request.Execute(method, u.String())
	if err != nil {
		return errors.Wrap(err, "resty execute error")
	}
	// see : https://api.akeneo.com/documentation/responses.html
	if resp.IsError() {
		return errors.Errorf("request error : %s", errResp.Message)
	}
	return nil
}

// GET creates a get request and execute it
// result must be a pointer to a struct
func (c *Client) GET(relPath string, ops, data, result any) error {
	err := c.createAndDo(http.MethodGet, relPath, ops, data, result)
	if err != nil {
		return errors.Wrap(err, "create and do error")
	}
	return nil
}

// POST creates a post request and execute it
// result must be a pointer to a struct
func (c *Client) POST(relPath string, ops, data, result any) error {
	err := c.createAndDo(http.MethodPost, relPath, ops, data, result)
	if err != nil {
		return errors.Wrap(err, "create and do error")
	}
	return nil
}
