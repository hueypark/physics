package message

type ActorUpdateShapeCircle struct {
	Id     int64   `json:"id,string"`
	Radius float64 `json:"radius"`
}

func (m ActorUpdateShapeCircle) Type() string {
	return ACTOR_UPDATE_SHAPE_CIRCLE
}
