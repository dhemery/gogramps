package gen

type Place struct {
	GrampsObject `json:",inline"`
	Private      bool             `json:"private"`
	Title        string           `json:"title"`
	Code         string           `json:"code"`
	Latitude     string           `json:"lat"`
	Longitude    string           `json:"long"`
	Citations    []CitationHandle `json:"citation_list"`
	Media        []MediaRef       `json:"media_list"`
	Notes        []NoteHandle     `json:"note_list"`
	Tags         []TagHandle      `json:"tag_list"`
}
