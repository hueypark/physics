package convex

import (
	"testing"

	"github.com/hueypark/physics/core/vector"
	"github.com/stretchr/testify/assert"
)

func TestNewConvex(t *testing.T) {
	a := assert.New(t)

	vertices := []vector.Vector{{0, 0}, {100, 0}, {100, -10}, {150, 100}, {100, 200}, {0, 210}, {-50, 100}, {30, 30}, {75, 30}}
	hull := []vector.Vector{{0, 210}, {100, 200}, {150, 100}, {100, -10}, {0, 0}, {-50, 100}}

	c := New(vertices)

	a.Equal(c.Hull(), hull)
}
