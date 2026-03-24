package gen

type EventCode int

type EventType struct {
	Code   EventCode `json:"value"`
	Custom string    `json:"string"`
}

type Event struct {
	GrampsObject `json:",inline"`
	Private      bool             `json:"private"`
	Type         EventType        `json:"type"`
	Description  string           `json:"description"`
	Date         Date             `json:"date"`
	Place        PlaceHandle      `json:"place"`
	Attributes   []AttributeRef   `json:"attribute_list"`
	Citations    []CitationHandle `json:"citation_list"`
	Media        []MediaRef       `json:"media_list"`
	Notes        []NoteHandle     `json:"note_list"`
	Tags         []TagHandle      `json:"tag_list"`
}
