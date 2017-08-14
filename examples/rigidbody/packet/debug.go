package packet

type DebugLine struct {
	Start Vector `json:"start"`
	End   Vector `json:"end"`
}

type DebugCircle struct {
	Position Vector  `json:"position"`
	Radius   float64 `json:"radius"`
}
