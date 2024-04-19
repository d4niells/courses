package tax

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateTax(t *testing.T) {
	tax, err := CalculateTax(500.0)
	assert.Nil(t, err)
	assert.Equal(t, tax, 5.0)

	tax, err = CalculateTax(0)
	assert.EqualError(t, err, "Greater than 0")
	assert.Equal(t, 0.0, tax)
}
