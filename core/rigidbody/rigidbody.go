package rigidbody

import (
	"github.com/hueypark/physics/core/vector"
	"github.com/hueypark/poseidon/core"
)

type Rigidbody struct {
	id          int64
	position    vector.Vector
	Velocity    vector.Vector
	Shape       shape
	inverseMass float64
	forceSum    vector.Vector
}

type shape interface {
	Type() int64
}

func New() *Rigidbody {
	r := Rigidbody{}
	r.id = core.Context.IdGenerator.Generate()

	return &r
}

func (r *Rigidbody) Id() int64 {
	return r.id
}

func (r *Rigidbody) Position() vector.Vector {
	return r.position
}

func (r *Rigidbody) Tick(delta float64) {
	if r.inverseMass <= 0.0 {
		return
	}

	r.position.AddScaledVector(r.Velocity, delta)

	acceleration := vector.Vector{}
	acceleration.AddScaledVector(r.forceSum, r.inverseMass)

	r.Velocity.AddScaledVector(acceleration, delta)

	r.forceSum.Clear()
}

func (r *Rigidbody) SetMass(mass float64) {
	r.inverseMass = 1.0 / mass
}

func (r *Rigidbody) SetShape(s shape) {
	r.Shape = s
}

func (r *Rigidbody) AddForce(force vector.Vector) {
	r.forceSum.Add(force)
}
