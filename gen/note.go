package gen

type NoteTypeCode int

type NoteType struct {
	Code   NoteTypeCode `json:"value"`
	Custom string       `json:"string"`
}

type Note struct {
	GrampsObject `json:",inline"`
	Private      bool        `json:"private"`
	Format       int         `json:"format"`
	Type         NoteType    `json:"type"`
	Tags         []TagHandle `json:"tag_list"`
}
