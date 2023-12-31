package goakeneo

import (
	"path"
)

const (
	attributeBasePath = "/api/rest/v1/attributes"
)

// AttributeService is an interface for interfacing with the attribute
type AttributeService interface {
	ListWithPagination(options any) ([]Attribute, Links, error)
	GetAttribute(code string, options any) (*Attribute, error)
	GetAttributeOptions(code string, options any) ([]AttributeOption, Links, error)
}

// attributeOp handles communication with the attribute related methods of the Akeneo API.
type attributeOp struct {
	client *Client
}

// ListWithPagination lists attributes with pagination
func (c *attributeOp) ListWithPagination(options any) ([]Attribute, Links, error) {
	attributeResponse := new(AttributesResponse)
	if err := c.client.GET(
		attributeBasePath,
		options,
		nil,
		attributeResponse,
	); err != nil {
		return nil, Links{}, err
	}
	return attributeResponse.Embedded.Items, attributeResponse.Links, nil
}

// GetAttribute gets an attribute by code
func (c *attributeOp) GetAttribute(code string, options any) (*Attribute, error) {
	sourcePath := path.Join(attributeBasePath, code)
	attribute := new(Attribute)
	if err := c.client.GET(
		sourcePath,
		options,
		nil,
		attribute,
	); err != nil {
		return nil, err
	}
	return attribute, nil
}

// GetAttributeOptions gets an attribute's options by code
func (c *attributeOp) GetAttributeOptions(code string, options any) ([]AttributeOption, Links, error) {
	sourcePath := path.Join(attributeBasePath, code, "options")
	attributeOptionsResponse := new(AttributeOptionsResponse)
	if err := c.client.GET(
		sourcePath,
		options,
		nil,
		attributeOptionsResponse,
	); err != nil {
		return nil, Links{}, err
	}
	return attributeOptionsResponse.Embedded.Items, attributeOptionsResponse.Links, nil
}

// AttributesResponse is the struct for a akeneo attributes response
type AttributesResponse struct {
	Links       Links          `json:"_links" mapstructure:"_links"`
	CurrentPage int            `json:"current_page" mapstructure:"current_page"`
	Embedded    attributeItems `json:"_embedded" mapstructure:"_embedded"`
}

type attributeItems struct {
	Items []Attribute `json:"items" mapstructure:"items"`
}

// AttributeOptionsResponse is the struct for a akeneo attribute options response
type AttributeOptionsResponse struct {
	Links       Links                `json:"_links" mapstructure:"_links"`
	CurrentPage int                  `json:"current_page" mapstructure:"current_page"`
	Embedded    attributeOptionItems `json:"_embedded" mapstructure:"_embedded"`
}

type attributeOptionItems struct {
	Items []AttributeOption `json:"items" mapstructure:"items"`
}
