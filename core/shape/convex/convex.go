package convex

import (
	"math"

	"github.com/hueypark/physics/core/math/rotator"
	"github.com/hueypark/physics/core/math/vector"
	"github.com/hueypark/physics/core/shape"
)

type Convex struct {
	vertices []vector.Vector
	hull     []vector.Vector
	edges    []Edge
}

type Edge struct {
	Start vector.Vector
	End   vector.Vector
}

func New(vertices []vector.Vector) *Convex {
	c := Convex{vertices, nil, nil}

	return &c
}

func (c *Convex) Type() int64 {
	return shape.CONVEX
}

// Hull is ccw
func (c *Convex) Hull() []vector.Vector {
	if c.hull == nil {
		minX, maxX := c.getExtremePoints()
		c.hull = append(c.quickHull(c.vertices, maxX, minX), c.quickHull(c.vertices, minX, maxX)...)
	}

	return c.hull
}

// Edge is ccw
func (c *Convex) Edges() []Edge {
	if c.edges == nil {
		hull := c.Hull()
		for i, vertex := range hull {
			nextIndex := i + 1
			if len(hull) <= nextIndex {
				nextIndex = 0
			}
			c.edges = append(c.edges, Edge{vertex, hull[nextIndex]})
		}
	}

	return c.edges
}

func MinkowskiDifference(a Convex, posA vector.Vector, rotA rotator.Rotator, b Convex, posB vector.Vector, rotB rotator.Rotator) *Convex {
	vertices := []vector.Vector{}

	for _, vertexA := range a.Hull() {
		for _, vertexB := range b.Hull() {
			vertexRotA := rotA.RotateVector(vertexA)
			vertexRotB := rotB.RotateVector(vertexB)
			worldA := vector.Add(vertexRotA, posA)
			worldB := vector.Subtract(vector.Vector{}, vector.Add(vertexRotB, posB))
			vertices = append(vertices, vector.Add(worldA, worldB))
		}
	}

	return New(vertices)
}

func (c *Convex) quickHull(points []vector.Vector, start, end vector.Vector) []vector.Vector {
	pointDistanceIndicators := c.getLhsPointDistanceIndicatorMap(points, start, end)
	if len(pointDistanceIndicators) == 0 {
		return []vector.Vector{end}
	}

	farthestPoint := c.getFarthestPoint(pointDistanceIndicators)

	newPoints := []vector.Vector{}
	for point := range pointDistanceIndicators {
		newPoints = append(newPoints, point)
	}

	return append(
		c.quickHull(newPoints, farthestPoint, end),
		c.quickHull(newPoints, start, farthestPoint)...)
}

func (c *Convex) InHull(position vector.Vector, rotation rotator.Rotator, point vector.Vector) bool {
	for _, edge := range c.Edges() {
		if vector.Subtract(point, vector.Add(position, rotation.RotateVector(edge.Start))).OnTheRight(vector.Subtract(vector.Add(position, rotation.RotateVector(edge.End)), vector.Add(position, rotation.RotateVector(edge.Start)))) == false {
			return false
		}
	}

	return true
}

func (c *Convex) getExtremePoints() (minX, maxX vector.Vector) {
	minX = vector.Vector{math.MaxFloat64, 0}
	maxX = vector.Vector{-math.MaxFloat64, 0}

	for _, p := range c.vertices {
		if p.X < minX.X {
			minX = p
		}

		if maxX.X < p.X {
			maxX = p
		}
	}

	return minX, maxX
}

func (c *Convex) getLhsPointDistanceIndicatorMap(points []vector.Vector, start, end vector.Vector) map[vector.Vector]float64 {
	pointDistanceIndicatorMap := make(map[vector.Vector]float64)

	for _, point := range points {
		distanceIndicator := c.getDistanceIndicator(point, start, end)
		if distanceIndicator > 0 {
			pointDistanceIndicatorMap[point] = distanceIndicator
		}
	}

	return pointDistanceIndicatorMap
}

func (c *Convex) getDistanceIndicator(point, start, end vector.Vector) float64 {
	vLine := vector.Subtract(end, start)

	vPoint := vector.Subtract(point, start)

	return vector.Cross(vLine, vPoint)
}

func (c *Convex) getFarthestPoint(pointDistanceIndicatorMap map[vector.Vector]float64) (farthestPoint vector.Vector) {
	maxDistanceIndicator := -math.MaxFloat64
	for point, distanceIndicator := range pointDistanceIndicatorMap {
		if maxDistanceIndicator < distanceIndicator {
			maxDistanceIndicator = distanceIndicator
			farthestPoint = point
		}
	}

	return farthestPoint
}
