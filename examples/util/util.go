package util

import (
	"golang.org/x/image/colornames"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"

	"github.com/hueypark/physics/core/math/vector"
)

func DrawCircle(imd *imdraw.IMDraw, position vector.Vector, radius float64) {
	imd.Color = colornames.White

	imd.Push(pixel.V(position.X, position.Y))
	imd.Circle(radius, 1)
}

func DrawConvex(imd *imdraw.IMDraw, position vector.Vector, vertices []vector.Vector) {
	imd.Color = colornames.White

	for _, vertex := range vertices {
		worldPosition := vector.Add(position, vertex)
		imd.Push(pixel.V(worldPosition.X, worldPosition.Y))
	}

	first := vector.Add(position, vertices[0])
	imd.Push(pixel.V(first.X, first.Y))
	imd.Line(1)
}

func DrawDebugLine(imd *imdraw.IMDraw, start, end vector.Vector) {
	imd.Color = colornames.Limegreen

	imd.Push(pixel.V(start.X, start.Y), pixel.V(end.X, end.Y))
	imd.Line(2)
}
