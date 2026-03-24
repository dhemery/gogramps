package gen

type PlaceTypeCode int

type PlaceType struct {
	Code   PlaceTypeCode `json:"value,omitzero"`
	Custom string        `json:"string,omitzero"`
}

type Place struct {
	GrampsObject `json:",inline"`
	Private      bool             `json:"private,omitzero"`
	Title        string           `json:"title,omitzero"`
	Name         PlaceName        `json:"name,omitzero"`
	Type         PlaceType        `json:"place_type,omitzero"`
	EnclosedBy   []PlaceRef       `json:"placeref_list,omitempty"`
	Code         string           `json:"code,omitzero"`
	Latitude     string           `json:"lat,omitzero"`
	Longitude    string           `json:"long,omitzero"`
	Citations    []CitationHandle `json:"citation_list,omitempty"`
	Media        []MediaRef       `json:"media_list,omitempty"`
	Notes        []NoteHandle     `json:"note_list,omitempty"`
	Tags         []TagHandle      `json:"tag_list,omitempty"`
}

type PlaceName struct {
	Name     string `json:"value,omitzero"`
	Date     DateValue   `json:"date,omitzero"`
	Language string `json:"language,omitzero"`
}
