package rotator

import (
	"math"

	"github.com/hueypark/physics/core/math/matrix"
	"github.com/hueypark/physics/core/math/util"
	"github.com/hueypark/physics/core/math/vector"
)

type Rotator float64

func (r Rotator) RotateVector(v vector.Vector) vector.Vector {
	return r.RotationMatrix().TransformVector(v)
}

func (r Rotator) RotationMatrix() (m matrix.Matrix) {
	rad := util.DegToRad(float64(r))

	c := math.Cos(rad)
	s := math.Sin(rad)

	m.M[0][0] = c
	m.M[0][1] = -s
	m.M[1][0] = s
	m.M[1][1] = c

	return m
}
