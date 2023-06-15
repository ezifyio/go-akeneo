package goakeneo

// Connector is the struct to use to interact with the Akeneo API
type Connector struct {
	ClientID string `json:"client_id" mapstructure:"client_id"`
	Secret   string `json:"secret" mapstructure:"secret"`
	UserName string `json:"username" mapstructure:"username"`
	Password string `json:"password" mapstructure:"password"`
}

// NewClient creates a new Akeneo client
func (c Connector) NewClient(opts ...Option) (*Client, error) {
	return NewClient(c, opts...)
}
