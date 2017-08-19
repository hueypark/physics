package shape

type Circle struct {
	Radius float64
}

func (c *Circle) Type() int64 {
	return CIRCLE
}
