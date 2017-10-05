package main

import (
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"

	"github.com/hueypark/physics/core"
	"github.com/hueypark/physics/core/body"
	"github.com/hueypark/physics/core/shape"
	"github.com/hueypark/physics/core/shape/circle"
	"github.com/hueypark/physics/core/shape/convex"
	"github.com/hueypark/physics/core/vector"
	"golang.org/x/image/colornames"
)

const WINDOW_WIDTH = 1024
const WINDOW_HEIGHT = 768

func main() {
	pixelgl.Run(run)
}

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Minkowski Difference",
		Bounds: pixel.R(-WINDOW_WIDTH/2, -WINDOW_HEIGHT/2, WINDOW_WIDTH/2, WINDOW_HEIGHT/2),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}
	imd := imdraw.New(nil)

	world := physics.New()

	circleA := body.New()
	circleA.SetStatic()
	circleA.SetShape(circle.New(50))
	circleA.SetPosition(vector.ZERO())
	world.Add(circleA)

	convexB := body.New()
	convexB.SetStatic()
	convexB.SetShape(convex.New([]vector.Vector{{-50, -50}, {-100, 0}, {70, 70}, {50, -50}, {50, 50}, {-50, 50}}))
	convexB.SetPosition(vector.Vector{100, 0})
	world.Add(convexB)

	delta := time.Second / 30
	ticker := time.NewTicker(delta)
	leftButtonClicked := false
	rightButtonClicked := false
	for range ticker.C {
		if win.Closed() {
			break
		}

		if win.JustPressed(pixelgl.MouseButtonLeft) {
			leftButtonClicked = true
		}

		if win.JustReleased(pixelgl.MouseButtonLeft) {
			leftButtonClicked = false
		}

		if leftButtonClicked {
			pos := win.MousePosition()
			circleA.SetPosition(vector.Vector{pos.X, pos.Y})
		}

		if win.JustPressed(pixelgl.MouseButtonRight) {
			rightButtonClicked = true
		}

		if win.JustReleased(pixelgl.MouseButtonRight) {
			rightButtonClicked = false
		}

		if rightButtonClicked {
			pos := win.MousePosition()
			convexB.SetPosition(vector.Vector{pos.X, pos.Y})
		}

		win.Clear(colornames.Black)
		imd.Clear()

		for _, b := range world.Bodys() {
			switch b.Shape.Type() {
			case shape.BULLET:
				drawCircle(imd, b.Position(), 1)
			case shape.CIRCLE:
				c := b.Shape.(*circle.Circle)
				drawCircle(imd, b.Position(), c.Radius)
			case shape.CONVEX:
				c := b.Shape.(*convex.Convex)
				drawConvex(imd, b.Position(), c.Hull())
			}
		}

		for _, m := range world.Manifolds() {
			for _, c := range m.Contacts() {
				start := vector.Add(c, vector.Multiply(m.Normal(), -10))
				end := vector.Add(c, vector.Multiply(m.Normal(), 10))
				drawDebugLine(imd, start, end)
			}
		}

		world.Tick(delta.Seconds())

		imd.Draw(win)
		win.Update()
	}
}

func drawCircle(imd *imdraw.IMDraw, position vector.Vector, radius float64) {
	imd.Push(pixel.V(position.X, position.Y))
	imd.Circle(radius, 1)
}

func drawConvex(imd *imdraw.IMDraw, position vector.Vector, vertices []vector.Vector) {
	for _, vertex := range vertices {
		worldPosition := vector.Add(position, vertex)
		imd.Push(pixel.V(worldPosition.X, worldPosition.Y))
	}

	first := vector.Add(position, vertices[0])
	imd.Push(pixel.V(first.X, first.Y))
	imd.Line(1)
}

func drawDebugLine(imd *imdraw.IMDraw, start, end vector.Vector) {
	imd.Push(pixel.V(start.X, start.Y), pixel.V(end.X, end.Y))
	imd.Line(1)
}
