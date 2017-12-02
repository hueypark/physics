package message

type ActorUpdate struct {
	Id  int64   `json:"id,string"`
	Pos Vector  `json:"pos"`
	Rot float64 `json:"rot"`
}

func (m ActorUpdate) Type() string {
	return ACTOR_UPDATE
}
