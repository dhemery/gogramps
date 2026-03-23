package gen

type Person struct {
	GrampsObject
	Citations        []CitationHandle `json:"citation_list"`
	FamiliesAsParent []FamilyHandle   `json:"family_list"`
	FamiliesAsChild  []FamilyHandle   `json:"parent_family_list"`
	Notes            []NoteHandle     `json:"note_list"`
}
