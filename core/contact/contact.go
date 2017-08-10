package contact

import (
	"github.com/hueypark/physics/core/body"
	"github.com/hueypark/physics/core/vector"
)

type Contact struct {
	lhs      body.Body
	rhs      body.Body
	normal   vector.Vector // lhs to rhs
	contacts []vector.Vector
}
