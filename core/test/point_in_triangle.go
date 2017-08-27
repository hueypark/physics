package test

import (
	"github.com/hueypark/physics/core/vector"
)

func PointInTriangle(p, a, b, c vector.Vector) bool {
	if vector.Cross(vector.Subtract(p, a), vector.Subtract(b, a)) < 0 {
		return false
	}

	if vector.Cross(vector.Subtract(p, b), vector.Subtract(c, b)) < 0 {
		return false
	}

	if vector.Cross(vector.Subtract(p, c), vector.Subtract(a, c)) < 0 {
		return false
	}

	return true
}
