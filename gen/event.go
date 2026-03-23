package gen

type Event struct {
	GrampsObject `json:",inline"`
	Private      bool             `json:"private"`
	Description  string           `json:"description"`
	Place        PlaceHandle      `json:"place"`
	Attributes   []AttributeRef   `json:"attribute_list"`
	Citations    []CitationHandle `json:"citation_list"`
	Media        []MediaRef       `json:"media_list"`
	Notes        []NoteHandle     `json:"note_list"`
	Tags         []TagHandle      `json:"tag_list"`
}
