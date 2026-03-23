package gen

type Family struct {
	GrampsObject `json:",inline"`
	Private      bool             `json:"private"`
	Complete     int              `json:"complete"`
	Mother       PersonHandle     `json:"mother_handle"`
	Father       PersonHandle     `json:"father_handle"`
	Attributes   []AttributeRef   `json:"attribute_list"`
	Citations    []CitationHandle `json:"citation_list"`
	Media        []MediaRef       `json:"media_list"`
	Notes        []NoteHandle     `json:"note_lis1t"`
	Tags         []TagHandle      `json:"tag_list"`
}
