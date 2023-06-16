package goakeneo

const (
	channelBasePath = "/api/rest/v1/channels"
)

// ChannelService is the interface to interact with the Akeneo Channel API
type ChannelService interface {
	ListWithPagination(options any) ([]Channel, Links, error)
}

type channelOp struct {
	client *Client
}

// ListWithPagination lists channels with pagination
// options should be url.Values
func (c *channelOp) ListWithPagination(options any) ([]Channel, Links, error) {
	channelResponse := new(ChannelsResponse)
	if err := c.client.GET(
		channelBasePath,
		options,
		nil,
		channelResponse,
	); err != nil {
		return nil, Links{}, err
	}
	return channelResponse.Embedded.Items, channelResponse.Links, nil
}

// ChannelsResponse is the struct for an akeneo channels response
type ChannelsResponse struct {
	Links       Links        `json:"_links" mapstructure:"_links"`
	CurrentPage int          `json:"current_page" mapstructure:"current_page"`
	Embedded    channelItems `json:"_embedded" mapstructure:"_embedded"`
}

type channelItems struct {
	Items []Channel `json:"items" mapstructure:"items"`
}
