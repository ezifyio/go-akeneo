package goakeneo

// Event akeneo events dispatch as json
// see: https://api.akeneo.com/events-documentation/subscription.html
type Event struct {
	Action        string   `json:"action,omitempty" mapstructure:"action"`
	Author        string   `json:"author,omitempty" mapstructure:"author"`
	AuthorType    string   `json:"author_type,omitempty" mapstructure:"author_type"`
	EventID       string   `json:"event_id,omitempty" mapstructure:"event_id"`
	EventDatatime string   `json:"event_datetime,omitempty" mapstructure:"event_datetime"`
	PimSource     string   `json:"pim_source,omitempty" mapstructure:"pim_source"`
	Data          resource `json:"data,omitempty" mapstructure:"data"`
}

type resource struct {
	Resource map[string]any `json:"resource,omitempty" mapstructure:"resource"`
}
