package contact

import (
	"math"

	"github.com/hueypark/physics/core/shape"
	"github.com/hueypark/physics/core/shape/circle"
	"github.com/hueypark/physics/core/shape/convex"
	"github.com/hueypark/physics/core/vector"
)

func (c *Contact) DetectCollision() {
	lhsType := c.lhs.Shape.Type()
	rhsType := c.rhs.Shape.Type()

	switch lhsType {
	case shape.BULLET:
		switch rhsType {
		case shape.BULLET:
			break
		case shape.CIRCLE:
			c.bulletToCircle(c.rhs.Shape.(*circle.Circle))
			break
		case shape.CONVEX:
			c.bulletToConvex(c.rhs.Shape.(*convex.Convex))
			break
		}
		break
	case shape.CIRCLE:
		switch rhsType {
		case shape.BULLET:
			c.circleToBullet(c.lhs.Shape.(*circle.Circle))
			break
		case shape.CIRCLE:
			c.circleToCircle(c.lhs.Shape.(*circle.Circle), c.rhs.Shape.(*circle.Circle))
			break
		case shape.CONVEX:
			c.circleToConvex(c.lhs.Shape.(*circle.Circle), c.rhs.Shape.(*convex.Convex))
			break
		}
		break
	case shape.CONVEX:
		switch rhsType {
		case shape.BULLET:
			c.convexToBullet(c.lhs.Shape.(*convex.Convex))
		case shape.CIRCLE:
			c.convexToCircle(c.lhs.Shape.(*convex.Convex), c.rhs.Shape.(*circle.Circle))
			break
		case shape.CONVEX:
			c.convexToConvex(c.lhs.Shape.(*convex.Convex), c.rhs.Shape.(*convex.Convex))
			break
		}
		break
	}
}

func (c *Contact) bulletToCircle(rhs *circle.Circle) {
	c.normal = vector.Subtract(c.rhs.Position(), c.lhs.Position())

	distanceSquared := c.normal.SizeSquared()

	if distanceSquared >= rhs.Radius*rhs.Radius {
		return
	}

	distance := math.Sqrt(distanceSquared)

	c.penetration = rhs.Radius - distance
	c.normal.Normalize()
	c.points = append(c.points, c.lhs.Position())
}

func (c *Contact) circleToBullet(lhs *circle.Circle) {
	c.normal = vector.Subtract(c.rhs.Position(), c.lhs.Position())

	distanceSquared := c.normal.SizeSquared()

	if distanceSquared >= lhs.Radius*lhs.Radius {
		return
	}

	distance := math.Sqrt(distanceSquared)

	c.penetration = lhs.Radius - distance
	c.normal.Normalize()
	c.points = append(c.points, c.lhs.Position())
}

func (c *Contact) bulletToConvex(rhs *convex.Convex) {
}

func (c *Contact) convexToBullet(lhs *convex.Convex) {
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

func (c *Contact) circleToConvex(lhs *circle.Circle, rhs *convex.Convex) {

}

func (c *Contact) convexToCircle(lhs *convex.Convex, rhs *circle.Circle) {

}

func (c *Contact) convexToConvex(lhs, rhs *convex.Convex) {

}
