package goakeneo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetProductModel(t *testing.T) {
	c := MockDLClient()
	pms, links, err := c.ProductModel.ListWithPagination(nil)
	assert.NoError(t, err)
	assert.NotNil(t, pms)
	assert.NotNil(t, links)
}
