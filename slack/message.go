package slack

// Message is
type Message struct {
	Text   string  `json:"text,omitempty"`
	Blocks []Block `json:"blocks"`
}
