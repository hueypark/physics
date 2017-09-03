package circle

import "github.com/hueypark/physics/core/shape"

type Circle struct {
	Radius float64
}

func (c *Circle) Type() int64 {
	return shape.CIRCLE
}
