package gen

type PlaceTypeCode int

type PlaceType struct {
	Code   PlaceTypeCode `json:"value"`
	Custom string        `json:"string"`
}

type Place struct {
	GrampsObject `json:",inline"`
	Private      bool             `json:"private"`
	Type         PlaceType        `json:"place_type"`
	Title        string           `json:"title"`
	Code         string           `json:"code"`
	Latitude     string           `json:"lat"`
	Longitude    string           `json:"long"`
	Citations    []CitationHandle `json:"citation_list"`
	Media        []MediaRef       `json:"media_list"`
	Notes        []NoteHandle     `json:"note_list"`
	Tags         []TagHandle      `json:"tag_list"`
}
