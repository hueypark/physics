package util

import (
	"golang.org/x/image/colornames"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"

	"github.com/hueypark/physics/core/contact"
	"github.com/hueypark/physics/core/math/rotator"
	"github.com/hueypark/physics/core/math/vector"
)

func DrawCircle(imd *imdraw.IMDraw, position vector.Vector, radius float64) {
	imd.Color = colornames.White

	imd.Push(pixel.V(position.X, position.Y))
	imd.Circle(radius, 1)
}

func DrawConvex(imd *imdraw.IMDraw, position vector.Vector, rotation rotator.Rotator, vertices []vector.Vector) {
	imd.Color = colornames.White

	for _, vertex := range vertices {
		vertex = rotation.RotateVector(vertex)
		worldPosition := vector.Add(position, vertex)
		imd.Push(pixel.V(worldPosition.X, worldPosition.Y))
	}

	first := rotation.RotateVector(vertices[0])
	firstWorldPosition := vector.Add(position, first)
	imd.Push(pixel.V(firstWorldPosition.X, firstWorldPosition.Y))
	imd.Line(1)
}

func DrawDebugLine(imd *imdraw.IMDraw, start, end vector.Vector) {
	imd.Color = colornames.Limegreen

	imd.Push(pixel.V(start.X, start.Y), pixel.V(end.X, end.Y))
	imd.Line(2)
}

func DrawContacts(imd *imdraw.IMDraw, contacts []*contact.Contact) {
	for _, c := range contacts {
		for _, p := range c.Points() {
			end := vector.Add(p, vector.Multiply(c.Normal(), -c.Penetration()))
			DrawDebugLine(imd, p, end)
		}
	}
}
