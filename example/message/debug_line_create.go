package message

type DebugLineCreate struct {
	Start Vector `json:"start"`
	End   Vector `json:"end"`
}

func (m DebugLineCreate) Type() string {
	return DEBUG_LINE_CREATE
}
