package gen

type Person struct {
	GrampsObject     `json:",inline"`
	Name             PersonName       `json:"primary_name"`
	AlternateNames   []PersonName     `json:"alternate_names,omitempty"`
	Gender           int              `json:"gender"`
	FamiliesAsParent []FamilyHandle   `json:"family_list,omitempty"`
	FamiliesAsChild  []FamilyHandle   `json:"parent_family_list,omitempty"`
	Associations     []PersonRef      `json:"person_ref_list"`
	Attributes       []AttributeRef   `json:"attribute_list,omitempty"`
	Citations        []CitationHandle `json:"citation_list,omitempty"`
	Events           []EventRef       `json:"event_ref_list,omitempty"`
	Media            []MediaRef       `json:"media_list,omitempty"`
	Notes            []NoteHandle     `json:"note_list,omitempty"`
	Tags             []TagHandle      `json:"tag_list,omitempty"`
}

type PersonName struct {
	Private   bool             `json:"private,omitzero"`
	Title     string           `json:"title,omitzero"`
	Given     string           `json:"first_name,omitzero"`
	Surnames  []Surname        `json:"surname_list"`
	Date      DateValue        `json:"date,omitzero"`
	Call      string           `json:"call,omitzero"`
	FamNick   string           `json:"famnick,omitzero"`
	Nick      string           `json:"nick,omitzero"`
	DisplayAs int              `json:"display_as,omitzero"`
	GroupAs   string           `json:"group_as,omitzero"`
	SortAs    int              `json:"sort_as,omitzero"`
	Citations []CitationHandle `json:"citation_list,omitempty"`
	Notes     []Note           `json:"note_list,omitempty"`
}

type Surname struct {
	Connector  string     `json:"connector,omitzero"`
	Prefix     string     `json:"prefix,omitzero"`
	Surname    string     `json:"surname,omitzero"`
	OriginType OriginType `json:"origin_type,omitzero"`
}

type OriginTypeCode int

type OriginType struct {
	Code   OriginTypeCode `json:"value,omitzero"`
	Custom string         `json:"string,omitempty"`
}
