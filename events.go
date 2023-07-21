package goakeneo

const (
	EventTypeProductCreated      = "product.created"
	EventTypeProductUpdated      = "product.updated"
	EventTypeProductRemoved      = "product.removed"
	EventTypeProductModelCreated = "product_model.created"
	EventTypeProductModelUpdated = "product_model.updated"
	EventTypeProductModelRemoved = "product_model.removed"
)

// Event akeneo events dispatch as json
// see: https://api.akeneo.com/events-documentation/subscription.html
type Event struct {
	Action        string       `json:"action,omitempty" mapstructure:"action"`
	Author        string       `json:"author,omitempty" mapstructure:"author"`
	AuthorType    string       `json:"author_type,omitempty" mapstructure:"author_type"`
	EventID       string       `json:"event_id,omitempty" mapstructure:"event_id"`
	EventDatatime string       `json:"event_datetime,omitempty" mapstructure:"event_datetime"`
	PimSource     string       `json:"pim_source,omitempty" mapstructure:"pim_source"`
	Data          dataResource `json:"data,omitempty" mapstructure:"data"`
}

// EventType return the event type
func (e *Event) EventType() string {
	return e.Action
}

// ID return the resource id
func (e *Event) ID() string {
	switch e.Action {
	case EventTypeProductCreated, EventTypeProductUpdated, EventTypeProductRemoved:
		return e.Data.Resource.Identifier
	case EventTypeProductModelCreated, EventTypeProductModelUpdated, EventTypeProductModelRemoved:
		return e.Data.Resource.Code
	default:
		return ""
	}
}

type dataResource struct {
	Resource resource `json:"resource,omitempty" mapstructure:"resource"`
}

type resource struct {
	UUID                   string                           `json:"uuid,omitempty" mapstructure:"uuid"`
	Code                   string                           `json:"code,omitempty" mapstructure:"code"`
	Identifier             string                           `json:"identifier,omitempty" mapstructure:"identifier"`
	Enabled                bool                             `json:"enabled,omitempty" mapstructure:"enabled"`
	Family                 string                           `json:"family,omitempty" mapstructure:"family"`
	FamilyVariant          string                           `json:"family_variant,omitempty" mapstructure:"family_variant"`
	Categories             []string                         `json:"categories,omitempty" mapstructure:"categories"`
	Groups                 []string                         `json:"groups,omitempty" mapstructure:"groups"`
	Parent                 string                           `json:"parent,omitempty" mapstructure:"parent"`
	Values                 []ProductValue                   `json:"values,omitempty" mapstructure:"values"`
	QuantifiedAssociations map[string]quantifiedAssociation `json:"quantified_associations,omitempty" mapstructure:"quantified_associations"`
	Associations           map[string]association           `json:"associations,omitempty" mapstructure:"associations"`
	Created                string                           `json:"created,omitempty" mapstructure:"created"`
	Updated                string                           `json:"updated,omitempty" mapstructure:"updated"`
}
