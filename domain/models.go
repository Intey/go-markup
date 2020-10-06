package domain

// Mark object
type Mark struct {
	ID       int64       `json:"id,omitempty"`
	Position interface{} `json:"position"`
	Entity   string      `json:"entity"`
}
