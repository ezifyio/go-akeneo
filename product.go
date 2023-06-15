package goakeneo

import (
	"path"

	"github.com/pkg/errors"
)

const (
	productBasePath = "/api/rest/v1/products"
	// todo: product uuid path
)

// ProductService is the interface to interact with the Akeneo Product API
type ProductService interface {
	GetAllProducts(options any) (<-chan Product, chan error)
	ListWithPagination(options any) ([]Product, Links, error)
	GetProduct(id string, options any) (*Product, error)
}

type productOp struct {
	client *Client
}

// GetAllProducts lists all products, returns a channel to iterate over products
func (p *productOp) GetAllProducts(options any) (<-chan Product, chan error) {
	prodChan := make(chan Product, 1)
	errChan := make(chan error, 1)
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
		if err != nil {
			errChan <- err
			return
		}
		for {
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
	productResponse := new(productsResponse)
	if err := p.client.GET(
		productBasePath,
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
	sourcePath := path.Join(productBasePath, id)
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

// productsResponse is the struct for a akeneo products response
type productsResponse struct {
	Links       Links        `json:"_links" mapstructure:"_links"`
	CurrentPage int          `json:"current_page" mapstructure:"current_page"`
	Embedded    productItems `json:"_embedded" mapstructure:"_embedded"`
}

type productItems struct {
	Items []Product `json:"items" mapstructure:"items"`
}
