package adt

import (
	"fmt"
	"testing"
)

func TestNewSqList(t *testing.T) {
	book1 := NewBook(1, "数据结构", 100)
	book2 := NewBook(2, "数据结构2", 200)
	bookList := NewSqList()
	fmt.Println(bookList, book1, book2)
}
