package contact

import (
	"math"

	"github.com/hueypark/physics/core/body"
	"github.com/hueypark/physics/core/closest_point"
	"github.com/hueypark/physics/core/shape"
	"github.com/hueypark/physics/core/shape/circle"
	"github.com/hueypark/physics/core/shape/convex"
	"github.com/hueypark/physics/core/math/vector"
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
			c.normal, c.penetration, c.points = bulletToConvex(c.lhs, c.rhs)
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
			c.normal, c.penetration, c.points = circleToConvex(c.lhs, c.rhs)
			break
		}
		break
	case shape.CONVEX:
		switch rhsType {
		case shape.BULLET:
			c.swap()
			c.normal, c.penetration, c.points = bulletToConvex(c.lhs, c.rhs)
		case shape.CIRCLE:
			c.swap()
			c.normal, c.penetration, c.points = circleToConvex(c.lhs, c.rhs)
			break
		case shape.CONVEX:
			c.normal, c.penetration, c.points = convexToConvex(c.lhs, c.rhs)
			break
		}
		break
	}
}

func (c *Contact) swap() {
	c.lhs, c.rhs = c.rhs, c.lhs
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

func bulletToConvex(lhs, rhs *body.Body) (normal vector.Vector, penetration float64, points []vector.Vector) {
	rhsConvex := rhs.Shape.(*convex.Convex)

	penetration = math.MaxFloat64

	for _, edge := range rhsConvex.Edges() {
		worldStart := vector.Add(rhs.Position(), edge.Start)
		worldEnd := vector.Add(rhs.Position(), edge.End)
		edgeVector := vector.Subtract(worldEnd, worldStart)
		pointVector := vector.Subtract(lhs.Position(), worldStart)

		if !pointVector.OnTheRight(edgeVector) {
			normal = vector.Vector{}
			penetration = 0
			return normal, penetration, points
		}

		perpendicular := vector.Vector{-edgeVector.Y, edgeVector.X}
		perpendicular.Normalize()

		lhsVector := vector.Subtract(lhs.Position(), worldStart)

		proj := vector.Multiply(perpendicular, vector.Dot(lhsVector, perpendicular))

		if proj.Size() < penetration {
			normal = perpendicular
			penetration = proj.Size()
		}
	}

	points = append(points, lhs.Position())
	return normal, penetration, points
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

func circleToConvex(lhs, rhs *body.Body) (normal vector.Vector, penetration float64, points []vector.Vector) {
	lhsCircle := lhs.Shape.(*circle.Circle)
	rhsConvex := rhs.Shape.(*convex.Convex)

	penetration = math.MaxFloat64
	var selectedEdge convex.Edge

	for _, edge := range rhsConvex.Edges() {
		worldStart := vector.Add(rhs.Position(), edge.Start)
		worldEnd := vector.Add(rhs.Position(), edge.End)
		edgeVector := vector.Subtract(worldEnd, worldStart)

		perpendicular := vector.Vector{-edgeVector.Y, edgeVector.X}
		perpendicular.Normalize()

		lhsVector := vector.Subtract(lhs.Position(), worldStart)

		projSize := vector.Dot(lhsVector, perpendicular)

		if projSize < -lhsCircle.Radius {
			return vector.ZERO(), 0, points
		}

		if projSize < penetration {
			normal = perpendicular
			penetration = projSize
			selectedEdge = edge
		}
	}

	worldStart := vector.Add(rhs.Position(), selectedEdge.Start)
	worldEnd := vector.Add(rhs.Position(), selectedEdge.End)
	closestPoint := closest_point.LineSegmentToPoint(lhs.Position(), worldStart, worldEnd)
	if lhsCircle.Radius*lhsCircle.Radius < vector.Subtract(lhs.Position(), closestPoint).SizeSquared() {
		return vector.ZERO(), 0, points
	}

	points = append(points, closestPoint)

	return normal, penetration, points
}

func convexToConvex(lhs, rhs *body.Body) (normal vector.Vector, penetration float64, points []vector.Vector) {
	lhsConvex := lhs.Shape.(*convex.Convex)
	rhsConvex := rhs.Shape.(*convex.Convex)

	minkowskiDifference := convex.MinkowskiDifference(*rhsConvex, rhs.Position(), rhs.Rotation(), *lhsConvex, lhs.Position(), lhs.Rotation())

	penetration = math.MaxFloat64

	for _, edge := range minkowskiDifference.Edges() {
		edgeVector := vector.Subtract(edge.End, edge.Start)
		pointVector := vector.Subtract(vector.ZERO(), edge.Start)

		if !pointVector.OnTheRight(edgeVector) {
			normal = vector.Vector{}
			penetration = 0
			return normal, penetration, points
		}

		perpendicular := vector.Vector{-edgeVector.Y, edgeVector.X}
		perpendicular.Normalize()

		lhsVector := vector.Subtract(vector.Vector{}, edge.Start)

		proj := vector.Multiply(perpendicular, vector.Dot(lhsVector, perpendicular))

		if proj.Size() < penetration {
			normal = perpendicular
			penetration = proj.Size()
		}
	}

	for _, point := range lhsConvex.Hull() {
		worldPoint := vector.Add(lhs.Position(), lhs.Rotation().RotateVector(point))
		if rhsConvex.InHull(rhs.Position(), rhs.Rotation(), worldPoint) {
			points = append(points, worldPoint)
		}
	}

	for _, point := range rhsConvex.Hull() {
		worldPoint := vector.Add(rhs.Position(), rhs.Rotation().RotateVector(point))
		if lhsConvex.InHull(lhs.Position(), lhs.Rotation(), worldPoint) {
			points = append(points, worldPoint)
		}
	}

	return normal, penetration, points
}
