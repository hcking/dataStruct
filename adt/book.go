package adt

const GBookMaxSize = 1000

type Book struct {
	no    int64   // 唯一编号
	name  string  // 书名
	price float64 // 价格
}
