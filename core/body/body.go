package body

import (
	"github.com/hueypark/physics/core/context"
	"github.com/hueypark/physics/core/math/rotator"
	"github.com/hueypark/physics/core/math/vector"
)

type Body struct {
	id              int64
	position        vector.Vector
	rotation        rotator.Rotator
	Velocity        vector.Vector
	angularVelocity float64
	Shape           shape
	inverseMass     float64
	inverseInertia  float64
	forceSum        vector.Vector
}

type shape interface {
	Type() int64
}

func New() *Body {
	r := Body{}
	r.id = context.Context.IdGenerator.Generate()

	return &r
}

func (r *Body) Id() int64 {
	return r.id
}

func (r *Body) Position() vector.Vector {
	return r.position
}

func (r *Body) SetPosition(position vector.Vector) {
	r.position = position
}

func (r *Body) Rotation() rotator.Rotator {
	return r.rotation
}

func (r *Body) SetRotation(rotation rotator.Rotator) {
	r.rotation = rotation
}

func (r *Body) SetAngularVelocity(degrees float64) {
	r.angularVelocity = degrees
}

func (r *Body) Tick(delta float64) {
	if r.inverseMass <= 0.0 {
		return
	}

	r.position.AddScaledVector(r.Velocity, delta)

	acceleration := vector.Vector{}
	acceleration.AddScaledVector(r.forceSum, r.inverseMass)

	r.Velocity.AddScaledVector(acceleration, delta)

	r.forceSum.Clear()

	r.rotation.AddScaled(r.angularVelocity, delta)
}

func (r *Body) SetMass(mass float64) {
	r.inverseMass = 1.0 / mass
	r.inverseInertia = 1.0 / mass
}

func (r *Body) InverseMass() float64 {
	return r.inverseMass
}

func (r *Body) SetStatic() {
	r.inverseMass = 0
}

func (r *Body) SetShape(s shape) {
	r.Shape = s
}

func (r *Body) AddForce(force vector.Vector) {
	r.forceSum.Add(force)
}

func (r *Body) AddImpluse(impulse vector.Vector) {
	r.Velocity.AddScaledVector(impulse, r.inverseMass)
}
