package main

import (
	"time"

	"github.com/hueypark/physics/core"
	"github.com/hueypark/physics/core/body"
	"github.com/hueypark/physics/core/shape"
	"github.com/hueypark/physics/examples/rigidbody/packet"
	"github.com/hueypark/poseidon/core/server"
)

func main() {
	s := server.New()
	world := physics.New()

	delta := time.Second / 30
	ticker := time.NewTicker(delta)
	respawnTime := time.Duration(0)
	go func() {
		for range ticker.C {
			const RESPAWN_TIME = time.Second * 5
			respawnTime -= delta

			if respawnTime < 0 {
				respawnTime = RESPAWN_TIME
				world.Add(createCircle())
			}

			world.Tick(delta.Seconds())
			for _, body := range world.GetBodys() {
				s.Broadcast(packet.Actor{
					body.Id(),
					packet.Vector{body.Position().X, body.Position().Y},
					packet.Circle(body.Shape.(*shape.Circle).Radius),
				})
			}
		}
	}()

	s.Start(":8080")
}

func createCircle() *body.Body {
	b := body.New()
	b.SetMass(1)
	b.SetShape(&shape.Circle{10})

	return b
}
