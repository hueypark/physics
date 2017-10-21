package physics

import (
	"github.com/hueypark/physics/core/body"
	"github.com/hueypark/physics/core/contact"
	"github.com/hueypark/physics/core/context"
	"github.com/hueypark/physics/core/math/vector"
)

type World struct {
	bodys                 map[int64]*body.Body
	gravity               vector.Vector
	contscts              []*contact.Contact
	reservedDeleteBodyIds []int64
}

func New() World {
	return World{
		bodys:   make(map[int64]*body.Body),
		gravity: vector.Vector{0.0, -100.0}}
}

func (w *World) Tick(delta float64) {
	w.deleteReserveDeleteBodys()

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
	w.bodys[context.Context.IdGenerator.Generate()] = body
}

func (w *World) ReservedDelete(id int64) {
	w.reservedDeleteBodyIds = append(w.reservedDeleteBodyIds, id)
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

func (w *World) deleteReserveDeleteBodys() {
	for _, id := range w.reservedDeleteBodyIds {
		delete(w.bodys, id)
	}

	w.reservedDeleteBodyIds = []int64{}
}
