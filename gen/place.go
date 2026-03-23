package gen

type Place struct {
	GrampsObject `json:",inline"`
	Private      bool         `json:"private"`
	Title        string       `json:"title"`
	Code         string       `json:"code"`
	Latitude     string       `json:"lat"`
	Longitude    string       `json:"long"`
	Notes        []NoteHandle `json:"note_list"`
}
