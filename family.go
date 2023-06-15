package goakeneo

import "fmt"

const (
	familyBasePath = "/api/rest/v1/families"
)

// FamilyService is the interface to interact with the Akeneo Family API
type FamilyService interface {
	ListWithPagination(options any) ([]Family, Links, error)
	GetFamily(familyCode string) (*Family, error)
	GetFamilyVariants(familyCode string) ([]FamilyVariant, error)
	GetFamilyVariant(familyCode string, familyVariantCode string) (*FamilyVariant, error)
}

type familyOp struct {
	client *Client
}

// ListWithPagination lists families with pagination
// options should be
func (f *familyOp) ListWithPagination(options any) ([]Family, Links, error) {
	familyResponse := new(familiesResponse)
	if err := f.client.GET(
		familyBasePath,
		options,
		nil,
		familyResponse,
	); err != nil {
		return nil, Links{}, err
	}
	return familyResponse.Embedded.Items, familyResponse.Links, nil
}

// GetFamily gets a family by code
func (f *familyOp) GetFamily(familyCode string) (*Family, error) {
	path := fmt.Sprintf("%s/%s", familyBasePath, familyCode)
	family := new(Family)
	if err := f.client.GET(
		path,
		nil,
		nil,
		family,
	); err != nil {
		return nil, err
	}
	return family, nil
}

// GetFamilyVariants gets a family variants by code
func (f *familyOp) GetFamilyVariants(familyCode string) ([]FamilyVariant, error) {
	path := fmt.Sprintf("%s/%s/variants", familyBasePath, familyCode)
	result := new(familyVariantsResponse)
	if err := f.client.GET(
		path,
		nil,
		nil,
		result,
	); err != nil {
		return nil, err
	}
	return result.Embedded.Items, nil
}

// GetFamilyVariant gets a family variant by code
func (f *familyOp) GetFamilyVariant(familyCode string, familyVariantCode string) (*FamilyVariant, error) {
	path := fmt.Sprintf("%s/%s/variants/%s", familyBasePath, familyCode, familyVariantCode)
	result := new(FamilyVariant)
	if err := f.client.GET(
		path,
		nil,
		nil,
		result,
	); err != nil {
		return nil, err
	}
	return result, nil
}

// familiesResponse is the struct for a akeneo families response
type familiesResponse struct {
	Links       Links       `json:"_links" mapstructure:"_links"`
	CurrentPage int         `json:"current_page" mapstructure:"current_page"`
	Embedded    familyItems `json:"_embedded" mapstructure:"_embedded"`
}

type familyItems struct {
	Items []Family `json:"items" mapstructure:"items"`
}

// familyVariantsResponse is the struct for a akeneo family variants response
type familyVariantsResponse struct {
	Links       Links              `json:"_links" mapstructure:"_links"`
	CurrentPage int                `json:"current_page" mapstructure:"current_page"`
	Embedded    familyVariantItems `json:"_embedded" mapstructure:"_embedded"`
}

type familyVariantItems struct {
	Items []FamilyVariant `json:"items" mapstructure:"items"`
}
