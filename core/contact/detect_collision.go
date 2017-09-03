package contact

import (
	"math"

	"github.com/hueypark/physics/core/shape"
	"github.com/hueypark/physics/core/shape/circle"
	"github.com/hueypark/physics/core/vector"
)

func (c *Contact) DetectCollision() {
	lhsType := c.lhs.Shape.Type()
	rhsType := c.rhs.Shape.Type()

	if lhsType == shape.CIRCLE && rhsType == shape.CIRCLE {
		c.circleToCircle(c.lhs.Shape.(*circle.Circle), c.rhs.Shape.(*circle.Circle))
	}
}

func (c *Contact) circleToCircle(lhsCircle, rhsCircle *circle.Circle) {
	c.normal = vector.Subtract(c.rhs.Position(), c.lhs.Position())

	distanceSquared := c.normal.SizeSquared()
	radius := lhsCircle.Radius + rhsCircle.Radius

	if distanceSquared >= radius*radius {
		return
	}

	distance := math.Sqrt(distanceSquared)

	c.penetration = radius - distance
	c.normal.Normalize()
	c.points = append(c.points, vector.Add(vector.Multiply(c.normal, lhsCircle.Radius), c.lhs.Position()))
}
