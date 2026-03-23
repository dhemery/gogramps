package gen

type Media struct {
	GrampsObject `json:",inline"`
	Private      bool             `json:"private"`
	Description  string           `json:"desc"`
	MimeType     string           `json:"mime"`
	Path         string           `json:"path"`
	Checksum     string           `json:"checksum"`
	Attributes   []AttributeRef   `json:"attribute_list"`
	Citations    []CitationHandle `json:"citation_list"`
	Notes        []NoteHandle     `json:"note_list"`
	Tags         []TagHandle      `json:"tag_list"`
}
