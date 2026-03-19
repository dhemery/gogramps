// Package gen is a compiled representation of the family tree.
package gen

type Person struct {
	Primary
	Gender string     `json:"gender"`
	Name   PersonName `json:"name"`
}

type PersonName struct {
	Title   string `json:"title,omitempty"`
	First   string `json:"first,omitempty"`
	Surname string `json:"surname,omitempty"`
	Suffix  string `json:"suffix,omitempty"`
	Call    string `json:"call,omitempty"`
	Nick    string `json:"nick,omitempty"`
}
