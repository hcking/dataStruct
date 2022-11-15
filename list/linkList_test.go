package list

import (
	"fmt"
	"testing"
)

func TestNewListNode(t *testing.T) {
	l := NewNode()
	fmt.Println(l)
}
