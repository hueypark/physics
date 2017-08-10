package physics

import (
	"github.com/hueypark/physics/core/body"
	"github.com/hueypark/physics/core/contact"
	"github.com/hueypark/physics/core/vector"
)

type World struct {
	bodys   map[int64]*body.Body
	gravity vector.Vector
}

type source struct {
	lhs, rhs *body.Body
}

func New() World {
	return World{make(map[int64]*body.Body), vector.Vector{0.0, -100.0}}
}

func (w *World) Tick(delta float64) {
	sources := w.generateSources()

	w.generateContacts(sources)

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

func (w *World) GetBodys() map[int64]*body.Body {
	return w.bodys
}

func (w *World) generateSources() []source {
	sources := []source{}

	for _, lhs := range w.bodys {
		for _, rhs := range w.bodys {
			if lhs.Id() <= rhs.Id() {
				continue
			}

			sources = append(sources, source{lhs, rhs})
		}
	}

	return sources
}

func (w *World) generateContacts(sources []source) []contact.Contact {
	contacts := []contact.Contact{}

	return contacts
}
