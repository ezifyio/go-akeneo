package goakeneo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMediaOp_ListPagination(t *testing.T) {
	c := MockDLClient()
	ms, links, err := c.MediaFile.ListPagination(nil)
	assert.NoError(t, err)
	assert.NotNil(t, ms)
	assert.NotNil(t, links)
}

func TestMediaOp_GetByCode(t *testing.T) {
	c := MockDLClient()
	m, err := c.MediaFile.GetByCode("b/a/7/9/ba795607155860d543ab1d1f97a91a0dba7d98a8_____________.png", nil)
	assert.NoError(t, err)
	assert.NotNil(t, m)
}

func TestMediaOp_Download(t *testing.T) {
	c := MockDLClient()
	err := c.MediaFile.Download("/1/3/e/d/13ed17a77f6ff8748758083641d3a33e4c651d7e_0______.jpeg", "test.png", nil)
	assert.NoError(t, err)
}
