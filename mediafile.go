package goakeneo

const mediaBasePath = "/api/rest/v1/media-files"

// MediaFileService see: https://api.akeneo.com/api-reference.html#media-files
type MediaFileService interface {
	ListPagination(options any) ([]MediaFile, Links, error)
}

type mediaOp struct {
	client *Client
}

// ListPagination lists media files with pagination
func (c *mediaOp) ListPagination(options any) ([]MediaFile, Links, error) {
	mediaResponse := new(MediaFileResponse)
	if err := c.client.GET(
		mediaBasePath,
		options,
		nil,
		mediaResponse,
	); err != nil {
		return nil, Links{}, err
	}
	return mediaResponse.Embedded.Items, mediaResponse.Links, nil
}

type MediaFileResponse struct {
	Links       Links      `json:"_links,omitempty" mapstructure:"_links"`
	CurrentPage int        `json:"current_page,omitempty" mapstructure:"current_page"`
	Embedded    mediaItems `json:"_embedded,omitempty" mapstructure:"_embedded"`
}

type mediaItems struct {
	Items []MediaFile `json:"items,omitempty" mapstructure:"items"`
}
