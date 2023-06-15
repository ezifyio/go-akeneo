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

// Product is the struct for an akeneo product
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

// Link is the struct for an akeneo link
type Link struct {
	Href string `json:"href,omitempty"`
}

// ProductValue is the interface for an akeneo product value
// see: https://api.akeneo.com/concepts/products.html#the-data-format
type ProductValue interface {
	ValueType() string
}

// StringValue is the struct for an akeneo text type product value
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

// StringCollectionValue is the struct for an akeneo collection type product value
type StringCollectionValue struct {
	Locale string   `json:"locale,omitempty" mapstructure:"locale"`
	Scope  string   `json:"scope,omitempty" mapstructure:"scope"`
	Data   []string `json:"data,omitempty" mapstructure:"data"`
}

// ValueType returns the value type, see ValueTypeConst
func (StringCollectionValue) ValueType() string {
	return ValueTypeStringCollection
}

// NumberValue is the struct for an akeneo number type product value
// pim_catalog_number : data is an int when decimal is false ,float64 string when decimal is true
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

// MetricValue is the struct for an akeneo metric type product value
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

// PriceValue is the struct for an akeneo price type product value
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

// BooleanValue is the struct for an akeneo boolean type product value
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

// SimpleSelectValue is the struct for an akeneo simple select type product value
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

// MultiSelectValue is the struct for an akeneo multi select type product value
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

// TableValue is the struct for an akeneo table type product value
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

// QuantifiedAssociations is the struct for an akeneo quantified associations
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

// Family is the struct for an akeneo family
type Family struct {
	Links                 Links               `json:"_links,omitempty" mapstructure:"_links"`
	Code                  string              `json:"code,omitempty" mapstructure:"code"`                                     // The code of the family
	AttributeAsLabel      string              `json:"attribute_as_label,omitempty" mapstructure:"attribute_as_label"`         // The code of the attribute used as label for the family
	AttributeAsImage      string              `json:"attribute_as_image,omitempty" mapstructure:"attribute_as_image"`         // Attribute code used as the main picture in the user interface (only since v2.fmt
	Attributes            []string            `json:"attributes,omitempty" mapstructure:"attributes"`                         //  Attributes codes that compose the family
	AttributeRequirements map[string][]string `json:"attribute_requirements,omitempty" mapstructure:"attribute_requirements"` //  • Attributes codes of the family that are required for the completeness calculation for the channel `channelCode`
	Labels                map[string]string   `json:"labels,omitempty" mapstructure:"labels"`                                 //  Translatable labels. Ex: {"en_US": "T-shirt", "fr_FR": "T-shirt"}
}

// FamilyVariant is the struct for an akeneo family variant
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

// Attribute is the struct for an akeneo attribute,see:
// https://api.akeneo.com/api-reference.html#Attribute
type Attribute struct {
	Links               Links             `json:"_links" mapstructure:"_links"`
	Code                string            `json:"code" mapstructure:"code"`
	Type                string            `json:"type" mapstructure:"type"`
	Labels              map[string]string `json:"labels" mapstructure:"labels"`
	Group               string            `json:"group" mapstructure:"group"`
	GroupLabels         map[string]string `json:"group_labels" mapstructure:"group_labels"`
	SortOrder           int               `json:"sort_order" mapstructure:"sort_order"`
	Localizable         bool              `json:"localizable" mapstructure:"localizable"`                       // whether the attribute is localizable or not,i.e. whether it can be translated or not
	Scopable            bool              `json:"scopable" mapstructure:"scopable"`                             // whether the attribute is scopable or not,i.e. whether it can have different values depending on the channel or not
	AvailableLocales    []string          `json:"available_locales" mapstructure:"available_locales"`           // the list of activated locales for the attribute values
	Unique              bool              `json:"unique" mapstructure:"unique"`                                 // whether the attribute value is unique or not
	UseableAsGridFilter bool              `json:"useable_as_grid_filter" mapstructure:"useable_as_grid_filter"` // whether the attribute can be used as a filter in the product grid or not
	MaxCharacters       int               `json:"max_characters" mapstructure:"max_characters"`                 // the maximum number of characters allowed for the value of the attribute
	ValidationRule      string            `json:"validation_rule" mapstructure:"validation_rule"`               // validation rule code to validate the attribute value
	ValidationRegexp    string            `json:"validation_regexp" mapstructure:"validation_regexp"`           // validation regexp to validate the attribute value
	WysiwygEnabled      bool              `json:"wysiwyg_enabled" mapstructure:"wysiwyg_enabled"`               // whether the attribute can have a value per channel or not
	NumberMin           string            `json:"number_min" mapstructure:"number_min"`                         // the minimum value allowed for the value of the attribute
	NumberMax           string            `json:"number_max" mapstructure:"number_max"`                         // the maximum value allowed for the value of the attribute
	DecimalsAllowed     bool              `json:"decimals_allowed" mapstructure:"decimals_allowed"`             // whether decimals are allowed for the attribute or not
	NegativeAllowed     bool              `json:"negative_allowed" mapstructure:"negative_allowed"`             // whether negative numbers are allowed for the attribute or not
	MetricFamily        string            `json:"metric_family" mapstructure:"metric_family"`                   // the metric family of the attribute
	DefaultMetricUnit   string            `json:"default_metric_unit" mapstructure:"default_metric_unit"`       // the default metric unit of the attribute
	DateMin             string            `json:"date_min" mapstructure:"date_min"`                             // the minimum date allowed for the value of the attribute
	DateMax             string            `json:"date_max" mapstructure:"date_max"`                             // the maximum date allowed for the value of the attribute
	AllowedExtensions   []string          `json:"allowed_extensions" mapstructure:"allowed_extensions"`         // the list of allowed extensions for the value of the attribute
	MaxFileSize         string            `json:"max_file_size" mapstructure:"max_file_size"`                   // the maximum file size allowed for the value of the attribute
	ReferenceDataName   string            `json:"reference_data_name" mapstructure:"reference_data_name"`       // the reference data name of the attribute
	DefaultValue        bool              `json:"default_value" mapstructure:"default_value"`                   // the default value of the attribute
	TableConfiguration  []string          `json:"table_configuration" mapstructure:"table_configuration"`       // the table configuration of the attribute
}

// AttributeOption is the struct for an akeneo attribute option,see:
type AttributeOption struct {
	Links     Links             `json:"_links" mapstructure:"_links"`
	Code      string            `json:"code" mapstructure:"code"`
	Attribute string            `json:"attribute" mapstructure:"attribute"`
	SortOrder int               `json:"sort_order" mapstructure:"sort_order"`
	Labels    map[string]string `json:"labels" mapstructure:"labels"`
}

// Category is the struct for an akeneo category
type Category struct {
	Links    Links                    `json:"_links" mapstructure:"_links"`
	Code     string                   `json:"code" mapstructure:"code"`
	Parent   string                   `json:"parent" mapstructure:"parent"`
	Updated  string                   `json:"updated" mapstructure:"updated"`
	Position int                      `json:"position" mapstructure:"position"` // since 7.0 with query parameter "with_positions=true"
	Labels   map[string]string        `json:"labels" mapstructure:"labels"`
	Values   map[string]categoryValue `json:"values" mapstructure:"values"`
}

// categoryValue is the struct for an akeneo category value
// todo : Data field is not yet implemented well
type categoryValue struct {
	Data          any    `json:"data" mapstructure:"data"`           //  AttributeValue
	Type          string `json:"type" mapstructure:"type"`           //  AttributeType
	Locale        string `json:"locale" mapstructure:"locale"`       //  AttributeLocale
	Channel       string `json:"channel" mapstructure:"channel"`     //  AttributeChannel
	AttributeCode string `json:"attribute" mapstructure:"attribute"` //  AttributeCode with uuid, i.e. "description|96b88bf4-c2b7-4b64-a1f9-5d4876c02c26"
}

// Channel is the struct for an akeneo channel
type Channel struct {
	Links           Links             `json:"_links" mapstructure:"_links"`
	Code            string            `json:"code" mapstructure:"code"`
	Currencies      []string          `json:"currencies" mapstructure:"currencies"`
	Locales         []string          `json:"locales" mapstructure:"locales"`
	CategoryTree    string            `json:"category_tree" mapstructure:"category_tree"`
	ConversionUnits map[string]string `json:"conversion_units" mapstructure:"conversion_units"`
	Labels          map[string]string `json:"labels" mapstructure:"labels"`
}

// Locale is the struct for an akeneo locale
type Locale struct {
	Links   Links  `json:"_links" mapstructure:"_links"`
	Code    string `json:"code" mapstructure:"code"`
	Enabled bool   `json:"enabled" mapstructure:"enabled"`
}

// MediaFile is the struct for an akeneo media file
type MediaFile struct {
	Links            Links  `json:"_links,omitempty" mapstructure:"_links"`
	Code             string `json:"code,omitempty" mapstructure:"code"`
	OriginalFilename string `json:"original_filename,omitempty" mapstructure:"original_filename"`
	MimeType         string `json:"mime_type,omitempty" mapstructure:"mime_type"`
	Size             int    `json:"size,omitempty" mapstructure:"size"`
	Extension        string `json:"extension,omitempty" mapstructure:"extension"`
}
