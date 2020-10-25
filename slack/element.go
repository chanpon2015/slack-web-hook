package slack

// Element is
type Element struct {
	Type        string       `json:"type"`
	Text        *Text        `json:"text,omitempty"`
	Placeholder *Placeholder `json:"placeholder,omitempty"`
	URL         string       `json:"url,omitempty"`
	Value       string       `json:"value,omitempty"`
	Options     *[]Option    `json:"options,omitempty"`
}
