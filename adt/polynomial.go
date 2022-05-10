package adt

// Polynomial 多项式类型
type Polynomial struct {
	p float64 // 系数
	e int64   // 指数
}

func NewPolynomial(p float64, e int64) *Polynomial {
	pn := &Polynomial{}
	pn.p = p
	pn.e = e
	return pn
}
