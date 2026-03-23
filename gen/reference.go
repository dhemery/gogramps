package gen

type MediaRef struct {
	Private   bool             `json:"private"`
	Media     MediaHandle      `json:"ref"`
	Citations []CitationHandle `json:"citation_list"`
	Notes     []NoteHandle     `json:"note_list"`
}
