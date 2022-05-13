package arith

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverse(t *testing.T) {
	arr0 := []int{5, 4, 3, 2, 1}
	arr1 := []int{1, 2, 3, 4, 5}
	arr2 := Reverse(arr1)
	assert.Equal(t, arr0, arr2)
}
