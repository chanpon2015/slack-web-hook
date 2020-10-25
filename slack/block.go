package slack

// Block is
type Block struct {
	Type      string     `json:"type"`
	Text      *Text      `json:"text,omitempty"`
	Elements  *[]Element `json:"elements,omitempty"`
	Accessory *Accessory `json:"accessory,omitempty"`
}

// AddDivider is
func (m *Message) AddDivider() *Block {
	b := Block{
		Type: "divider",
	}
	m.Blocks = append(m.Blocks, b)
	return &b
}
