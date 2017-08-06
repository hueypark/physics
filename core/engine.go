package physics

type Actor interface {
	Tick(delta float64)
}

type Engine struct {
	actors map[int64]Actor
}

func New() Engine {
	return Engine{make(map[int64]Actor)}
}

func (e *Engine) Tick(delta float64) {
	for _, object := range e.actors {
		object.Tick(delta)
	}
}

func (e *Engine) Add(a Actor) {
	e.actors[Context.IdGenerator.Generate()] = a
}

func (e *Engine) GetObjects() map[int64]Actor {
	return e.actors
}
