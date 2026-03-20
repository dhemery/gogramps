// Package gramps unmarshals Gramps data structures from XML.
package gramps

import (
	"encoding/xml"
)

type DB struct {
	XMLName xml.Name `xml:"database"`
	Unknown
	Header
	Tags         []Tag        `xml:"tags>tag"`
	Citations    []Citation   `xml:"citations>citation"`
	Events       []Event      `xml:"events>event"`
	Families     []Family     `xml:"families>family"`
	Media        []Media      `xml:"objects>object"`
	Notes        []Note       `xml:"notes>note"`
	People       People       `xml:"people"`
	Places       []Place      `xml:"places>placeobj"`
	Repositories []Repository `xml:"repositories>repository"`
	Sources      []Source     `xml:"sources>source"`
}

type Header struct {
	XMLName    xml.Name   `xml:"header"`
	Created    Created    `xml:"created"`
	Researcher Researcher `xml:"researcher"`
	MediaPath  string     `xml:"mediapath"`
}

type Created struct {
	Date    string `xml:"date,attr"`
	Version string `xml:"version,attr"`
}

type Researcher struct {
	XMLName xml.Name `xml:"researcher"`
}

type People struct {
	Unknown
	People     []Person `xml:"person"`

	HomePerson string   `xml:"home,attr"`
}

func (p *People) hasUnknowns() bool {
	return p.Unknown.hasUnknowns()
}

func (db *DB) CollectUnknowns() []any {
	var unknowns []any

	if db.People.hasUnknowns() {
		people := db.People
		people.People = nil
		unknowns = append(unknowns, people)
	}

	for _, p := range db.People.People {
		if p.hasUnknowns() {
			unknowns = append(unknowns, p)
		}
	}

	return unknowns
}
