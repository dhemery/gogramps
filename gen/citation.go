package gen

type Citation struct {
	GrampsObject `json:",inline"`
	Private      bool         `json:"private"`
	Source       SourceHandle `json:"source_handle"`
	Page         string       `json:"page"`
	Confidence   int          `json:"confidence"`
	Notes        []NoteHandle `json:"note_list"`
	Tags         []TagHandle  `json:"tag_list"`
}
