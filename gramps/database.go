// Package gramps implements Gramps data structures.
package gramps

import (
	"encoding/xml"
)

type Created struct {
	Date    string `xml:"date,attr"`
	Version string `xml:"version,attr"`
}

type Researcher struct {
	XMLName    xml.Name   `xml:"researcher"`
}

type Header struct {
	XMLName    xml.Name   `xml:"header"`
	Created    Created    `xml:"created"`
	Researcher Researcher `xml:"researcher"`
	MediaPath  string     `xml:"mediapath"`
}

type Database struct {
	XMLName      xml.Name `xml:"database"`
	Header       Header
	Tags         any `xml:"tags"`
	Events       any `xml:"events"`
	People       any `xml:"people"`
	Families     any `xml:"families"`
	Citations    any `xml:"citations"`
	Sources      any `xml:"sources"`
	Places       any `xml:"places"`
	Objects      any `xml:"objects"`
	Repositories any `xml:"repositories"`
	Notes        any `xml:"notes"`
}
