package test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/hueypark/physics/core/vector"
)

func TestNewConvex(t *testing.T) {
	a := assert.New(t)

	point := vector.Vector{30, 30}
	triangleVertexA := vector.Vector{0, 0}
	triangleVertexB := vector.Vector{100, 0}
	triangleVertexC := vector.Vector{0, 100}

	a.True(PointInTriangle(point, triangleVertexA, triangleVertexB, triangleVertexC))
}
