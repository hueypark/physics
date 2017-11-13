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
	"github.com/hueypark/physics/core/math/rotator"
	"github.com/hueypark/physics/core/math/vector"
	"github.com/hueypark/physics/core/shape/bullet"
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
	plane := createConvex(
		[]vector.Vector{
			{300, -100},
			{300, -120},
			{-300, -100},
			{-300, -120},
		},
		vector.Vector{0, -200},
		rotator.ZERO(),
		vector.Vector{})
	plane.SetStatic()
	world.Add(plane)

	delta := time.Second / 30
	ticker := time.NewTicker(delta)
	respawnTime := time.Duration(0)
	for range ticker.C {
		if win.Closed() {
			break
		}

		const RESPAWN_TIME = time.Millisecond * 500
		respawnTime -= delta

		if respawnTime < 0 {
			respawnTime = RESPAWN_TIME
			world.Add(
				createRandomShape(
					vector.Vector{random.FRandom(-300, 300), 0},
					rotator.Rotator{random.FRandom(180.0, 360.0)},
					vector.Vector{0, 0}))
		}

		world.Tick(delta.Seconds())

		win.Clear(colornames.Black)
		imd.Clear()

		util.DrawWorld(imd, world)

		imd.Draw(win)

		win.Update()
	}
}

func createRandomShape(position vector.Vector, rotation rotator.Rotator, velocity vector.Vector) *body.Body {
	var b *body.Body
	switch random.Random(1, 2) {
	case 0:
		b = createBullet(position, velocity)
	case 1:
		b = createCircle(random.FRandom(10, 50), position, velocity)
	case 2:
		b = createConvex(
			[]vector.Vector{
				{random.FRandom(-50, 50), random.FRandom(-50, 50)},
				{random.FRandom(-50, 50), random.FRandom(-50, 50)},
				{random.FRandom(-50, 50), random.FRandom(-50, 50)},
				{random.FRandom(-50, 50), random.FRandom(-50, 50)},
				{random.FRandom(-50, 50), random.FRandom(-50, 50)},
				{random.FRandom(-50, 50), random.FRandom(-50, 50)},
				{random.FRandom(-50, 50), random.FRandom(-50, 50)},
				{random.FRandom(-50, 50), random.FRandom(-50, 50)},
				{random.FRandom(-50, 50), random.FRandom(-50, 50)},
				{random.FRandom(-50, 50), random.FRandom(-50, 50)},
			},
			position,
			rotation,
			velocity)
	}

	return b
}

func createBullet(position vector.Vector, velocity vector.Vector) *body.Body {
	b := body.New()
	b.SetMass(1)
	b.SetShape(bullet.New())
	b.SetPosition(position)
	b.Velocity = velocity

	return b
}

func createCircle(radius float64, position vector.Vector, velocity vector.Vector) *body.Body {
	b := body.New()
	b.SetMass(1)
	b.SetShape(&circle.Circle{radius})
	b.SetPosition(position)
	b.Velocity = velocity

	return b
}

func createConvex(vertices []vector.Vector, position vector.Vector, rotation rotator.Rotator, velocity vector.Vector) *body.Body {
	b := body.New()
	b.SetMass(1)
	b.SetShape(convex.New(vertices))
	b.SetPosition(position)
	//b.SetRotation(rotation)
	b.Velocity = velocity

	return b
}

func isOutbound(position vector.Vector) bool {
	const MARGIN = 300
	if position.X < -WINDOW_WIDTH/2-MARGIN ||
		position.X > WINDOW_WIDTH/2+MARGIN ||
		position.Y < -WINDOW_HEIGHT/2-MARGIN ||
		position.Y > WINDOW_HEIGHT/2+MARGIN {
		return true
	}
	return false
}
