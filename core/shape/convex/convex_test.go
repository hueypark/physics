package convex

import (
	"testing"

	"github.com/hueypark/physics/core/vector"
	"github.com/stretchr/testify/assert"
)

func TestNewConvex(t *testing.T) {
	a := assert.New(t)

	vertices := []vector.Vector{{0, 0}, {100, 0}, {100, -10}, {150, 100}, {100, 200}, {0, 210}, {-50, 100}, {30, 30}, {75, 30}}
	hull := []vector.Vector{{-50, 100}, {0, 0}, {100, -10}, {150, 100}, {100, 200}, {0, 210}}

	c := New(vertices)

	a.Equal(hull, c.Hull())
}
