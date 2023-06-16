package goakeneo

const (
	localeBasePath = "/api/rest/v1/locales"
)

// LocaleService is the interface to interact with the Akeneo Locale API
type LocaleService interface {
	ListWithPagination(options any) ([]Locale, Links, error)
}

type localeOp struct {
	client *Client
}

// ListWithPagination lists locales with pagination
func (c *localeOp) ListWithPagination(options any) ([]Locale, Links, error) {
	localeResponse := new(LocalesResponse)
	if err := c.client.GET(
		localeBasePath,
		options,
		nil,
		localeResponse,
	); err != nil {
		return nil, Links{}, err
	}
	return localeResponse.Embedded.Items, localeResponse.Links, nil
}

// LocalesResponse is the struct for a akeneo locales response
type LocalesResponse struct {
	Links       Links       `json:"_links" mapstructure:"_links"`
	CurrentPage int         `json:"current_page" mapstructure:"current_page"`
	Embedded    localeItems `json:"_embedded" mapstructure:"_embedded"`
}

// localeItems is the struct for a akeneo locale items response
type localeItems struct {
	Items []Locale `json:"items" mapstructure:"items"`
}
