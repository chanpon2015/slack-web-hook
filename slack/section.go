package slack

// SectionOption is section option.
type SectionOption func(*SectionConfig) error

// SectionConfig is
type SectionConfig struct {
	TextType  string
	Accessory *Accessory
}

// WithSetPlainText is text type plain_text.
func WithSetPlainText() SectionOption {
	return func(c *SectionConfig) error {
		c.TextType = "plain_text"
		return nil
	}
}

// WithSetMrkdwnText is text type mrkdwn.
func WithSetMrkdwnText() SectionOption {
	return func(c *SectionConfig) error {
		c.TextType = "mrkdwn"
		return nil
	}
}

// WithSetImageAccessory is image accessory.
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

// Accessory is accessory.
type Accessory struct {
	Type     string `json:"type"`
	ImageURL string `json:"image_url"`
	AltText  string `json:"alt_text"`
}

// CreateSection is selection.
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
