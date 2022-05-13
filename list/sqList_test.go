package list

import (
	"dataStruct/adt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewSqList(t *testing.T) {
	book1 := adt.NewBook(1, "data struct 1", 100)
	book2 := adt.NewBook(2, "data struct 2", 200)
	bookList := NewSqList(5)
	bookList.Append(book1)
	bookList.Append(book2)
	assert.Equal(t, bookList.Get(0), book1)
	assert.Equal(t, bookList.Get(1), book2)
	assert.Equal(t, bookList.Get(2), nil)
	assert.Equal(t, bookList.Get(100), nil)
}
