package message

type ActorUpdateShapeConvex struct {
	Id     int64    `json:"id,string"`
	Points []Vector `json:"points"`
}

func (m ActorUpdateShapeConvex) Type() string {
	return ACTOR_UPDATE_SHAPE_CONVEX
}
