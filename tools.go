package goakeneo

import (
	"net/url"

	"github.com/google/go-querystring/query"
	"github.com/pkg/errors"
)

// StructToURLValues converts a struct to url.Values
func structToURLValues(s interface{}) (url.Values, error) {
	v, err := query.Values(s)
	if err != nil {
		return nil, errors.Wrap(err, "unable to convert struct to url.Values, check the struct fields")
	}
	return v, nil
}
