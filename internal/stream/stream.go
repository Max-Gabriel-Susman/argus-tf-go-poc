package stream

type Stream struct {
	ID    string      `json:"id"`
	Label string      `json:"label"`
	Input interface{} `json:"input"`
}

func NewStream(id, label string, input interface{}) *Stream {
	return &Stream{
		ID:    id,
		Label: label,
		Input: input,
	}
}
