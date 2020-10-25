package slack

// AddButton is
func (b *Block) AddButton(text, url, value string) *Block {
	element := Element{
		Type: "button",
		Text: &Text{
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

// AddChannelsSelect is
func (b *Block) AddChannelsSelect() *Block {
	element := Element{
		Type: "channels_select",
	}
	*b.Elements = append(*b.Elements, element)
	return b
}

// AddUsersSelect is
func (b *Block) AddUsersSelect() *Block {
	element := Element{
		Type: "users_select",
		Placeholder: &Placeholder{
			Type:  "plain_text",
			Emoji: true,
		},
	}
	*b.Elements = append(*b.Elements, element)
	return b
}

// AddStaticSelect is
func (b *Block) AddStaticSelect(options []Option) *Block {
	element := Element{
		Type:    "static_select",
		Options: &options,
	}
	*b.Elements = append(*b.Elements, element)
	return b
}
