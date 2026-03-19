// Package gramps unmarshals Gramps data structures from XML.
package gramps

import (
	"encoding/xml"
)

// A TableObject is an object stored in its own record in a Gramps table.
type TableObject struct {
	// Identifes the record in the table.
	Handle string `xml:"handle,attr"`
	// The date and time of the latest update,
	// represented as a number of seconds since the Unix epoch.
	Change uint64 `xml:"change,attr"`
}

// A PrimaryObject represents one of Gramps's fundamental genealogical record types.
type PrimaryObject struct {
	TableObject
	Privacy
	Tags
	ID string `xml:"id,attr"`
}

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
	People       []Person     `xml:"people>person"`
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

