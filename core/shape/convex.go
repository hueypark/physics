package shape

import "github.com/hueypark/physics/core/vector"

type Convex struct {
	Vertices []vector.Vector
}

func NewConvex(vertices []vector.Vector) *Convex {
	c := Convex{}
	c.Vertices = vertices

	return &c
}

func (c *Convex) Type() int64 {
	return CONVEX
}
