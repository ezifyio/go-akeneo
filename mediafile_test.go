package goakeneo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMediaOp_ListPagination(t *testing.T) {
	c := MockClient()
	ms, links, err := c.MediaFile.ListPagination(nil)
	assert.NoError(t, err)
	assert.NotNil(t, ms)
	assert.NotNil(t, links)
}

func TestMediaOp_GetByCode(t *testing.T) {
	c := MockClient()
	m, err := c.MediaFile.GetByCode("0/2/5/2/025271bc214518f255d543282237df62cf97c96d_Copy_of_Green_Spreadsheet_Data_Analysis_Instagram_Post__1600____900_px_.png", nil)
	assert.NoError(t, err)
	assert.NotNil(t, m)
}

func TestMediaOp_Download(t *testing.T) {
	c := MockClient()
	err := c.MediaFile.Download("0/2/5/2/025271bc214518f255d543282237df62cf97c96d_Copy_of_Green_Spreadsheet_Data_Analysis_Instagram_Post__1600____900_px_.png", "test.png", nil)
	assert.NoError(t, err)
}
