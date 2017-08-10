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
	r := body.New()
	r.SetMass(1)
	r.SetShape(&shape.Circle{1})
	world.Add(r)

	delta := time.Second / 30
	ticker := time.NewTicker(delta)
	go func() {
		for range ticker.C {
			world.Tick(delta.Seconds())
			for _, object := range world.GetBodys() {
				s.Broadcast(packet.Actor{
					object.Id(),
					packet.Vector{object.Position().X, object.Position().Y},
					packet.Circle(5),
				})
			}
		}
	}()

	s.Start(":8080")
}
