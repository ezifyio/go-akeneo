package goakeneo

import (
	"encoding/base64"
	"net/http"
	"net/url"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
)

const (
	authBasePath = "api/oauth/v1/token"
)

// AuthService is the interface to implement to authenticate to the Akeneo API
type AuthService interface {
	GrantByPassword() error
	GrantByRefreshToken() error
	ShouldRefreshToken() bool
	AutoRefreshToken() error
}

type authOp struct {
	client *Client
}

// GrantByPassword authenticates to the Akeneo API using the password grant type
func (a *authOp) GrantByPassword() error {
	result := new(authResponse)
	request := authByPasswordRequest{
		GrantType: "password",
		Username:  a.client.connector.UserName,
		Password:  a.client.connector.Password,
	}
	rel, _ := url.Parse(authBasePath)
	// Make the full url based on the relative path
	u := a.client.baseURL.ResolveReference(rel)
	_, err := resty.New().R().
		SetHeader("Content-Type", defaultContentType).
		SetHeader("Authorization", base64BasicAuth(a.client.connector.ClientID, a.client.connector.Secret)).
		SetBody(request).
		SetResult(result).
		Execute(http.MethodPost, u.String())
	if err != nil {
		return errors.Wrap(err, "unable to authenticate to the Akeneo API")
	}
	if err := result.validate(); err != nil {
		return errors.Wrap(err, "invalid response from the Akeneo API")
	}
	a.client.token = result.AccessToken
	a.client.refreshToken = result.RefreshToken
	a.client.tokenExp = time.Now().Add(time.Duration(result.ExpiresIn) * time.Second)
	return nil
}

// GrantByRefreshToken authenticates to the Akeneo API using the refresh token grant type
func (a *authOp) GrantByRefreshToken() error {
	result := new(authResponse)
	request := authByRefreshTokenRequest{
		GrantType:    "refresh_token",
		RefreshToken: a.client.refreshToken,
	}
	if err := a.client.POST(
		authBasePath,
		nil,
		request,
		result,
	); err != nil {
		return err
	}
	if err := result.validate(); err != nil {
		return err
	}
	a.client.token = result.AccessToken
	a.client.refreshToken = result.RefreshToken
	a.client.tokenExp = time.Now().Add(time.Duration(result.ExpiresIn) * time.Second)
	return nil
}

// ShouldRefreshToken returns true if the token should be refreshed
func (a *authOp) ShouldRefreshToken() bool {
	// time.Now is 5 minutes before the actual expiration
	return time.Now().Add(5 * time.Minute).After(a.client.tokenExp)
}

// AutoRefreshToken refreshes the token if needed
func (a *authOp) AutoRefreshToken() error {
	if a.ShouldRefreshToken() {
		err := a.GrantByRefreshToken()
		if err != nil {
			return a.GrantByPassword()
		}
	}
	return nil
}

type authResponse struct {
	AccessToken  string `json:"access_token,omitempty"`
	ExpiresIn    int64  `json:"expires_in,omitempty"`
	Scope        string `json:"scope,omitempty"`
	TokenType    string `json:"token_type,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
}

func (a *authResponse) validate() error {
	if a.AccessToken == "" || a.RefreshToken == "" || a.ExpiresIn == 0 {
		return errors.New("invalid auth response")
	}
	return nil
}

type authByPasswordRequest struct {
	GrantType string `json:"grant_type"`
	Username  string `json:"username,omitempty"`
	Password  string `json:"password,omitempty"`
}

type authByRefreshTokenRequest struct {
	GrantType    string `json:"grant_type"`
	RefreshToken string `json:"refresh_token,omitempty"`
}

func base64BasicAuth(clientID, clientSecret string) string {
	authCredentials := clientID + ":" + clientSecret
	return "Basic " + base64.StdEncoding.EncodeToString([]byte(authCredentials))
}
