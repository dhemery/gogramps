package gen

type Repository struct {
	GrampsObject `json:",inline"`
	Private      bool         `json:"private"`
	Name         string       `json:"name"`
	Notes        []NoteHandle `json:"note_list"`
	Tags         []TagHandle  `json:"tag_list"`
}
