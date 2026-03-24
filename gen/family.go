package gen

type FamilyTypeCode int

const (
	FamMarried = FamilyTypeCode(iota)
	FamUnmarried
	FamCivilUnion
	FamUnknown
	FamCustom
)

type FamilyType struct {
	Code   FamilyTypeCode `json:"value,omitzero"`
	Custom string         `json:"string,omitzero"`
}

type Family struct {
	GrampsObject `json:",inline"`
	Private      bool             `json:"private,omitzero"`
	Complete     int              `json:"complete,omitzero"`
	Type         FamilyType       `json:"type,omitzero"`
	Mother       PersonHandle     `json:"mother_handle,omitzero"`
	Father       PersonHandle     `json:"father_handle,omitzero"`
	Children     []ChildRef       `json:"child_ref_list,omitempty"`
	Attributes   []AttributeRef   `json:"attribute_list,omitempty"`
	Citations    []CitationHandle `json:"citation_list,omitempty"`
	Events       []EventRef       `json:"event_ref_list,omitempty"`
	Media        []MediaRef       `json:"media_list,omitempty"`
	Notes        []NoteHandle     `json:"note_list,omitempty"`
	Tags         []TagHandle      `json:"tag_list,omitempty"`
}
