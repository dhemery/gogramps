package gen

type Event struct {
	GrampsObject `json:",inline"`
	Private      bool             `json:"private"`
	Description  string           `json:"description"`
	Place        PlaceHandle      `json:"place"`
	Citations    []CitationHandle `json:"citation_list"`
	Notes        []NoteHandle     `json:"note_list"`
}
