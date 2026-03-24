package gen

type Person struct {
	GrampsObject     `json:",inline"`
	Gender           int              `json:"gender"`
	FamiliesAsParent []FamilyHandle   `json:"family_list,omitempty"`
	FamiliesAsChild  []FamilyHandle   `json:"parent_family_list,omitempty"`
	Attributes       []AttributeRef   `json:"attribute_list,omitempty"`
	Citations        []CitationHandle `json:"citation_list,omitempty"`
	Events           []EventRef       `json:"event_ref_list,omitempty"`
	Media            []MediaRef       `json:"media_list,omitempty"`
	Notes            []NoteHandle     `json:"note_list,omitempty"`
	Tags             []TagHandle      `json:"tag_list,omitempty"`
}
