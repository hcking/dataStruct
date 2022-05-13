package list

import (
	"fmt"
	"testing"
)

func TestNewListNode(t *testing.T) {
	l := NewListNode()
	fmt.Println(l)
}
