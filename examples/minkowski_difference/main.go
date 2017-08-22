package main

import (
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"

	"github.com/hueypark/physics/core"
	"github.com/hueypark/physics/core/vector"
	"github.com/hueypark/physics/core/body"
	"github.com/hueypark/physics/core/shape"
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
	b := body.New()
	b.SetStatic()
	b.SetShape(shape.NewConvex([]vector.Vector{{-50,-50},{50,-50},{50,50},{-50, 50}}))
	b.SetPosition(vector.Vector{50, 50})
	world.Add(b)

	delta := time.Second / 30
	ticker := time.NewTicker(delta)
	for range ticker.C {
		if win.Closed() {
			break
		}

		win.Clear(colornames.Black)
		imd.Clear()

		for _, b := range world.Bodys() {
			if b.Shape.Type() == shape.CONVEX {
				convex := b.Shape.(*shape.Convex)
				drawConvex(imd, b.Position(), convex.Vertices)
			}
		}

		world.Tick(delta.Seconds())

		imd.Draw(win)
		win.Update()
	}
}

func drawConvex(imd *imdraw.IMDraw, position vector.Vector, vertices []vector.Vector) {
	imd.Color = colornames.Limegreen
	for _, vertex := range vertices {
		worldPosition := vector.Add(position, vertex)
		imd.Push(pixel.V(worldPosition.X, worldPosition.Y))
	}

	first := vector.Add(position, vertices[0])
	imd.Push(pixel.V(first.X, first.X))
	imd.Line(5)
}
