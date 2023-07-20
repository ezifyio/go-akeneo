package goakeneo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCategory(t *testing.T) {
	c := MockDLClient()
	categories, links, err := c.Category.ListWithPagination(nil)
	assert.NoError(t, err)
	assert.NotNil(t, categories)
	assert.NotNil(t, links)
	if len(categories) > 0 {
		category, err := c.Category.Get(categories[0].Code)
		assert.NoError(t, err)
		assert.NotNil(t, category)
	}
}
