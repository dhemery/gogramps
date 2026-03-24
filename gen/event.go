package gen

type EventCode int

type EventType struct {
	Code   EventCode `json:"value,omitzero"`
	Custom string    `json:"string,omitzero"`
}

type Event struct {
	GrampsObject `json:",inline"`
	Private      bool             `json:"private,omitzero"`
	Type         EventType        `json:"type,omitzero"`
	Description  string           `json:"description,omitzero"`
	Date         Date             `json:"date,omitzero"`
	Place        PlaceHandle      `json:"place,omitzero"`
	Attributes   []AttributeRef   `json:"attribute_list,omitempty"`
	Citations    []CitationHandle `json:"citation_list,omitempty"`
	Media        []MediaRef       `json:"media_list,omitempty"`
	Notes        []NoteHandle     `json:"note_list,omitempty"`
	Tags         []TagHandle      `json:"tag_list,omitempty"`
}
