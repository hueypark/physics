package mainfold

import (
	"math"

	"github.com/hueypark/physics/core/body"
	"github.com/hueypark/physics/core/shape"
	"github.com/hueypark/physics/core/vector"
)

const RESTITUTION = 0.5

type Manifold struct {
	lhs         *body.Body
	rhs         *body.Body
	normal      vector.Vector // lhs to rhs
	penetration float64
	contacts    []vector.Vector
}

func New(lhs, rhs *body.Body) *Manifold {
	return &Manifold{lhs: lhs, rhs: rhs}
}

func (m *Manifold) DetectCollision() {
	lhsCircle := m.lhs.Shape.(*shape.Circle)
	rhsCircle := m.rhs.Shape.(*shape.Circle)

	m.normal = vector.Subtract(m.rhs.Position(), m.lhs.Position())

	distanceSquared := m.normal.SizeSquared()
	radius := lhsCircle.Radius + rhsCircle.Radius

	if distanceSquared >= radius*radius {
		return
	}

	distance := math.Sqrt(distanceSquared)

	m.penetration = radius - distance
	m.normal.Normalize()
	m.contacts = append(m.contacts, vector.Add(vector.Multiply(m.normal, lhsCircle.Radius), m.lhs.Position()))
}

func (m *Manifold) SolveCollision() {
	m.addImpulse()
}

func (m *Manifold) Contacts() []vector.Vector {
	return m.contacts
}

func (m *Manifold) Normal() vector.Vector {
	return m.normal
}

func (m *Manifold) Penetration() float64 {
	return m.penetration
}

func (m *Manifold) addImpulse() {
	for range m.contacts {
		relativeVelocity := vector.Subtract(m.rhs.Velocity, m.lhs.Velocity)

		velAlongNormal := vector.Dot(relativeVelocity, m.normal)
		if velAlongNormal > 0 {
			return
		}

		contactVelocity := velAlongNormal * -(1 + RESTITUTION)

		inverseMassSum := m.lhs.InverseMass() + m.rhs.InverseMass()

		impulse := vector.Multiply(m.normal, contactVelocity)
		impulse.Multiply(1 / inverseMassSum)

		m.lhs.AddImpluse(vector.Multiply(impulse, -1))
		m.rhs.AddImpluse(impulse)
	}
}
