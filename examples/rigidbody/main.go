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
	r.Position = vector.Vector{0, 0, 0}
	r.Velocity = vector.Vector{-10, 0, 0}
	engine.Add(r)

	delta := time.Second / 30
	ticker := time.NewTicker(delta)
	go func() {
		for range ticker.C {
			engine.Tick(delta.Seconds())
			for _, object := range engine.GetObjects() {
				s.Broadcast(object)
			}
		}
	}()

	s.Start(":8080")
}
