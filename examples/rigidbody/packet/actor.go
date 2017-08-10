package packet

const (
	SHAPE_CIRCLE = "circle"
)

type Vector struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

func Circle(radius float64) circle {
	return circle{SHAPE_CIRCLE, radius}
}

type Actor struct {
	Id       int64       `json:"id,string"`
	Position Vector      `json:"position"`
	Shape    interface{} `json:"shape"`
}

type DeleteActor struct {
	Id int64 `json:"id,string"`
}

type circle struct {
	Type   string  `json:"type"`
	Radius float64 `json:"radius"`
}
