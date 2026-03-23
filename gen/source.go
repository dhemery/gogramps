package gen

type Source struct {
	GrampsObject `json:",inline"`
	Private      bool         `json:"private"`
	Title        string       `json:"title"`
	Abbreviation string       `json:"abbreviation"`
	Author       string       `json:"author"`
	PubInfo      string       `json:"pub_info"`
	Notes        []NoteHandle `json:"note_list"`
}
