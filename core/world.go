package physics

import (
	"github.com/hueypark/physics/core/body"
	"github.com/hueypark/physics/core/contact"
)

type World struct {
	bodys map[int64]*body.Body
}

type contactSource struct {
	lhs body.Body
	rhs body.Body
}

func New() World {
	return World{make(map[int64]*body.Body)}
}

func (e *World) Tick(delta float64) {
	sources := e.generateSources()

	e.generateContacts(sources)

	for _, actor := range e.bodys {
		actor.Tick(delta)
	}
}

func (e *World) Add(body *body.Body) {
	e.bodys[Context.IdGenerator.Generate()] = body
}

func (e *World) GetObjects() map[int64]*body.Body {
	return e.bodys
}

func (e *World) generateSources() []contactSource {
	sources := []contactSource{}

	return sources
}

func (e *World) generateContacts([]contactSource) []contact.Contact {
	contacts := []contact.Contact{}

	return contacts
}
