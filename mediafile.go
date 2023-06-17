package goakeneo

import (
	"os"
	"path"
)

const mediaBasePath = "/api/rest/v1/media-files"

// MediaFileService see: https://api.akeneo.com/api-reference.html#media-files
type MediaFileService interface {
	ListPagination(options any) ([]MediaFile, Links, error)
	GetByCode(code string, options any) (*MediaFile, error)
	Download(code string, options any) ([]byte, error)
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

// GetByCode gets a media file by code
func (c *mediaOp) GetByCode(code string, options any) (*MediaFile, error) {
	result := new(MediaFile)
	sourcePath := path.Join(mediaBasePath, code)
	if err := c.client.GET(
		sourcePath,
		options,
		nil,
		result,
	); err != nil {
		return nil, err
	}
	return result, nil
}

// Download downloads a media file by code
func (c *mediaOp) Download(code, filePath string, options any) error {
	options = nil // options are not supported for downloading media files
	result := new(MediaFile)
	sourcePath := path.Join(mediaBasePath, code)
	if err := c.client.GET(
		sourcePath,
		options,
		nil,
		result,
	); err != nil {
		return err
	}
	downloadURL := result.Links.Download.Href
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	if err := c.client.download(downloadURL, file); err != nil {
		return err
	}
	return nil
}

type MediaFileResponse struct {
	Links       Links      `json:"_links,omitempty" mapstructure:"_links"`
	CurrentPage int        `json:"current_page,omitempty" mapstructure:"current_page"`
	Embedded    mediaItems `json:"_embedded,omitempty" mapstructure:"_embedded"`
}

type mediaItems struct {
	Items []MediaFile `json:"items,omitempty" mapstructure:"items"`
}
