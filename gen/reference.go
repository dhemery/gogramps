package gen

type AttributeType struct {
	// The index of the type in the standard attribute type enumeration,
	// or 0 if the attribute uses a custom type.
	Index int `json:"value"`
	// The custom attribute type, or "" if the attribute uses a standard
	// type.
	String string `json:"string"`
}

type AttributeRef struct {
	Type  AttributeType `json:"type"`
	Value string        `json:"value"`
}

type MediaRef struct {
	Private   bool             `json:"private"`
	Media     MediaHandle      `json:"ref"`
	Citations []CitationHandle `json:"citation_list"`
	Notes     []NoteHandle     `json:"note_list"`
}
