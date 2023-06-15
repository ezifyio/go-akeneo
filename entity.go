package goakeneo

import (
	"net/url"
	"strconv"
)

// ValueTypeConst
const (
	ValueTypeString           = "string"
	ValueTypeStringCollection = "string_collection"
	ValueTypeNumber           = "number"
	ValueTypeMetric           = "metric"
	ValueTypePrice            = "price"
	ValueTypeBoolean          = "boolean"
	ValueTypeSimpleSelect     = "simple_select"
	ValueTypeMultiSelect      = "multi_select"
	ValueTypeTable            = "table"
)

// Product is the struct for a akeneo product
type Product struct {
	Links                  Links                            `json:"_links,omitempty" mapstructure:"_links"`
	UUID                   string                           `json:"uuid,omitempty" mapstructure:"uuid"` // Since Akeneo 7.0
	Identifier             string                           `json:"identifier,omitempty" mapstructure:"identifier"`
	Enabled                bool                             `json:"enabled,omitempty" mapstructure:"enabled"`
	Family                 string                           `json:"family,omitempty" mapstructure:"family"`
	Categories             []string                         `json:"categories,omitempty" mapstructure:"categories"`
	Groups                 []string                         `json:"groups,omitempty" mapstructure:"groups"`
	Parent                 string                           `json:"parent,omitempty" mapstructure:"parent"` // code of the parent product model when the product is a variant
	Values                 map[string][]ProductValue        `json:"values,omitempty" mapstructure:"values"`
	Associations           map[string]association           `json:"associations,omitempty" mapstructure:"associations"`
	QuantifiedAssociations map[string]quantifiedAssociation `json:"quantified_associations,omitempty" mapstructure:"quantified_associations"` // Since Akeneo 5.0
	Created                string                           `json:"created,omitempty" mapstructure:"created"`
	Updated                string                           `json:"updated,omitempty" mapstructure:"updated"`
	QualityScores          []QualityScore                   `json:"quality_scores,omitempty" mapstructure:"quality_scores"` // Since Akeneo 5.0,WithQualityScores must be true in the request
	Completenesses         []any                            `json:"completenesses,omitempty" mapstructure:"completenesses"` // Since Akeneo 6.0,WithCompleteness must be true in the request
	Metadata               map[string]string                `json:"metadata,omitempty" mapstructure:"metadata"`             // Enterprise Edition only
}

// Links is the struct for akeneo links
type Links struct {
	Self     Link `json:"self,omitempty"`
	First    Link `json:"first,omitempty"`
	Previous Link `json:"previous,omitempty"`
	Next     Link `json:"next,omitempty"`
	Download Link `json:"download,omitempty"`
}

// HasNext returns true if there is a next link
func (l Links) HasNext() bool {
	return l.Next.Href != ""
}

// NextOptions returns the options for the next link
func (l Links) NextOptions() url.Values {
	u, err := url.Parse(l.Next.Href)
	if err != nil {
		return nil
	}
	return u.Query()
}

// Link is the struct for a akeneo link
type Link struct {
	Href string `json:"href,omitempty"`
}

// ProductValue is the interface for a akeneo product value
// see: https://api.akeneo.com/concepts/products.html#the-data-format
type ProductValue interface {
	ValueType() string
}

// StringValue is the struct for a akeneo text type product value
// pim_catalog_text or pim_catalog_textarea : data is a string
// pim_catalog_file or pim_catalog_image: data is the file path
// pim_catalog_date : data is a string in ISO-8601 format
type StringValue struct {
	Locale string `json:"locale,omitempty" mapstructure:"locale"`
	Scope  string `json:"scope,omitempty" mapstructure:"scope"`
	Data   string `json:"data,omitempty" mapstructure:"data"`
}

// ValueType returns the value type, see ValueTypeConst
func (StringValue) ValueType() string {
	return ValueTypeString
}

// StringCollectionValue is the struct for a akeneo collection type product value
type StringCollectionValue struct {
	Locale string   `json:"locale,omitempty" mapstructure:"locale"`
	Scope  string   `json:"scope,omitempty" mapstructure:"scope"`
	Data   []string `json:"data,omitempty" mapstructure:"data"`
}

// ValueType returns the value type, see ValueTypeConst
func (StringCollectionValue) ValueType() string {
	return ValueTypeStringCollection
}

// NumberValue is the struct for a akeneo number type product value
// pim_catalog_number : data is a int when decimal is false ,float64 string when decimal is true
// so the data will be parsed as ValueTypeString when decimal is true
type NumberValue struct {
	Locale string `json:"locale,omitempty" mapstructure:"locale"`
	Scope  string `json:"scope,omitempty" mapstructure:"scope"`
	Data   int    `json:"data,omitempty" mapstructure:"data"`
}

// ValueType returns the value type, see ValueTypeConst
func (NumberValue) ValueType() string {
	return ValueTypeNumber
}

// MetricValue is the struct for a akeneo metric type product value
// pim_catalog_metric : data amount is a float64 string when decimal is true, int when decimal is false
type MetricValue struct {
	Locale string `json:"locale,omitempty" mapstructure:"locale"`
	Scope  string `json:"scope,omitempty" mapstructure:"scope"`
	Data   metric `json:"data,omitempty" mapstructure:"data"`
}

type metric struct {
	Amount any    `json:"amount,omitempty" mapstructure:"amount"`
	Unit   string `json:"unit,omitempty" mapstructure:"unit"`
}

// ValueType returns the value type, see ValueTypeConst
func (MetricValue) ValueType() string {
	return ValueTypeMetric
}

// Amount returns the amount as string
func (v MetricValue) Amount() string {
	if f, ok := v.Data.Amount.(string); ok {
		return f
	}
	i, ok := v.Data.Amount.(int)
	if !ok {
		return ""
	}
	return strconv.Itoa(i)
}

// PriceValue is the struct for a akeneo price type product value
// pim_catalog_price : data amount is a float64 string when decimal is true, int when decimal is false
type PriceValue struct {
	Locale string `json:"locale,omitempty" mapstructure:"locale"`
	Scope  string `json:"scope,omitempty" mapstructure:"scope"`
	Data   price  `json:"data,omitempty" mapstructure:"data"`
}

type price struct {
	Amount   any    `json:"amount,omitempty" mapstructure:"amount"`
	Currency string `json:"currency,omitempty" mapstructure:"currency"`
}

// ValueType returns the value type, see ValueTypeConst
func (PriceValue) ValueType() string {
	return ValueTypePrice
}

// Amount returns the amount as string
func (v PriceValue) Amount() string {
	if f, ok := v.Data.Amount.(string); ok {
		return f
	}
	i, ok := v.Data.Amount.(int)
	if !ok {
		return ""
	}
	return strconv.Itoa(i)
}

// BooleanValue is the struct for a akeneo boolean type product value
// pim_catalog_boolean : data is a bool
type BooleanValue struct {
	Locale string `json:"locale,omitempty" mapstructure:"locale"`
	Scope  string `json:"scope,omitempty" mapstructure:"scope"`
	Data   bool   `json:"data,omitempty" mapstructure:"data"`
}

// ValueType returns the value type, see ValueTypeConst
func (BooleanValue) ValueType() string {
	return ValueTypeBoolean
}

type linkedData struct {
	Attribute string            `json:"attribute,omitempty" mapstructure:"attribute"`
	Code      string            `json:"code,omitempty" mapstructure:"code"`
	Labels    map[string]string `json:"labels,omitempty" mapstructure:"labels"`
}

// SimpleSelectValue is the struct for a akeneo simple select type product value
type SimpleSelectValue struct {
	Locale     string     `json:"locale,omitempty" mapstructure:"locale"`
	Scope      string     `json:"scope,omitempty" mapstructure:"scope"`
	Data       string     `json:"data,omitempty" mapstructure:"data"`
	LinkedData linkedData `json:"linked_data,omitempty" mapstructure:"linked_data"`
}

// ValueType returns the value type, see ValueTypeConst
func (SimpleSelectValue) ValueType() string {
	return ValueTypeSimpleSelect
}

// MultiSelectValue is the struct for a akeneo multi select type product value
type MultiSelectValue struct {
	Locale     string                `json:"locale,omitempty" mapstructure:"locale"`
	Scope      string                `json:"scope,omitempty" mapstructure:"scope"`
	Data       []string              `json:"data,omitempty" mapstructure:"data"`
	LinkedData map[string]linkedData `json:"linked_data,omitempty" mapstructure:"linked_data"`
}

// ValueType returns the value type, see ValueTypeConst
func (MultiSelectValue) ValueType() string {
	return ValueTypeMultiSelect
}

// TableValue is the struct for a akeneo table type product value
// pim_catalog_table : data is a []map[string]any
type TableValue struct {
	Locale string `json:"locale,omitempty" mapstructure:"locale"`
	Scope  string `json:"scope,omitempty" mapstructure:"scope"`
	Data   []map[string]any
}

// ValueType returns the value type, see ValueTypeConst
func (TableValue) ValueType() string {
	return ValueTypeTable
}

type association struct {
	Groups        []string `json:"groups" mapstructure:"groups"`
	Products      []string `json:"products" mapstructure:"products"`
	ProductModels []string `json:"product_models" mapstructure:"product_models"`
}

// QuantifiedAssociations is the struct for a akeneo quantified associations
type quantifiedAssociation struct {
	Products      []productQuantity      `json:"products" mapstructure:"products"`
	ProductModels []productModelQuantity `json:"product_models" mapstructure:"product_models"`
}

type productQuantity struct {
	Identifier string `json:"identifier" mapstructure:"identifier"`
	Quantity   int    `json:"quantity" mapstructure:"quantity"`
}

type productModelQuantity struct {
	Code     string `json:"code" mapstructure:"code"`
	Quantity int    `json:"quantity" mapstructure:"quantity"`
}

// QualityScore is the struct for quality score
type QualityScore struct {
	Scope  string `json:"scope,omitempty" validate:"required"`
	Locale string `json:"locale,omitempty" validate:"required"`
	Data   string `json:"data,omitempty" validate:"required"`
}

// Family is the struct for a akeneo family
type Family struct {
	Links                 Links               `json:"_links,omitempty" mapstructure:"_links"`
	Code                  string              `json:"code,omitempty" mapstructure:"code"`                                     // The code of the family
	AttributeAsLabel      string              `json:"attribute_as_label,omitempty" mapstructure:"attribute_as_label"`         // The code of the attribute used as label for the family
	AttributeAsImage      string              `json:"attribute_as_image,omitempty" mapstructure:"attribute_as_image"`         // Attribute code used as the main picture in the user interface (only since v2.fmt
	Attributes            []string            `json:"attributes,omitempty" mapstructure:"attributes"`                         //  Attributes codes that compose the family
	AttributeRequirements map[string][]string `json:"attribute_requirements,omitempty" mapstructure:"attribute_requirements"` //  • Attributes codes of the family that are required for the completeness calculation for the channel `channelCode`
	Labels                map[string]string   `json:"labels,omitempty" mapstructure:"labels"`                                 //  Translatable labels. Ex: {"en_US": "T-shirt", "fr_FR": "T-shirt"}
}

// FamilyVariant is the struct for a akeneo family variant
type FamilyVariant struct {
	Links                Links                 `json:"_links,omitempty" mapstructure:"_links"`
	Code                 string                `json:"code,omitempty" mapstructure:"code"`                                     // The code of the family variant
	Lables               map[string]string     `json:"labels,omitempty" mapstructure:"labels"`                                 // Translatable labels. Ex: {"en_US": "T-shirt", "fr_FR": "T-shirt"}
	VariantAttributeSets []variantAttributeSet `json:"variant_attribute_sets,omitempty" mapstructure:"variant_attribute_sets"` // The variant attribute sets of the family variant
}

type variantAttributeSet struct {
	Level      int      `json:"level,omitempty" mapstructure:"level"`           // The level of the variant attribute set
	Axes       []string `json:"axes,omitempty" mapstructure:"axes"`             // The axes of the variant attribute set
	Attributes []string `json:"attributes,omitempty" mapstructure:"attributes"` // The attributes of the variant attribute set
}
