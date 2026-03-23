package gen

type Media struct {
	GrampsObject `json:",inline"`
	Private      bool         `json:"private"`
	Description  string       `json:"desc"`
	MimeType     string       `json:"mime"`
	Path         string       `json:"path"`
	Checksum     string       `json:"checksum"`
	Notes        []NoteHandle `json:"note_list"`
}
