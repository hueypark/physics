package bullet

import "github.com/hueypark/physics/core/shape"

type Bullet struct {
}

func New() *Bullet {
	return &Bullet{}
}

func (b *Bullet) Type() int64 {
	return shape.BULLET
}
