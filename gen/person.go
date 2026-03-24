package gen

type Person struct {
	GrampsObject     `json:",inline"`
	Name             PersonName       `json:"primary_name"`
	AlternateNames   []PersonName     `json:"alternate_names"`
	Gender           int              `json:"gender"`
	FamiliesAsParent []FamilyHandle   `json:"family_list,omitempty"`
	FamiliesAsChild  []FamilyHandle   `json:"parent_family_list,omitempty"`
	Attributes       []AttributeRef   `json:"attribute_list,omitempty"`
	Citations        []CitationHandle `json:"citation_list,omitempty"`
	Events           []EventRef       `json:"event_ref_list,omitempty"`
	Media            []MediaRef       `json:"media_list,omitempty"`
	Notes            []NoteHandle     `json:"note_list,omitempty"`
	Tags             []TagHandle      `json:"tag_list,omitempty"`
}

type PersonName struct {
	Private   bool             `json:"private"`
	Title     string           `json:"title"`
	Given     string           `json:"first_name"`
	Surnames  []Surname        `json:"surname_list"`
	Date      Date             `json:"date"`
	Call      string           `json:"call"`
	FamNick   string           `json:"famnick"`
	Nick      string           `json:"nick"`
	DisplayAs int              `json:"display_as"`
	GroupAs   string           `json:"group_as"`
	SortAs    int              `json:"sort_as"`
	Citations []CitationHandle `json:"citation_list"`
	Notes     []Note           `json:"note_list"`
}

type Surname struct {
	Connector string `json:"connector"`
	Prefix    string `json:"prefix"`
	Surname   string `json:"surname"`
}
