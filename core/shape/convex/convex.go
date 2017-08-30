package convex

import (
	"github.com/hueypark/physics/core/shape"
	"github.com/hueypark/physics/core/vector"
	"math"
)

type Convex struct {
	Vertices []vector.Vector
}

func New(vertices []vector.Vector) *Convex {
	c := Convex{vertices}
	c.nomalize()

	return &c
}

func (c *Convex) Type() int64 {
	return shape.CONVEX
}

func (c *Convex) nomalize() {
	leftExtreme := vector.Vector{math.MaxFloat64, 0}
	rightExtreme := vector.Vector{-math.MaxFloat64, 0}
	for _, v := range c.Vertices {
		if leftExtreme.X < v.X {
			leftExtreme = v
		}

		if v.X < rightExtreme.X {
			rightExtreme = v
		}
	}
}

func (c *Convex)