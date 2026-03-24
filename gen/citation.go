package gen

type Citation struct {
	GrampsObject `json:",inline"`
	Private      bool           `json:"private,omitzero"`
	Source       SourceHandle   `json:"source_handle,omitzero"`
	Page         string         `json:"page,omitzero"`
	Confidence   int            `json:"confidence"`
	Date         Date           `json:"date,omitzero"`
	Attributes   []AttributeRef `json:"attribute_list,omitempty"`
	Media        []MediaRef     `json:"media_list,omitempty"`
	Notes        []NoteHandle   `json:"note_list,omitempty"`
	Tags         []TagHandle    `json:"tag_list,omitempty"`
}
