package slack

// Accessory is
type Accessory struct {
	Type     string `json:"type"`
	ImageURL string `json:"image_url"`
	AltText  string `json:"alt_text"`
}

// SectionOption is
type SectionOption func(*SectionConfig) error

// SectionConfig is
type SectionConfig struct {
	TextType  string
	Accessory *Accessory
}

// WithSetPlainText is
func WithSetPlainText() SectionOption {
	return func(c *SectionConfig) error {
		c.TextType = "plain_text"
		return nil
	}
}

// WithSetMrkdwnText is
func WithSetMrkdwnText() SectionOption {
	return func(c *SectionConfig) error {
		c.TextType = "mrkdwn"
		return nil
	}
}

// WithSetImageAccessory is
func WithSetImageAccessory(url, text string) SectionOption {
	return func(c *SectionConfig) error {
		c.Accessory = &Accessory{
			Type:     "image",
			ImageURL: url,
			AltText:  text,
		}
		return nil
	}
}

// Block is
type Block struct {
	Type      string     `json:"type"`
	Text      *Text      `json:"text,omitempty"`
	Elements  *[]Element `json:"elements,omitempty"`
	Accessory *Accessory `json:"accessory,omitempty"`
}

// CreateSection is
func (m *Message) CreateSection(opts ...SectionOption) *Block {
	c := SectionConfig{
		TextType: "plain_text",
	}
	for _, opt := range opts {
		if err := opt(&c); err != nil {
			return nil
		}
	}
	b := Block{
		Type: "section",
		Text: &Text{
			Type: c.TextType,
			Text: "*section*",
		},
		Accessory: c.Accessory,
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

// AddDivider is
func (m *Message) AddDivider() *Block {
	b := Block{
		Type: "divider",
	}
	m.Blocks = append(m.Blocks, b)
	return &b
}

// Element is
type Element struct {
	Type        string       `json:"type"`
	Text        *Text        `json:"text,omitempty"`
	Placeholder *Placeholder `json:"placeholder,omitempty"`
	URL         string       `json:"url,omitempty"`
	Value       string       `json:"value,omitempty"`
	Options     *[]Option    `json:"options,omitempty"`
}

// Text is
type Text struct {
	Type  string `json:"type"`
	Text  string `json:"text"`
	Emoji bool   `json:"emoji,omitempty"`
}

// Placeholder is
type Placeholder struct {
	Type  string `json:"type"`
	Text  string `json:"text"`
	Emoji bool   `json:"emoji"`
}

// Option is
type Option struct {
	Text  Text   `json:"text"`
	Value string `json:"value"`
}

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
