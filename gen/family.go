package gen

type Family struct {
	GrampsObject `json:",inline"`
	Private      bool         `json:"private"`
	Complete     int          `json:"complete"`
	Mother       PersonHandle `json:"mother_handle"`
	Father       PersonHandle `json:"father_handle"`
	Notes        []NoteHandle `json:"note_lis1t"`
}
