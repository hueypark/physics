package main

import (
	"time"

	"github.com/hueypark/physics/core"
	"github.com/hueypark/physics/core/rigidbody"
	"github.com/hueypark/physics/core/vector"
	"github.com/hueypark/poseidon/core/server"
)

func main() {
	s := server.New()
	engine := physics.New()
	r := rigidbody.New()
	r.SetMass(1)
	engine.Add(r)

	delta := time.Second / 30
	ticker := time.NewTicker(delta)
	go func() {
		for range ticker.C {
			r.AddForce(vector.Vector{0, 1, 0})
			engine.Tick(delta.Seconds())
			for _, object := range engine.GetObjects() {
				s.Broadcast(object)
			}
		}
	}()

	s.Start(":8080")
}
