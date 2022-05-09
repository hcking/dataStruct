package adt

// GPolynomialMaxSize 多项式可能达到的最大长度
const GPolynomialMaxSize = 1000

// Polynomial 多项式类型
type Polynomial struct {
	p float64 // 系数
	e int64   // 指数
}
