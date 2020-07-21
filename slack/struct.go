package slack

// Block is
type Block struct {
	Type     string     `json:"type"`
	Text     Text       `json:"text"`
	Elements *[]Element `json:"elements,omitempty"`
}

// CreateSection is
func (m *Message) CreateSection() *Block {
	b := Block{
		Type: "section",
		Text: Text{
			Type: "mrkdwn",
			Text: "*section*",
		},
	}
	m.Blocks = append(m.Blocks, b)
	return &b
}

// CreateActions is
func (m *Message) CreateActions() *Block {
	b := Block{
		Type:     "actions",
		Elements: &[]Element{},
	}
	m.Blocks = append(m.Blocks, b)
	return &b
}

// Element is
type Element struct {
	Type  string `json:"type"`
	Text  Text   `json:"text"`
	URL   string `json:"url,omitempty"`
	Value string `json:"value,omitempty"`
}

// Text is
type Text struct {
	Type  string `json:"type"`
	Text  string `json:"text"`
	Emoji bool   `json:"emoji,omitempty"`
}

// AddButton is
func (b *Block) AddButton(text, url, value string) *Block {
	element := Element{
		Type: "button",
		Text: Text{
			Type:  "plain_text",
			Text:  text,
			Emoji: false,
		},
		URL:   url,
		Value: value,
	}
	*b.Elements = append(*b.Elements, element)
	return b
}
