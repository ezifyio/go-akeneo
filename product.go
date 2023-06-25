package goakeneo

import (
	"context"
	"path"

	"github.com/pkg/errors"
)

const (
	productBasePath = "/api/rest/v1/products"
	// product uuid path since akeneo 7
	productUUIDBasePath = "/api/rest/v1/products-uuid"
)

// ProductService is the interface to interact with the Akeneo Product API
type ProductService interface {
	GetAllProducts(ctx context.Context, options any) (<-chan Product, chan error)
	ListWithPagination(options any) ([]Product, Links, error)
	GetProduct(id string, options any) (*Product, error)
	UpdateOrCreateProducts(products []Product) (PatchProductResponse, error)
}

type productOp struct {
	client *Client
}

// GetAllProducts lists all products, returns a channel to iterate over products
func (p *productOp) GetAllProducts(ctx context.Context, options any) (<-chan Product, chan error) {
	prodChan := make(chan Product, 1)
	errChan := make(chan error)
	go func() {
		defer close(errChan)
		defer close(prodChan)
		defer func() {
			if r := recover(); r != nil {
				err := errors.Errorf("unable to get all products: %v", r)
				errChan <- err
			}
		}()
		prods, links, err := p.ListWithPagination(options)
		for {
			if err != nil {
				errChan <- err
				return
			}
			select {
			case <-ctx.Done():
				return
			default:
			}
			for _, prod := range prods {
				prodChan <- prod
			}
			if !links.HasNext() {
				break
			}
			prods, links, err = p.ListWithPagination(links.NextOptions())
		}
	}()
	return prodChan, errChan
}

// ListWithPagination lists products with pagination
func (p *productOp) ListWithPagination(options any) ([]Product, Links, error) {
	// products-uuid path since akeneo 7
	basePath := productBasePath
	if p.client.osVersion >= AkeneoPimVersion7 {
		basePath = productUUIDBasePath
	}
	productResponse := new(ProductsResponse)
	if err := p.client.GET(
		basePath,
		options,
		nil,
		productResponse,
	); err != nil {
		return nil, Links{}, err
	}
	return productResponse.Embedded.Items, productResponse.Links, nil
}

// GetProduct gets a product by its identifier
func (p *productOp) GetProduct(id string, options any) (*Product, error) {
	// products-uuid path since akeneo 7
	basePath := productBasePath
	if p.client.osVersion >= AkeneoPimVersion7 {
		basePath = productUUIDBasePath
	}
	sourcePath := path.Join(basePath, id)
	product := new(Product)
	if err := p.client.GET(
		sourcePath,
		options,
		nil,
		product,
	); err != nil {
		return nil, err
	}
	return product, nil
}

func (p *productOp) UpdateOrCreateProducts(products []Product) (PatchProductResponse, error) {
	result := new(PatchProductResponse)
	if err := p.client.PATCH(
		productBasePath,
		nil,
		products,
		result,
	); err != nil {
		return nil, err
	}
	return *result, nil
}

// ProductsResponse is the struct for an akeneo products response
type ProductsResponse struct {
	Links       Links        `json:"_links,omitempty" mapstructure:"_links"`
	CurrentPage int          `json:"current_page,omitempty" mapstructure:"current_page"`
	Embedded    productItems `json:"_embedded,omitempty" mapstructure:"_embedded"`
}

type productItems struct {
	Items []Product `json:"items,omitempty" mapstructure:"items"`
}

type PatchProductResponseLine struct {
	Line       int    `json:"line,omitempty" mapstructure:"line"`
	Identifier string `json:"identifier,omitempty" mapstructure:"identifier"`
	Code       string `json:"code,omitempty" mapstructure:"code"`
	StatusCode int    `json:"status_code,omitempty" mapstructure:"status_code"`
	Message    string `json:"message,omitempty" mapstructure:"message"`
}
type PatchProductRequest []Product
type PatchProductResponse []PatchProductResponseLine
