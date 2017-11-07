package main

import (
	"time"

	"golang.org/x/image/colornames"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"

	"github.com/hueypark/physics/core"
	"github.com/hueypark/physics/core/body"
	"github.com/hueypark/physics/core/math/vector"
	"github.com/hueypark/physics/core/shape"
	"github.com/hueypark/physics/core/shape/circle"
	"github.com/hueypark/physics/core/shape/convex"
	"github.com/hueypark/physics/examples/util"
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

	circleB := body.New()
	circleB.SetStatic()
	circleB.SetShape(circle.New(100))
	circleB.SetPosition(vector.Vector{100, 0})
	world.Add(circleB)

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
			circleB.SetPosition(vector.Vector{pos.X, pos.Y})
		}

		win.Clear(colornames.Black)
		imd.Clear()

		for _, b := range world.Bodys() {
			switch b.Shape.Type() {
			case shape.BULLET:
				util.DrawCircle(imd, b.Position(), 1)
			case shape.CIRCLE:
				c := b.Shape.(*circle.Circle)
				util.DrawCircle(imd, b.Position(), c.Radius)
			case shape.CONVEX:
				c := b.Shape.(*convex.Convex)
				util.DrawConvex(imd, b.Position(), b.Rotation(), c.Hull())
			}
		}

		util.DrawContacts(imd, world.Contacts())

		world.Tick(delta.Seconds())

		imd.Draw(win)
		win.Update()
	}
}
