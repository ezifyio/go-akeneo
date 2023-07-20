package goakeneo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLocaleOp_ListWithPagination(t *testing.T) {
	c := MockDLClient()
	locales, pagi, err := c.Locale.ListWithPagination(nil)
	assert.NoError(t, err)
	assert.NotNil(t, locales)
	assert.NotNil(t, pagi)
}
