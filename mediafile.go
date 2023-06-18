package goakeneo

import (
	"bytes"
	"encoding/json"
	"github.com/pkg/errors"
	"io"
	"mime/multipart"
	"net/url"
	"os"
	"path"
)

const mediaBasePath = "/api/rest/v1/media-files"

// MediaFileService see: https://api.akeneo.com/api-reference.html#media-files
type MediaFileService interface {
	ListPagination(options any) ([]MediaFile, Links, error)
	GetByCode(code string, options any) (*MediaFile, error)
	Download(code, filePath string, options any) error
	Create(filePath string, association MediaFileAssociation) (string, error)
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
	options = nil // options are not supported for downloading media files yet
	sourcePath := path.Join(mediaBasePath, code, "download")
	sourceP, _ := url.Parse(sourcePath)
	downloadURL := c.client.baseURL.ResolveReference(sourceP).String()
	if err := c.client.download(downloadURL, filePath); err != nil {
		return err
	}
	return nil
}

// Create creates a media file
func (c *mediaOp) Create(filePath string, association MediaFileAssociation) (string, error) {
	// check if file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return "", errors.Wrapf(err, "file %s does not exist", filePath)
	}
	f, err := os.Open(filePath)
	if err != nil {
		return "", errors.Wrapf(err, "failed to open file %s", filePath)
	}
	defer f.Close()
	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)
	fileWriter, err := writer.CreateFormFile("file", path.Base(filePath))
	if err != nil {
		return "", errors.Wrapf(err, "failed to create form file %s", filePath)
	}
	// copy image file to form file writer
	if _, err = io.Copy(fileWriter, f); err != nil {
		return "", errors.Wrapf(err, "failed to copy file %s", filePath)
	}
	// add association
	if association.Type() == "product" {
		if err = writer.WriteField("product", association.ToJSONString()); err != nil {
			return "", errors.Wrapf(err, "failed to write field product %s", filePath)
		}
	} else if association.Type() == "product_model" {
		if err = writer.WriteField("product_model", association.ToJSONString()); err != nil {
			return "", errors.Wrapf(err, "failed to write field product_model %s", filePath)
		}
	}
	// close writer
	if err = writer.Close(); err != nil {
		return "", errors.Wrapf(err, "failed to close writer %s", filePath)
	}
	uri, err := c.client.upload(mediaBasePath, writer)
	if err != nil {
		return "", errors.Wrapf(err, "failed to upload file %s", filePath)
	}
	return uri, nil
}

type MediaFileAssociation interface {
	ToJSONString() string
	Type() string
}
type AssociatedProduct struct {
	Identifier string `json:"identifier,omitempty" mapstructure:"identifier"`
	Attribute  string `json:"attribute,omitempty" mapstructure:"attribute"`
	Scope      string `json:"scope,omitempty" mapstructure:"scope"`
	Locale     string `json:"locale,omitempty" mapstructure:"locale"`
}

func (p AssociatedProduct) ToJSONString() string {
	j, _ := json.Marshal(p)
	return string(j)
}

func (p AssociatedProduct) Type() string {
	return "product"
}

type AssociatedProductModel struct {
	Code      string `json:"code,omitempty" mapstructure:"code"`
	Attribute string `json:"attribute,omitempty" mapstructure:"attribute"`
	Scope     string `json:"scope,omitempty" mapstructure:"scope"`
	Locale    string `json:"locale,omitempty" mapstructure:"locale"`
}

func (p AssociatedProductModel) ToJSONString() string {
	j, _ := json.Marshal(p)
	return string(j)
}

func (p AssociatedProductModel) Type() string {
	return "product_model"
}

type MediaFileResponse struct {
	Links       Links      `json:"_links,omitempty" mapstructure:"_links"`
	CurrentPage int        `json:"current_page,omitempty" mapstructure:"current_page"`
	Embedded    mediaItems `json:"_embedded,omitempty" mapstructure:"_embedded"`
}

type mediaItems struct {
	Items []MediaFile `json:"items,omitempty" mapstructure:"items"`
}
