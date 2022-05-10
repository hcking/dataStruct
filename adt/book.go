package adt

// Book 图书类型
type Book struct {
	no    int64   // 唯一编号
	name  string  // 书名
	price float64 // 价格
}

func NewBook(no int64, name string, price float64) *Book {
	b := &Book{}
	b.no = no
	b.name = name
	b.price = price
	return b
}
