package gen

type RepositoryTypeCode = int
type RepositoryType struct {
	Code   RepositoryTypeCode `json:"value"`
	Custom string             `json:"string"`
}

type Repository struct {
	GrampsObject `json:",inline"`
	Private      bool           `json:"private"`
	Name         string         `json:"name"`
	Type         RepositoryType `json:"type"`
	Notes        []NoteHandle   `json:"note_list"`
	Tags         []TagHandle    `json:"tag_list"`
}
