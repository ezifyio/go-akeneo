package goakeneo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLocaleOp_ListWithPagination(t *testing.T) {
	c := MockClient()
	locales, pagi, err := c.Locale.ListWithPagination(nil)
	assert.NoError(t, err)
	assert.NotNil(t, locales)
	assert.NotNil(t, pagi)
}
