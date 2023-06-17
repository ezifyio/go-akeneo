package goakeneo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProducts(t *testing.T) {
	c := MockClient()
	p, err := c.Product.GetProduct("code-1409-cj0cq", nil)
	assert.NoError(t, err)

	for key, vs := range p.Values {
		for _, v := range vs {
			if v.IsLocalized() {
				continue
			}
			result, err := v.ParseValue()
			assert.NotNil(t, key)
			assert.NoError(t, err)
			assert.NotNil(t, result)
		}
	}

}
