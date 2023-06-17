package goakeneo

import (
	"github.com/pkg/errors"
	"path"
)

const (
	productModelBasePath = "/api/rest/v1/product-models"
)

type ProductModelService interface {
	ListWithPagination(options any) ([]ProductModel, Links, error)
	GetProductModel(code string, options any) (*ProductModel, error)
	Crate(pm ProductModel) error
}

type productModelOp struct {
	client *Client
}

// Crate creates a product model
func (p *productModelOp) Crate(pm ProductModel) error {
	if err := pm.validateBeforeCreate(); err != nil {
		return errors.Wrap(err, "failed to validate product model before create")
	}
	if err := p.client.POST(
		productModelBasePath,
		nil,
		pm,
		nil,
	); err != nil {
		return err
	}
	return nil
}

// ListWithPagination lists product models with pagination
func (p *productModelOp) ListWithPagination(options any) ([]ProductModel, Links, error) {
	productModelResponse := new(ProductModelsResponse)
	if err := p.client.GET(
		productModelBasePath,
		options,
		nil,
		productModelResponse,
	); err != nil {
		return nil, Links{}, err
	}
	return productModelResponse.Embedded.Items, productModelResponse.Links, nil
}

// GetProductModel gets a product model by code
func (p *productModelOp) GetProductModel(code string, options any) (*ProductModel, error) {
	sourcePath := path.Join(productModelBasePath, code)
	productModel := new(ProductModel)
	if err := p.client.GET(
		sourcePath,
		options,
		nil,
		productModel,
	); err != nil {
		return nil, err
	}
	return productModel, nil
}

// ProductModelsResponse is the struct for the response of the ListWithPagination function
type ProductModelsResponse struct {
	Links       Links             `json:"_links" mapstructure:"_links"`
	CurrentPage int               `json:"current_page" mapstructure:"current_page"`
	Embedded    productModelItems `json:"_embedded" mapstructure:"_embedded"`
}

type productModelItems struct {
	Items []ProductModel `json:"items" mapstructure:"items"`
}
