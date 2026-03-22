package gen

type Citation struct {
	GrampsObject `json:",inline"`
	Private      bool      `json:"private"`
	Source       SourceRef `json:"source_handle"`
	Page         string    `json:"page"`
	Confidence   int       `json:"confidence"`
	// Unknown      map[string]any `json:",unknown"`
}


