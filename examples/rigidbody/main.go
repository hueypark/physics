package main

import (
	"time"

	"github.com/hueypark/framework/core/random"
	"github.com/hueypark/physics/core"
	"github.com/hueypark/physics/core/body"
	"github.com/hueypark/physics/core/shape"
	"github.com/hueypark/physics/core/vector"
	"github.com/hueypark/physics/examples/rigidbody/packet"
	"github.com/hueypark/poseidon/core/server"
)

func main() {
	s := server.New()
	world := physics.New()
	plane := createCircle(100, vector.Vector{0, -200}, vector.Vector{})
	plane.SetStatic()
	world.Add(plane)

	delta := time.Second / 30
	ticker := time.NewTicker(delta)
	respawnTime := time.Duration(0)
	go func() {
		for range ticker.C {
			const RESPAWN_TIME = time.Millisecond * 100
			respawnTime -= delta

			if respawnTime < 0 {
				respawnTime = RESPAWN_TIME
				world.Add(createCircle(random.FRandom(5, 20), vector.Vector{random.FRandom(-300, 300), 0}, vector.Vector{0, random.FRandom(100, 300)}))
			}

			world.Tick(delta.Seconds())
			for _, m := range world.Manifolds() {
				for _, c := range m.Contacts() {
					start := vector.Add(c, vector.Multiply(m.Normal(), -10))
					end := vector.Add(c, vector.Multiply(m.Normal(), 10))

					s.Broadcast(packet.DebugLine{
						packet.Vector{start.X, start.Y},
						packet.Vector{end.X, end.Y},
					})
				}
			}

			for _, b := range world.Bodys() {
				s.Broadcast(packet.Actor{
					b.Id(),
					packet.Vector{b.Position().X, b.Position().Y},
					packet.Circle(b.Shape.(*shape.Circle).Radius),
				})
			}
		}
	}()

	s.Start(":8080")
}

func createCircle(radius float64, position vector.Vector, velocity vector.Vector) *body.Body {
	b := body.New()
	b.SetMass(1)
	b.SetShape(&shape.Circle{radius})
	b.SetPosition(position)
	b.Velocity = velocity

	return b
}
