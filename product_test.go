package goakeneo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProducts(t *testing.T) {
	c := MockClient()
	products, links, err := c.Product.ListWithPagination(nil)
	assert.NoError(t, err)
	assert.NotNil(t, products)
	assert.NotNil(t, links)
	for _, p := range products {
		for _, vs := range p.Values {
			for _, v := range vs {
				result, err := v.ParseValue()
				assert.NoError(t, err)
				assert.NotNil(t, result)
			}
		}
	}
}
