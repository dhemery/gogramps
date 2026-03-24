package gen

type AttributeTypeCode int
type AttributeType struct {
	Code   AttributeTypeCode `json:"value"`
	Custom string            `json:"string"`
}

type AttributeRef struct {
	Type  AttributeType `json:"type"`
	Value string        `json:"value"`
}

type EventRoleCode int
type EventRole struct {
	Code   EventRoleCode `json:"value"`
	Custom string        `json:"string"`
}

// An EventRef is the relationship between a person or family and an event.
type EventRef struct {
	Private    bool             `json:"private"`
	Event      EventHandle      `json:"ref"`
	Role       EventRole        `json:"role"`
	Attributes []AttributeRef   `json:"attribute_list"`
	Citations  []CitationHandle `json:"citation_list"`
	Notes      []NoteHandle     `json:"note_list"`
}

type MediaRef struct {
	Private   bool             `json:"private"`
	Media     MediaHandle      `json:"ref"`
	Citations []CitationHandle `json:"citation_list"`
	Notes     []NoteHandle     `json:"note_list"`
}
