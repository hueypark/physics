package main

import (
	"image/color"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"

	"github.com/hueypark/physics/core"
	"github.com/hueypark/physics/core/body"
	"github.com/hueypark/physics/core/math/vector"
	"github.com/hueypark/physics/core/shape"
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

	convexA := body.New()
	convexA.SetStatic()
	convexA.SetShape(convex.New([]vector.Vector{{-50, -50}, {50, -50}, {0, 100}, {50, 50}, {-50, 50}}))
	convexA.SetPosition(vector.Vector{0, 0})
	convexA.SetRotation(90)
	world.Add(convexA)

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
			convexA.SetPosition(vector.Vector{pos.X, pos.Y})
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
			if b.Shape.Type() == shape.CONVEX {
				convex := b.Shape.(*convex.Convex)
				util.DrawConvex(imd, b.Position(), b.Rotation(), convex.Hull())
			}
		}

		convexMD := convex.MinkowskiDifference(
			*convexA.Shape.(*convex.Convex), convexA.Position(), convexA.Rotation(),
			*convexB.Shape.(*convex.Convex), convexB.Position(), convexB.Rotation())
		util.DrawConvex(imd, vector.Vector{}, 0, convexMD.Hull())

		drawLine(imd, vector.Vector{0, WINDOW_HEIGHT}, vector.Vector{0, -WINDOW_HEIGHT}, colornames.White)
		drawLine(imd, vector.Vector{WINDOW_WIDTH, 0}, vector.Vector{-WINDOW_WIDTH, 0}, colornames.White)

		world.Tick(delta.Seconds())

		imd.Draw(win)
		win.Update()
	}
}

func drawLine(imd *imdraw.IMDraw, start, end vector.Vector, c color.Color) {
	imd.Color = c
	imd.Push(pixel.V(start.X, start.Y))
	imd.Push(pixel.V(end.X, end.Y))
	imd.Line(3)
}
