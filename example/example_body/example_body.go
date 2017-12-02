package example_body

import (
	"github.com/hueypark/framework/core/random"
	"github.com/hueypark/physics/core"
	"github.com/hueypark/physics/core/body"
	"github.com/hueypark/physics/core/math/rotator"
	"github.com/hueypark/physics/core/math/vector"
	"github.com/hueypark/physics/core/shape/bullet"
	"github.com/hueypark/physics/core/shape/circle"
	"github.com/hueypark/physics/core/shape/convex"
)

const WINDOW_WIDTH = 1024
const WINDOW_HEIGHT = 768

type ExampleBody struct {
	world         *physics.World
	respawnTime   float64
	onBodyCreated func(body body.Body)
}

func New() *ExampleBody {
	world := physics.New()
	plane := createConvex(
		[]vector.Vector{
			{X: 300, Y: -100},
			{X: 300, Y: -120},
			{X: -300, Y: -100},
			{X: -300, Y: -120},
		},
		vector.Vector{X: 0, Y: -200},
		rotator.ZERO(),
		vector.Vector{})
	plane.SetStatic()
	world.Add(plane)

	return &ExampleBody{world, 0, nil}
}

func (e *ExampleBody) SetOnBodyCreated(onBodyCreated func(body body.Body)) {
	e.onBodyCreated = onBodyCreated
}

func (e *ExampleBody) Tick(delta float64) {
	const RESPAWN_TIME = 0.5
	e.respawnTime -= delta

	if e.respawnTime < 0 {
		e.respawnTime = RESPAWN_TIME
		e.world.Add(
			e.createRandomShape(
				vector.Vector{X: random.FRandom(-300, 300), Y: 0},
				rotator.Rotator{Degrees: random.FRandom(180.0, 360.0)},
				vector.Vector{X: 0, Y: 0}))
	}

	e.world.Tick(delta)

	for _, b := range e.world.Bodys() {
		if isOutbound(b.Position()) {
			e.world.ReservedDelete(b.Id())
		}
	}
}

func (e *ExampleBody) World() *physics.World {
	return e.world
}

func (e *ExampleBody) createRandomShape(position vector.Vector, rotation rotator.Rotator, velocity vector.Vector) *body.Body {
	var b *body.Body
	switch random.Random(0, 2) {
	case 0:
		b = createBullet(position, velocity)
	case 1:
		b = createCircle(random.FRandom(10, 50), position, velocity)
	case 2:
		b = createConvex(
			[]vector.Vector{
				{X: random.FRandom(-50, 50), Y: random.FRandom(-50, 50)},
				{X: random.FRandom(-50, 50), Y: random.FRandom(-50, 50)},
				{X: random.FRandom(-50, 50), Y: random.FRandom(-50, 50)},
				{X: random.FRandom(-50, 50), Y: random.FRandom(-50, 50)},
				{X: random.FRandom(-50, 50), Y: random.FRandom(-50, 50)},
				{X: random.FRandom(-50, 50), Y: random.FRandom(-50, 50)},
				{X: random.FRandom(-50, 50), Y: random.FRandom(-50, 50)},
				{X: random.FRandom(-50, 50), Y: random.FRandom(-50, 50)},
				{X: random.FRandom(-50, 50), Y: random.FRandom(-50, 50)},
				{X: random.FRandom(-50, 50), Y: random.FRandom(-50, 50)},
			},
			position,
			rotation,
			velocity)
	}

	if e.onBodyCreated != nil {
		e.onBodyCreated(*b)
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
	b.SetShape(&circle.Circle{Radius: radius})
	b.SetPosition(position)
	b.Velocity = velocity

	return b
}

func createConvex(vertices []vector.Vector, position vector.Vector, rotation rotator.Rotator, velocity vector.Vector) *body.Body {
	b := body.New()
	b.SetMass(1)
	b.SetShape(convex.New(vertices))
	b.SetPosition(position)
	b.SetRotation(rotation)
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
