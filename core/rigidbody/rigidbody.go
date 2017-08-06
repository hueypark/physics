package rigidbody

import (
	"github.com/hueypark/physics/core/vector"
	"github.com/hueypark/poseidon/core"
)

type Rigidbody struct {
	Id          int64
	Position    vector.Vector
	Velocity    vector.Vector
	inverseMass float64
	forceSum    vector.Vector
}

func New() *Rigidbody {
	r := Rigidbody{}
	r.Id = core.Context.IdGenerator.Generate()

	return &r
}

func (r *Rigidbody) Tick(delta float64) {
	if r.inverseMass <= 0.0 {
		return
	}

	r.Position.AddScaledVector(r.Velocity, delta)

	acceleration := vector.Vector{}
	acceleration.AddScaledVector(r.forceSum, r.inverseMass)

	r.Velocity.AddScaledVector(acceleration, delta)

	r.forceSum.Clear()
}

func (r *Rigidbody) SetMass(mass float64) {
	r.inverseMass = 1.0 / mass
}

func (r *Rigidbody) AddForce(force vector.Vector) {
	r.forceSum.Add(force)
}
