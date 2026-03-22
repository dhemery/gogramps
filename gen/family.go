package gen

type Family struct {
	GrampsObject `json:",inline"`
	Private      bool      `json:"private"`
	Complete     int       `json:"complete"`
	Mother       PersonRef `json:"mother_handle"`
	Father       PersonRef `json:"father_handle"`
}
