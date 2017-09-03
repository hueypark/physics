package convex

import (
	"math"

	"github.com/hueypark/physics/core/shape"
	"github.com/hueypark/physics/core/vector"
)

type Convex struct {
	vertices []vector.Vector
	hull     []vector.Vector
}

func New(vertices []vector.Vector) *Convex {
	c := Convex{vertices, nil}

	return &c
}

func (c *Convex) Type() int64 {
	return shape.CONVEX
}

func (c *Convex) Hull() []vector.Vector {
	if c.hull == nil {
		minX, maxX := c.getExtremePoints()
		c.hull = append(c.quickHull(c.vertices, minX, maxX), c.quickHull(c.vertices, maxX, minX)...)
	}

	return c.hull
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
		c.quickHull(newPoints, start, farthestPoint),
		c.quickHull(newPoints, farthestPoint, end)...)
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

	return vector.Cross(vPoint, vLine)
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
