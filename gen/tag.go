package gen

type Tag struct {
	TableObject `json:",inline"`
	Color       string `json:"color"`
	Name        string `json:"name"`
	Priority    int    `json:"priority"`
}
