package util

import (
	"golang.org/x/image/colornames"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"

	"github.com/hueypark/physics/core"
	"github.com/hueypark/physics/core/contact"
	"github.com/hueypark/physics/core/math/rotator"
	"github.com/hueypark/physics/core/math/vector"
	"github.com/hueypark/physics/core/shape"
	"github.com/hueypark/physics/core/shape/circle"
	"github.com/hueypark/physics/core/shape/convex"
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

func DrawWorld(imd *imdraw.IMDraw, world *physics.World) {
	for _, b := range world.Bodys() {
		switch b.Shape.Type() {
		case shape.BULLET:
			DrawCircle(imd, b.Position(), 3)
		case shape.CIRCLE:
			c := b.Shape.(*circle.Circle)
			DrawCircle(imd, b.Position(), c.Radius)
		case shape.CONVEX:
			c := b.Shape.(*convex.Convex)
			DrawConvex(imd, b.Position(), b.Rotation(), c.Hull())
		}
	}

	DrawContacts(imd, world.Contacts())
}

func DrawContacts(imd *imdraw.IMDraw, contacts []*contact.Contact) {
	for _, c := range contacts {
		for _, p := range c.Points() {
			end := vector.Add(p, vector.Multiply(c.Normal(), -c.Penetration()))
			DrawDebugLine(imd, p, end)
		}
	}
}
