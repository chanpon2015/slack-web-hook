package slack

// CreateActions is
func (m *Message) CreateActions() *Block {
	b := Block{
		Type:     "actions",
		Elements: &[]Element{},
	}
	m.Blocks = append(m.Blocks, b)
	return &b
}

// CreateActions is actions.
func (w *WebHook) CreateActions() *Block {
	b := Block{
		Type:     "actions",
		Elements: &[]Element{},
	}
	w.Msg.Blocks = append(w.Msg.Blocks, b)
	return &b
}
