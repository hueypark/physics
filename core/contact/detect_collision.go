package contact

import (
	"math"

	"github.com/hueypark/physics/core/body"
	"github.com/hueypark/physics/core/math/vector"
	"github.com/hueypark/physics/core/shape"
	"github.com/hueypark/physics/core/shape/circle"
	"github.com/hueypark/physics/core/shape/convex"
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
			c.normal, c.penetration, c.points = bulletToCircle(c.lhs, c.rhs)
			break
		case shape.CONVEX:
			c.normal, c.penetration, c.points = bulletToConvex(c.lhs, c.rhs)
			break
		}
		break
	case shape.CIRCLE:
		switch rhsType {
		case shape.BULLET:
			c.swap()
			c.normal, c.penetration, c.points = bulletToCircle(c.lhs, c.rhs)
			break
		case shape.CIRCLE:
			c.normal, c.penetration, c.points = circleToCircle(c.lhs, c.rhs)
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

func bulletToCircle(lhs, rhs *body.Body) (normal vector.Vector, penetration float64, points []vector.Vector) {
	rhsCircle := rhs.Shape.(*circle.Circle)

	normal = vector.Subtract(rhs.Position(), lhs.Position())

	distanceSquared := normal.SizeSquared()

	if distanceSquared >= rhsCircle.Radius*rhsCircle.Radius {
		return
	}

	distance := math.Sqrt(distanceSquared)

	normal.Normalize()
	penetration = rhsCircle.Radius - distance
	points = append(points, lhs.Position())

	return normal, penetration, points
}

func bulletToConvex(lhs, rhs *body.Body) (normal vector.Vector, penetration float64, points []vector.Vector) {
	rhsConvex := rhs.Shape.(*convex.Convex)

	penetration = math.MaxFloat64

	for _, edge := range rhsConvex.Edges() {
		worldStart := rhs.Rotation().RotateVector(edge.Start)
		worldStart = vector.Add(rhs.Position(), worldStart)
		worldEnd := rhs.Rotation().RotateVector(edge.End)
		worldEnd = vector.Add(rhs.Position(), worldEnd)
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

func circleToCircle(lhs, rhs *body.Body) (normal vector.Vector, penetration float64, points []vector.Vector) {
	lhsCircle := lhs.Shape.(*circle.Circle)
	rhsCircle := rhs.Shape.(*circle.Circle)

	normal = vector.Subtract(rhs.Position(), lhs.Position())

	distanceSquared := normal.SizeSquared()
	radius := lhsCircle.Radius + rhsCircle.Radius

	if distanceSquared >= radius*radius {
		return
	}

	distance := math.Sqrt(distanceSquared)

	normal.Normalize()
	penetration = radius - distance
	points = append(points, vector.Add(
		lhs.Position(),
		vector.Add(vector.Multiply(normal, lhsCircle.Radius), vector.Multiply(normal, -0.5*penetration))))

	return normal, penetration, points
}

func circleToConvex(l, r *body.Body) (normal vector.Vector, penetration float64, points []vector.Vector) {
	lCircle := l.Shape.(*circle.Circle)
	rConvex := r.Shape.(*convex.Convex)

	minPenetration := math.MaxFloat64
	edgeNormal := vector.ZERO()
	for _, edge := range rConvex.Edges() {
		p := -vector.Dot(edge.Normal, vector.Subtract(l.Position(), vector.Add(r.Position(), edge.Start)))

		if p < -lCircle.Radius {
			return normal, penetration, points
		}

		if p < minPenetration {
			minPenetration = p
			edgeNormal = edge.Normal
		}
	}

	normal = vector.Invert(edgeNormal)
	penetration = lCircle.Radius + minPenetration
	points = append(points, vector.Add(
		l.Position(),
		vector.Add(vector.Multiply(normal, lCircle.Radius), vector.Multiply(normal, -0.5*penetration))))

	return normal, penetration, points
}

func convexToConvex(l, r *body.Body) (normal vector.Vector, penetration float64, points []vector.Vector) {
	lConvex := l.Shape.(*convex.Convex)
	rConvex := r.Shape.(*convex.Convex)

	lPenetration, lNormal, lPoint := findAxisLeastPenetration(lConvex, rConvex, l.Position(), r.Position())
	if lPenetration < 0.0 {
		return normal, penetration, points
	}

	rPenetration, rNormal, rPoint := findAxisLeastPenetration(rConvex, lConvex, r.Position(), l.Position())
	if rPenetration < 0.0 {
		return normal, penetration, points
	}

	if lPenetration < rPenetration {
		normal = lNormal
		penetration = lPenetration
		points = append(points, lPoint)
	} else {
		normal = vector.Invert(rNormal)
		penetration = rPenetration
		points = append(points, rPoint)
	}

	return normal, -penetration, points
}

func findAxisLeastPenetration(l, r *convex.Convex, lPos, rPos vector.Vector) (minPenetration float64, bestNormal vector.Vector, bestPoint vector.Vector) {
	minPenetration = math.MaxFloat64

	for _, edge := range l.Edges() {
		s := r.Support(vector.Invert(edge.Normal))

		v := vector.Add(edge.Start, lPos)
		v.Subtract(rPos)

		penetration := -vector.Dot(edge.Normal, vector.Subtract(s, v))

		if penetration < minPenetration {
			bestNormal = edge.Normal
			minPenetration = penetration
			bestPoint = vector.Add(vector.Add(s, rPos), vector.Multiply(bestNormal, penetration*0.5))
		}
	}

	return minPenetration, bestNormal, bestPoint
}
