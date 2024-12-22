package stream

type Source struct {
	ID    string      `json:"id"`
	Label string      `json:"label"`
	Input interface{} `json:"input"`
}

func NewStream(id, label string, input interface{}) *Source {
	return &Source{
		ID:    id,
		Label: label,
		Input: input,
	}
}
