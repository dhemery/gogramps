package gen

type Note struct {
	GrampsObject `json:",inline"`
	Private      bool        `json:"private"`
	Format       int         `json:"format"`
	Tags         []TagHandle `json:"tag_list"`
}
