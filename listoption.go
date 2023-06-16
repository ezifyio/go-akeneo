package goakeneo

import (
	"encoding/json"
	"html"
)

// SearchFilter is a map of search filters,see :
// https://api.akeneo.com/documentation/filter.html
type SearchFilter map[string][]map[string]interface{}

func (sf SearchFilter) String() string {
	t, _ := json.Marshal(sf)
	return html.UnescapeString(string(t))
}

// ListOptions is the struct for common list options
type ListOptions struct {
	Search    string `url:"search,omitempty"`
	Page      int    `url:"page,omitempty"`
	Limit     int    `url:"limit,omitempty"`
	WithCount bool   `url:"with_count,omitempty"`
}

// ProductListOptions specifies the product optional parameters
// see: https://api.akeneo.com/api-reference.html#get_products
type ProductListOptions struct {
	Scope                string `url:"scope,omitempty"`
	Locales              string `url:"locales,omitempty"`
	Attributes           string `url:"attributes,omitempty"`
	PaginationType       string `url:"pagination_type,omitempty"`
	SearchAfter          string `url:"search_after,omitempty"`
	WithAttributeOptions bool   `url:"with_attribute_options,omitempty"`
	WithCompleteness     bool   `url:"with_completeness,omitempty"`
	WithQualityScores    bool   `url:"with_quality_scores,omitempty"`
	ListOptions
}

// ProductModelListOptions specifies the product model optional parameters
// see :https://api.akeneo.com/api-reference.html#Productmodel
type ProductModelListOptions struct {
	Scope             string `url:"scope,omitempty"`
	Locales           string `url:"locales,omitempty"`
	Attributes        string `url:"attributes,omitempty"`
	PaginationType    string `url:"pagination_type,omitempty"`
	SearchAfter       string `url:"search_after,omitempty"`
	WithQualityScores bool   `url:"with_quality_scores,omitempty"`
	ListOptions
}

// FamilyListOptions specifies the family optional parameters
// see :https://api.akeneo.com/api-reference.html#Family
type FamilyListOptions struct {
	ListOptions
}

// FamilyVariantListOptions specifies the family variant optional parameters
// see :https://api.akeneo.com/api-reference.html#FamilyVariant
type FamilyVariantListOptions struct {
	Page      int  `url:"page,omitempty"`
	Limit     int  `url:"limit,omitempty"`
	WithCount bool `url:"with_count,omitempty"`
}

// AttributeListOptions specifies the attribute optional parameters
type AttributeListOptions struct {
	WithTableSelectOptions bool `url:"with_table_select_options,omitempty" json:"with_table_select_options,omitempty" mapstructure:"with_table_select_options"` // false by default,decreases performance when enabled
	ListOptions
}

// AttributeOptionListOptions specifies the attribute option optional parameters
type AttributeOptionListOptions struct {
	Page      int  `url:"page,omitempty"`
	Limit     int  `url:"limit,omitempty"`
	WithCount bool `url:"with_count,omitempty"`
}

// AttributeGroupListOptions specifies the attribute group optional parameters
type AttributeGroupListOptions struct {
	ListOptions
}

// AssociationTypeListOptions specifies the association type optional parameters
type AssociationTypeListOptions struct {
	Page      int  `url:"page,omitempty"`
	Limit     int  `url:"limit,omitempty"`
	WithCount bool `url:"with_count,omitempty"`
}

// CategoryListOptions specifies the category optional parameters
type CategoryListOptions struct {
	ListOptions
	WithPosition           bool `url:"with_position,omitempty"`
	WithEnrichedAttributes bool `url:"with_enriched_attributes,omitempty"`
}
