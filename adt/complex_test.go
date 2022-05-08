package adt

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewComplex(t *testing.T) {
	c0 := NewComplex(1, 1)
	c1 := NewComplex(1, 1)
	c2 := NewComplex(2, 2)
	c3 := NewComplex(3, 3)

	c1.add(c2)
	assert.Equal(t, c1, c3)

	c1.divide(c0)
	assert.Equal(t, c1, c3)

}
