package gen

type Source struct {
	GrampsObject `json:",inline"`
	Private      bool           `json:"private,omitzero"`
	Title        string         `json:"title,omitzero"`
	Abbreviation string         `json:"abbreviation,omitzero"`
	Author       string         `json:"author,omitzero"`
	PubInfo      string         `json:"pub_info,omitzero"`
	Attributes   []AttributeRef `json:"attribute_list,omitempty"`
	Media        []MediaRef     `json:"media_list,omitempty"`
	Notes        []NoteHandle   `json:"note_list,omitempty"`
	Tags         []TagHandle    `json:"tag_list,omitempty"`
}
