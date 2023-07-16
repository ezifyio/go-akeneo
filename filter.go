package goakeneo

import (
	"encoding/json"
	"html"
)

// SearchFilter is a map of search filters,see :
// https://api.akeneo.com/documentation/filter.html
type SearchFilter map[string][]map[string]interface{}

func (sf SearchFilter) String() string {
	t, _ := json.Marshal(sf)
	return html.UnescapeString(string(t))
}

// Add adds a new filter to the search filter
func (sf SearchFilter) Add(key, operator string, value any) {
	sf[key] = append(sf[key], map[string]interface{}{"operator": operator, "value": value})
}
