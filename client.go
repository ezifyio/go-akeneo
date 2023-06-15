package goakeneo

import (
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
	"go.uber.org/ratelimit"
)

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
}

func (c *Client) init() error {

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
	if c.httpClient == nil {
		c.httpClient = http.DefaultClient
	}
	if _, ok := pimVersionMap[c.osVersion]; !ok {
		return errors.Errorf("invalid osVersion %d", c.osVersion)
	}
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
		},
		connector: con,
		osVersion: AkeneoPimVersion6,
	}
	for _, opt := range opts {
		opt(c)
	}
	c.Auth = &authOp{c}
	if err := c.init(); err != nil {
		return nil, err
	}
	return c, nil
}

// createAndDoGetHeaders create a request and get the headers
func (c *Client) createAndDoGetHeaders(method, relPath string, opts, data, result any) (http.Header, error) {
	if err := c.Auth.AutoRefreshToken(); err != nil {
		return nil, err
	}
	rel, err := url.Parse(relPath)
	if err != nil {
		return nil, err
	}
	// Make the full url based on the relative path
	u := c.baseURL.ResolveReference(rel)

	request := resty.NewWithClient(c.httpClient).R().
		SetHeader("Content-Type", defaultContentType).
		SetHeader("Accept", defaultAccept).
		SetHeader("User-Agent", defaultUserAgent).
		SetAuthToken(c.token).
		SetResult(result)
	if opts != nil {
		if v, ok := opts.(url.Values); ok {
			request.SetQueryParamsFromValues(v)
		} else {
			// check if opts is a struct or a pointer to a struct
			t := reflect.TypeOf(opts)
			if t.Kind() == reflect.Ptr && t.Elem().Kind() == reflect.Struct || t.Kind() == reflect.Struct {
				v, err := structToURLValues(opts)
				if err != nil {
					return nil, errors.Wrap(err, "unable to convert struct to url values")
				}
				request.SetQueryParamsFromValues(v)
			} else {
				return nil, errors.New("opts must be a struct or a pointer to a struct or a url.Values")
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
		return nil, errors.Wrap(err, "resty execute error")
	}
	// todo: check status code
	return resp.Header(), nil
}

// GET creates a get request and execute it
func (c *Client) GET(relPath string, ops, data, result any) error {
	_, err := c.createAndDoGetHeaders(http.MethodGet, relPath, ops, data, result)
	if err != nil {
		return errors.Wrap(err, "create and do get headers error")
	}
	return nil
}

// POST creates a post request and execute it
func (c *Client) POST(relPath string, ops, data, result any) error {
	_, err := c.createAndDoGetHeaders(http.MethodPost, relPath, ops, data, result)
	if err != nil {
		return errors.Wrap(err, "create and do get headers error")
	}
	return nil
}
