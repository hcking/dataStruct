package adt

// Complex 复数
type Complex struct {
	realPart float64
	imagPart float64
}

func NewComplex(real float64, imag float64) *Complex {
	c := &Complex{}
	c.realPart = real
	c.imagPart = imag
	return c
}

func (c *Complex) add(c2 *Complex) {
	c.realPart += c2.realPart
	c.imagPart += c2.imagPart
}

func (c *Complex) multiply(c2 *Complex) {
	c.realPart *= c2.realPart
	c.imagPart *= c2.imagPart
}

func (c *Complex) divide(c2 *Complex) {
	c.realPart /= c2.realPart
	c.imagPart /= c2.imagPart
}

func ComplexEqual(c1 *Complex, c2 *Complex) bool {
	return c1.realPart == c2.realPart && c1.imagPart == c2.imagPart
}
