package convex

import (
	"testing"

	"github.com/hueypark/physics/core/vector"
	"github.com/stretchr/testify/assert"
)

func TestNewConvex(t *testing.T) {
	a := assert.New(t)

	vertices := []vector.Vector{{-100, -100}, {100, -100}, {100, 100}, {-100, 100}, {0, 0}}
	nomalizedVertices := []vector.Vector{{-100, -100}, {100, -100}, {100, 100}, {-100, 100}}

	c := New(vertices)

	a.Equal(c.Vertices, nomalizedVertices)
}
