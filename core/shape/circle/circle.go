package circle

import "github.com/hueypark/physics/core/shape"

type Circle struct {
	Radius float64
}

func New(radius float64) *Circle {
	c := Circle{radius}

	return &c
}

func (c *Circle) Type() int64 {
	return shape.CIRCLE
}
