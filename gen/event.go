package gen

type Event struct {
	GrampsObject `json:",inline"`
	Private      bool             `json:"private"`
	Description  string           `json:"description"`
	Media        []MediaRef       `json:"media_list"`
	Place        PlaceHandle      `json:"place"`
	Citations    []CitationHandle `json:"citation_list"`
	Notes        []NoteHandle     `json:"note_list"`
}
