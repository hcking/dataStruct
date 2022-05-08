package ADT

import (
	"fmt"
	"testing"
)

func TestNewCircle(t *testing.T) {
	c := NewCircle(5, 1, 1)
	fmt.Println(c, c.GetArea(), c.GetCircumferences())
}
