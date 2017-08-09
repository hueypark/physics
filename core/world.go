package physics

type Actor interface {
	Tick(delta float64)
}

type World struct {
	actors map[int64]Actor
}

func New() World {
	return World{make(map[int64]Actor)}
}

func (e *World) Tick(delta float64) {
	for _, object := range e.actors {
		object.Tick(delta)
	}
}

func (e *World) Add(a Actor) {
	e.actors[Context.IdGenerator.Generate()] = a
}

func (e *World) GetObjects() map[int64]Actor {
	return e.actors
}
