package gen

type AttributeTypeCode int
type AttributeType struct {
	Code   AttributeTypeCode `json:"value,omitzero"`
	Custom string            `json:"string,omitzero"`
}

type AttributeRef struct {
	Type  AttributeType `json:"type,omitzero"`
	Value string        `json:"value,omitzero"`
}

type ParentRelCode int
type ParentRel struct {
	Code   ParentRelCode `json:"value,omitzero"`
	Custom string        `json:"string,omitzero"`
}

// A ChildRef describes how the father and mother in a family relate to the
// referred child.
type ChildRef struct {
	Private   bool             `json:"private,omitzero"`
	Child     PersonHandle     `json:"ref,omitzero"`
	FatherRel ParentRel        `json:"frel,omitzero"`
	MotherRel ParentRel        `json:"mrel,omitzero"`
	Citations []CitationHandle `json:"citation_list,omitempty"`
	Notes     []NoteHandle     `json:"note_list,omitempty"`
}

type EventRoleCode int
type EventRole struct {
	Code   EventRoleCode `json:"value,omitzero"`
	Custom string        `json:"string,omitzero"`
}

// An EventRef is the relationship between a person or family and an event.
type EventRef struct {
	Private    bool             `json:"private,omitzero"`
	Event      EventHandle      `json:"ref,omitzero"`
	Role       EventRole        `json:"role,omitzero"`
	Attributes []AttributeRef   `json:"attribute_list,omitempty"`
	Citations  []CitationHandle `json:"citation_list,omitempty"`
	Notes      []NoteHandle     `json:"note_list,omitempty"`
}

type MediaRef struct {
	Private   bool             `json:"private,omitzero"`
	Media     MediaHandle      `json:"ref,omitzero"`
	Citations []CitationHandle `json:"citation_list,omitempty"`
	Notes     []NoteHandle     `json:"note_list,omitempty"`
}
