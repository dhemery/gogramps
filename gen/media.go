package gen

type Media struct {
	GrampsObject `json:",inline"`
	Private      bool             `json:"private,omitzero"`
	Date         DateValue             `json:"date,omitzero"`
	Description  string           `json:"desc,omitzero"`
	MimeType     string           `json:"mime,omitzero"`
	Path         string           `json:"path"`
	Checksum     string           `json:"checksum"`
	Attributes   []AttributeRef   `json:"attribute_list,omitempty"`
	Citations    []CitationHandle `json:"citation_list,omitempty"`
	Notes        []NoteHandle     `json:"note_list,omitempty"`
	Tags         []TagHandle      `json:"tag_list,omitempty"`
}
