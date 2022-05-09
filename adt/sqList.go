package adt

// 线性表

// PolynomialSqList 多项式顺序结构（线性表)类型
type PolynomialSqList struct {
	SqList
	p *Polynomial // 存储空间的基地址
}

// BookSqList 图书顺序结构（线性表)类型
type BookSqList struct {
	SqList
	b *Book
}

type SqList struct {
	length int64
}
