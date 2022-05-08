package ADT

import (
	"fmt"
	"testing"
)

func TestNewComplex(t *testing.T) {
	c0 := NewComplex(1, 1)
	c1 := NewComplex(1, 1)
	c2 := NewComplex(2, 2)

	fmt.Println(ComplexEqual(c0, c1), ComplexEqual(c0, c2))
	c1.add(c2)
	c0.add(c2)
	c1.multiply(c2)
	c1.divide(c2)
	fmt.Println(c1, c2, ComplexEqual(c0, c1))
}
