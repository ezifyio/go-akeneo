package goakeneo

import (
	"path"
)

const (
	familyBasePath = "/api/rest/v1/families"
)

// FamilyService is the interface to interact with the Akeneo Family API
// todo: query parameters check
type FamilyService interface {
	ListWithPagination(options any) ([]Family, Links, error)
	GetFamily(familyCode string, options any) (*Family, error)
	GetFamilyVariants(familyCode string, options any) ([]FamilyVariant, error)
	GetFamilyVariant(familyCode string, familyVariantCode string) (*FamilyVariant, error)
	CreateFamily(family Family) error
	UpdateOrCreate(familyCode, familyVariantCode string, familyVariant FamilyVariant) error
}

type familyOp struct {
	client *Client
}

// ListWithPagination lists families with pagination
func (f *familyOp) ListWithPagination(options any) ([]Family, Links, error) {
	familyResponse := new(FamiliesResponse)
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
// do not use options for now
// get family does not support options yet, but it may in the future
func (f *familyOp) GetFamily(familyCode string, options any) (*Family, error) {
	sourcePath := path.Join(familyBasePath, familyCode)
	family := new(Family)
	if err := f.client.GET(
		sourcePath,
		options,
		nil,
		family,
	); err != nil {
		return nil, err
	}
	return family, nil
}

// GetFamilyVariants gets a family variants by code
func (f *familyOp) GetFamilyVariants(familyCode string, options any) ([]FamilyVariant, error) {
	sourcePath := path.Join(familyBasePath, familyCode, "variants")
	result := new(FamilyVariantsResponse)
	if err := f.client.GET(
		sourcePath,
		options,
		nil,
		result,
	); err != nil {
		return nil, err
	}
	return result.Embedded.Items, nil
}

// GetFamilyVariant gets a family variant by code
func (f *familyOp) GetFamilyVariant(familyCode string, familyVariantCode string) (*FamilyVariant, error) {
	sourcePath := path.Join(familyBasePath, familyCode, "variants", familyVariantCode)
	result := new(FamilyVariant)
	if err := f.client.GET(
		sourcePath,
		nil,
		nil,
		result,
	); err != nil {
		return nil, err
	}
	return result, nil
}

// CreateFamily creates a family
func (f *familyOp) CreateFamily(family Family) error {
	if err := f.client.POST(
		familyBasePath,
		nil,
		family,
		nil,
	); err != nil {
		return err
	}
	return nil
}

// UpdateOrCreate updates or creates a family variant
func (f *familyOp) UpdateOrCreate(familyCode, familyVariantCode string, familyVariant FamilyVariant) error {
	sourcePath := path.Join(familyBasePath, familyCode, "variants", familyVariantCode)
	if err := f.client.PATCH(
		sourcePath,
		nil,
		familyVariant,
		nil,
	); err != nil {
		return err
	}
	return nil
}

// FamiliesResponse is the struct for an akeneo families response
type FamiliesResponse struct {
	Links       Links       `json:"_links,omitempty" mapstructure:"_links"`
	CurrentPage int         `json:"current_page,omitempty" mapstructure:"current_page"`
	Embedded    familyItems `json:"_embedded,omitempty" mapstructure:"_embedded"`
}

type familyItems struct {
	Items []Family `json:"items,omitempty" mapstructure:"items"`
}

// FamilyVariantsResponse is the struct for an akeneo family variants response
type FamilyVariantsResponse struct {
	Links       Links              `json:"_links,omitempty" mapstructure:"_links"`
	CurrentPage int                `json:"current_page,omitempty" mapstructure:"current_page"`
	Embedded    familyVariantItems `json:"_embedded,omitempty" mapstructure:"_embedded"`
}

type familyVariantItems struct {
	Items []FamilyVariant `json:"items,omitempty" mapstructure:"items"`
}
