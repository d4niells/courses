package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProduct(t *testing.T) {

	t.Run("create a valid product", func(t *testing.T) {
		p, err := NewProduct("Product 1", 10)
		assert.Nil(t, err)
		assert.NotNil(t, p)
		assert.NotEmpty(t, p.ID)
		assert.Equal(t, "Product 1", p.Name)
		assert.Equal(t, 10.0, p.Price)
	})

	t.Run("name is required", func(t *testing.T) {
		p, err := NewProduct("", 10.0)
		assert.Nil(t, p)
		assert.EqualError(t, ErrNameIsRequired, err.Error())
	})

	t.Run("price is required", func(t *testing.T) {
		p, err := NewProduct("Product 1", 0.0)
		assert.Nil(t, p)
		assert.EqualError(t, ErrPriceIsRequired, err.Error())
	})

	t.Run("price shouldn't less than zero", func(t *testing.T) {
		p, err := NewProduct("Product 1", -10.0)
		assert.Nil(t, p)
		assert.EqualError(t, ErrInvalidPrice, err.Error())
	})

}
