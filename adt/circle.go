package adt

const PI = 3.14

// Circle 圆
type Circle struct {
	r  float64
	x  float64
	y  float64
	pi float64
}

func NewCircle(r float64, x float64, y float64) *Circle {
	var c = &Circle{}
	c.r = r
	c.x = x
	c.y = y
	c.pi = PI
	return c
}

func (c *Circle) GetArea() float64 {
	return PI * (c.r * c.r)
}

func (c *Circle) GetCircumferences() float64 {
	return PI * c.r * 2
}
