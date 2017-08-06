package rigidbody

import (
	"github.com/hueypark/physics/core/vector"
	"github.com/hueypark/poseidon/core"
)

type Rigidbody struct {
	Id       int64
	Position vector.Vector
	Velocity vector.Vector
}

func New() *Rigidbody {
	r := Rigidbody{}
	r.Id = core.Context.IdGenerator.Generate()

	return &r
}

func (r *Rigidbody) Tick(delta float64) {
	r.Position.AddScaledVector(r.Velocity, delta)
}
