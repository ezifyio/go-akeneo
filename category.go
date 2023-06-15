package goakeneo

const (
	categoryBasePath = "/api/rest/v1/categories"
)

// CategoryService is an interface for interacting with the Akeneo Category API.
type CategoryService interface {
	ListWithPagination(options any) ([]Category, Links, error)
}

type categoryOp struct {
	client *Client
}

// ListWithPagination lists categories with pagination
func (c *categoryOp) ListWithPagination(options any) ([]Category, Links, error) {
	categoryResponse := new(categoriesResponse)
	if err := c.client.GET(
		categoryBasePath,
		options,
		nil,
		categoryResponse,
	); err != nil {
		return nil, Links{}, err
	}
	return categoryResponse.Embedded.Items, categoryResponse.Links, nil
}

// categoriesResponse is the struct for a akeneo categories response
type categoriesResponse struct {
	Links       Links         `json:"_links,omitempty" mapstructure:"_links"`
	CurrentPage int           `json:"current_page,omitempty" mapstructure:"current_page"`
	Embedded    categoryItems `json:"_embedded,omitempty" mapstructure:"_embedded"`
}

// categoryItems is the struct for a akeneo category items
type categoryItems struct {
	Items []Category `json:"items,omitempty" mapstructure:"items"`
}
