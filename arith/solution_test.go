package arith

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTopic1(t *testing.T) {
	var nums = []int{3, 2, 1, 0, 6}
	res := topic1(nums, 7)
	assert.Equal(t, []int{2, 4}, res)
}

func TestTopic1Map(t *testing.T) {
	var nums = []int{3, 2, 1, 0, 6}
	res := topic1Map(nums, 7)
	assert.Equal(t, []int{2, 4}, res)
}
