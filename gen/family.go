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
	Code   FamilyTypeCode `json:"value"`
	Custom string         `json:"string"`
}

type Family struct {
	GrampsObject `json:",inline"`
	Private      bool             `json:"private"`
	Complete     int              `json:"complete"`
	Type         FamilyType       `json:"type"`
	Mother       PersonHandle     `json:"mother_handle"`
	Father       PersonHandle     `json:"father_handle"`
	Children     []ChildRef       `json:"child_ref_list"`
	Attributes   []AttributeRef   `json:"attribute_list"`
	Citations    []CitationHandle `json:"citation_list"`
	Events       []EventRef       `json:"event_ref_list"`
	Media        []MediaRef       `json:"media_list"`
	Notes        []NoteHandle     `json:"note_lis1t"`
	Tags         []TagHandle      `json:"tag_list"`
}
