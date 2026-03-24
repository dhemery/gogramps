package gen

type RepositoryTypeCode = int
type RepositoryType struct {
	Code   RepositoryTypeCode `json:"value,omitzero"`
	Custom string             `json:"string,omitzero"`
}

type Repository struct {
	GrampsObject `json:",inline"`
	Private      bool           `json:"private,omitzero"`
	Name         string         `json:"name,omitzero"`
	Type         RepositoryType `json:"type,omitzero"`
	Notes        []NoteHandle   `json:"note_list,omitempty"`
	Tags         []TagHandle    `json:"tag_list,omitempty"`
}
