package message

type ActorCreate struct {
	Id  int64   `json:"id,string"`
	Pos Vector  `json:"pos"`
	Rot float64 `json:"rot"`
}

func (a ActorCreate) Type() string {
	return ACTOR_CREATE
}
