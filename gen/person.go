package gen

type Person struct {
	GrampsObject
	Gender           int              `json:"gender"`
	FamiliesAsParent []FamilyHandle   `json:"family_list"`
	FamiliesAsChild  []FamilyHandle   `json:"parent_family_list"`
	Attributes       []AttributeRef   `json:"attribute_list"`
	Citations        []CitationHandle `json:"citation_list"`
	Media            []MediaRef       `json:"media_list"`
	Notes            []NoteHandle     `json:"note_list"`
}
