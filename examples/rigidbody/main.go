package main

import (
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"

	"github.com/hueypark/framework/core/random"
	"github.com/hueypark/physics/core"
	"github.com/hueypark/physics/core/body"
	"github.com/hueypark/physics/core/shape"
	"github.com/hueypark/physics/core/vector"
)

const WINDOW_WIDTH = 1024
const WINDOW_HEIGHT = 768

func main() {
	pixelgl.Run(run)
}

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Rigidbody",
		Bounds: pixel.R(-WINDOW_WIDTH/2, -WINDOW_HEIGHT/2, WINDOW_WIDTH/2, WINDOW_HEIGHT/2),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}
	imd := imdraw.New(nil)

	world := physics.New()
	plane := createCircle(100, vector.Vector{0, -200}, vector.Vector{})
	plane.SetStatic()
	world.Add(plane)

	delta := time.Second / 30
	ticker := time.NewTicker(delta)
	respawnTime := time.Duration(0)
	for range ticker.C {
		if win.Closed() {
			break
		}

		const RESPAWN_TIME = time.Millisecond * 100
		respawnTime -= delta

		if respawnTime < 0 {
			respawnTime = RESPAWN_TIME
			world.Add(createCircle(random.FRandom(5, 20), vector.Vector{random.FRandom(-300, 300), 0}, vector.Vector{0, random.FRandom(100, 300)}))
		}

		world.Tick(delta.Seconds())

		win.Clear(colornames.Black)
		imd.Clear()

		for _, b := range world.Bodys() {
			if isOutbound(b.Position()) {
				world.ReservedDelete(b.Id())
			}

			if b.Shape.Type() == shape.CIRCLE {
				circle := b.Shape.(*shape.Circle)
				drawCircle(imd, b.Position(), circle.Radius)
			}
		}

		for _, m := range world.Manifolds() {
			for _, c := range m.Contacts() {
				start := vector.Add(c, vector.Multiply(m.Normal(), -10))
				end := vector.Add(c, vector.Multiply(m.Normal(), 10))
				drawDebugLine(imd, start, end)
			}
		}

		imd.Draw(win)

		win.Update()
	}
}

func createCircle(radius float64, position vector.Vector, velocity vector.Vector) *body.Body {
	b := body.New()
	b.SetMass(1)
	b.SetShape(&shape.Circle{radius})
	b.SetPosition(position)
	b.Velocity = velocity

	return b
}

func drawCircle(imd *imdraw.IMDraw, position vector.Vector, radius float64) {
	imd.Color = colornames.Limegreen
	imd.Push(pixel.V(position.X, position.Y))
	imd.Circle(radius, 1)
}

func drawDebugLine(imd *imdraw.IMDraw, start, end vector.Vector) {
	imd.Color = colornames.Limegreen
	imd.Push(pixel.V(start.X, start.Y), pixel.V(end.X, end.Y))
	imd.Line(1)
}

func isOutbound(position vector.Vector) bool {
	const MARGIN = 300
	if position.X < -WINDOW_WIDTH/2-MARGIN ||
		position.X > WINDOW_WIDTH/2+MARGIN ||
		position.Y < -WINDOW_HEIGHT/2-MARGIN ||
		position.Y > WINDOW_HEIGHT/2+MARGIN {
		//fmt.Println("OUT")
		return true
	}
	return false
}
