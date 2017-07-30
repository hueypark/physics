package vector

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNormalize(t *testing.T) {
	a := assert.New(t)

	v1 := Vector{10, 0, 0}
	v1.Normalize()

	nV1 := Vector{1, 0, 0}
	a.Equal(v1, nV1)

	v2 := Vector{3, 4, 0}
	v2.Normalize()

	nV2 := Vector{0.6, 0.8, 0}
	a.Equal(v2, nV2)
}

func TestProduct(t *testing.T) {
	a := assert.New(t)

	v1 := Vector{1,2,3}
	v2 := Vector{4,5,6}
	v3 := Product(v1, v2)

	a.Equal(v3, Vector{4, 10, 18})
}

func TestDot(t *testing.T) {
	a := assert.New(t)

	v1 := Vector{1,2,3}
	v2 := Vector{4,5,6}
	dot := Dot(v1, v2)

	a.Equal(dot, 32.0)
}

func TestCross(t *testing.T) {
	a := assert.New(t)

	v1 := Vector{1, 0, 0}
	v2 := Vector{0, 1, 0}
	cross := Cross(v1, v2)

	a.Equal(cross, Vector{0,0,1})

	v3 := Vector{1,0,0}
	v4 := Vector{2,0,0}
	cross2 := Cross(v3, v4)

	a.Equal(cross2, Vector{0,0,0})
}