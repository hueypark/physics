package physics

import (
	"github.com/hueypark/physics/core/body"
	"github.com/hueypark/physics/core/contact"
	"github.com/hueypark/physics/core/vector"
)

type World struct {
	bodys    map[int64]*body.Body
	gravity  vector.Vector
	contscts []*contact.Contact
}

func New() World {
	return World{
		bodys:   make(map[int64]*body.Body),
		gravity: vector.Vector{0.0, -100.0}}
}

func (w *World) Tick(delta float64) {
	w.contscts = w.broadPhase()
	for _, c := range w.contscts {
		c.DetectCollision()
		c.SolveCollision()
	}

	for _, b := range w.bodys {
		b.AddForce(w.gravity)
		b.Tick(delta)
	}
}

func (w *World) Add(body *body.Body) {
	w.bodys[Context.IdGenerator.Generate()] = body
}

func (w *World) SetGravity(gravity vector.Vector) {
	w.gravity = gravity
}

func (w *World) Bodys() map[int64]*body.Body {
	return w.bodys
}

func (w *World) Manifolds() []*contact.Contact {
	return w.contscts
}

func (w *World) broadPhase() []*contact.Contact {
	contacts := []*contact.Contact{}

	for _, lhs := range w.bodys {
		for _, rhs := range w.bodys {
			if lhs.Id() <= rhs.Id() {
				continue
			}

			contacts = append(contacts, contact.New(lhs, rhs))
		}
	}

	return contacts
}
