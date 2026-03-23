package gen

type Person struct {
	GrampsObject
	Citations    []CitationHandle `json:"citation_list"`
	Notes []NoteHandle `json:"note_list"`
}
