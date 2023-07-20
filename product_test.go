package goakeneo

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProducts(t *testing.T) {
	c := MockDLClient()
	p, err := c.Product.GetProduct("code-15768-vw128j", nil)
	assert.NoError(t, err)

	for key, vs := range p.Values {
		for _, v := range vs {
			result, err := v.ParseValue()
			if err != nil {
				t.Error(err)
				t.Errorf("key: %s, value: %v,result:%v", key, v, result)
			}
		}
	}
}

func TestProductOp_GetAllProducts(t *testing.T) {
	c := MockDLClient()
	prodChan, errChan := c.Product.GetAllProducts(context.Background(), nil)
	go func() {
		for err := range errChan {
			t.Error(err)
		}
	}()
	for p := range prodChan {
		if p.Identifier == "code-9200-eprcg" {
			if v, ok := p.Values["<spu>"]; ok {
				t.Logf("key: %s, value: %v", "<spu>", v)
			} else {
				t.Errorf("does not have key: %s", "<spu>")
			}
		}

	}
}
