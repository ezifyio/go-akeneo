package goakeneo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchFilter(t *testing.T) {
	sf := make(SearchFilter)
	sf.Add("CREATED", ">", "2019-01-01T00:00:00+01:00")
	sf.Add("UPDATED", "<", "2019-01-01T00:00:00+01:00")
	result := sf.String()
	assert.Equal(t, `{"CREATED":[{"operator":">","value":"2019-01-01T00:00:00+01:00"}],"UPDATED":[{"operator":"<","value":"2019-01-01T00:00:00+01:00"}]}`, result)
}
