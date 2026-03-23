package gen

type Person struct {
	GrampsObject
	Notes []NoteHandle `json:"note_list"`
}
