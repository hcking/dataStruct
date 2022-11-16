package arith

import (
	"fmt"
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

func TestTopic2(t *testing.T) {
	//topic2()
	return
}

func TestTopic3(t *testing.T) {
	s := "hello"
	res := topic3(s)
	fmt.Println(res)
	assert.Equal(t, res, 3)
	return
}

func TestTopic4(t *testing.T) {
	//findMedianSortedArrays()
}

func TestTopic(t *testing.T) {
	return
}
