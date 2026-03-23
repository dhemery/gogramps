package gen

type Family struct {
	GrampsObject `json:",inline"`
	Private      bool             `json:"private"`
	Complete     int              `json:"complete"`
	Mother       PersonHandle     `json:"mother_handle"`
	Father       PersonHandle     `json:"father_handle"`
	Media        []MediaRef       `json:"media_list"`
	Citations    []CitationHandle `json:"citation_list"`
	Notes        []NoteHandle     `json:"note_lis1t"`
}
