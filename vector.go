package vector

import "math"

type Vector struct {
	X float64
	Y float64
	Z float64
}

func (v *Vector) Invert() {
	v.X = -v.X
	v.Y = -v.Y
	v.Z = -v.Z
}

func (v *Vector) Magnitude() float64 {
	return math.Sqrt(v.SquareMagnitude())
}

func (v *Vector) SquareMagnitude() float64 {
	return (v.X * v.X) + (v.Y * v.Y) + (v.Z * v.Z)
}

func (v *Vector) Normalize() {
	l := v.Magnitude()
	if l > 0 {
		v.X /= l
		v.Y /= l
		v.Z /= l
	}
}

func (v *Vector) Multiply(val float64) {
	v.X *= val
	v.Y *= val
	v.Z *= val
}

func (v *Vector) Add(o Vector) {
	v.X += o.X
	v.Y += o.Y
	v.Z += o.Z
}

func (v *Vector) Subtract(o Vector) {
	v.X -= o.X
	v.Y -= o.Y
	v.Z -= o.Z
}

func Product(lhs, rhs Vector) Vector {
	return Vector{lhs.X * rhs.X, lhs.Y * rhs.Y, lhs.Z * rhs.Z}
}

func Dot(lhs, rhs Vector) float64 {
	return (lhs.X * rhs.X) + (lhs.Y * rhs.Y) + (lhs.Z * rhs.Z)
}

func Cross(lhs, rhs Vector) Vector {
	return Vector{
		(lhs.Y * rhs.Z) - (lhs.Z * rhs.Y),
		(lhs.Z * rhs.X) - (lhs.X * rhs.Z),
		(lhs.X * rhs.Y) - (lhs.Y * rhs.X)}
}
