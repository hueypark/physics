package main

import (
	"time"

	"github.com/toqueteos/webbrowser"

	"github.com/hueypark/physics/core/body"
	"github.com/hueypark/physics/core/shape"
	"github.com/hueypark/physics/core/shape/circle"
	"github.com/hueypark/physics/core/shape/convex"
	"github.com/hueypark/physics/example/example_body"
	"github.com/hueypark/physics/example/message"
	"github.com/hueypark/poseidon/core/server"
)

func main() {
	s := server.New()

	example := example_body.New()
	example.SetOnBodyCreated(func(b body.Body) {
		broadcastActorCreate(s, b)
	})

	s.SetOnConnect(func(id int64) {
		world := example.World()
		for _, b := range world.Bodys() {
			broadcastActorCreate(s, *b)
		}
	})

	go func() {
		delta := time.Second / 30
		ticker := time.NewTicker(delta)
		for range ticker.C {
			example.Tick(delta.Seconds())

			world := example.World()
			for _, b := range world.Bodys() {
				s.Broadcast(message.ActorUpdate{
					Id:  b.Id(),
					Pos: message.Vector{X: b.Position().X, Y: b.Position().Y},
					Rot: b.Rotation().Degrees})
			}
		}
	}()

	url := "127.0.0.1:8080"
	webbrowser.Open("http://" + url)
	s.Start(url)
}

func broadcastActorCreate(s *server.Server, b body.Body) {
	s.Broadcast(message.ActorCreate{
		Id:  b.Id(),
		Pos: message.Vector{X: b.Position().X, Y: b.Position().Y},
		Rot: b.Rotation().Degrees})
	switch b.Shape.Type() {
	case shape.BULLET:
		s.Broadcast(message.ActorUpdateShapeCircle{Id: b.Id(), Radius: 2})
	case shape.CIRCLE:
		c := b.Shape.(*circle.Circle)
		s.Broadcast(message.ActorUpdateShapeCircle{Id: b.Id(), Radius: c.Radius})
	case shape.CONVEX:
		c := b.Shape.(*convex.Convex)
		var points []message.Vector
		for _, vertice := range c.Hull() {
			points = append(points, message.Vector{X: vertice.X, Y: vertice.Y})
		}
		s.Broadcast(message.ActorUpdateShapeConvex{Id: b.Id(), Points: points})
	}
}
