// Package gramps implements Gramps data structures.
package gramps

import (
	"encoding/xml"
)

type Database struct {
	XMLName      xml.Name `xml:"database"`
	Header       Header
	Tags         []Tag        `xml:"tags>tag"`
	Events       []Event      `xml:"events>event"`
	People       []Person     `xml:"people>person"`
	Families     []Family     `xml:"families>family"`
	Citations    []Citation   `xml:"citations>citation"`
	Sources      []Source     `xml:"sources>source"`
	Places       []Place      `xml:"places>placeobj"`
	Media        []Media      `xml:"objects>object"`
	Repositories []Repository `xml:"repositories>repository"`
	Notes        []Note       `xml:"notes>note"`
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

type Tag struct {
	Handle   string `xml:"handle,attr"`
	Change   uint64 `xml:"change,attr"`
	Name     string `xml:"name,attr"`
	Color    string `xml:"color,attr"`
	Priority uint   `xml:"priority,attr"`
}

type Primary struct {
	Handle string `xml:"handle,attr"`
	Change uint64 `xml:"change,attr"`
	ID     string `xml:"id,attr"`
}

type Event struct {
	Primary
	Type        string      `xml:"type"`
	Date        DateVal     `xml:"dateval"`
	PlaceRef    PlaceRef    `xml:"place"`
	Description string      `xml:"description"`
	CitationRef CitationRef `xml:"citationref"`
}

type EventRef struct {
	Handle     string      `xml:"hlink,attr"`
	Role       string      `xml:"role,attr"`
	Attributes []Attribute `xml:"attribute"`
}

type DateVal struct {
	Val  string `xml:"val,attr"`
	Type string `xml:"type,attr"`
}

type Person struct {
	Primary
	Gender   string      `xml:"gender"`
	Name     PersonName  `xml:"name"`
	Events   []EventRef  `xml:"eventref"`
	ChildOf  []FamilyRef `xml:"childof"`
	ParentIn []FamilyRef `xml:"parentin"`
}

type PersonRef struct {
	Handle string `xml:"hlink,attr"`
}

type PersonName struct {
	First   string  `xml:"first"`
	Call    string  `xml:"call"`
	Nick    string  `xml:"nick"`
	Surname string  `xml:"surname"`
	Suffix  string  `xml:"suffix"`
	DateStr DateVal `xml:"datestr"`
}

type Attribute struct {
	Type  string `xml:"type,attr"`
	Value string `xml:"value,attr"`
}

type Family struct {
	Primary
	Rel      Rel         `xml:"rel"`
	Father   PersonRef   `xml:"father"`
	Mother   PersonRef   `xml:"mother"`
	Events   []EventRef  `xml:"eventref"`
	Children []PersonRef `xml:"childref"`
}

type FamilyRef struct {
	Handle string `xml:"hlink,attr"`
}

type Rel struct {
	Type string `xml:"type,attr"`
}

type Citation struct {
	Primary
	Date       DateVal     `xml:"datestr"`
	Page       string      `xml:"page"`
	Confidence uint8       `xml:"confidence"`
	Media      []MediaRef  `xml:"mediaref"`
	Sources    []SourceRef `xml:"sourceref"`
	Notes      []NoteRef   `xml:"noteref"`
}

type CitationRef struct {
	Handle string `xml:"hlink,attr"`
}

type Source struct {
	Primary
	Title        string          `xml:"stitle"`
	Author       string          `xml:"sauthor"`
	PubInfo      string          `xml:"spubinfo"`
	Media        []MediaRef      `xml:"mediaref"`
	Notes        []NoteRef       `xml:"noteref"`
	Repositories []RepositoryRef `xml:"reporef"`
}

type SourceRef struct {
	Handle string `xml:"hlink,attr"`
}

type Place struct {
	Primary
	Type          string      `xml:"type,attr"`
	Name          PlaceName   `xml:"pname"`
	Coordinates   Coordinates `xml:"coord"`
	EncompassedBy []PlaceRef  `xml:"placeref"`
}

type PlaceRef struct {
	Handle string `xml:"hlink,attr"`
	Date DateVal `xml:"dateval"`
}

type Coordinates struct {
	Longitude string `xml:"long,attr"`
	Latitude  string `xml:"lat,attr"`
}

type PlaceName struct {
	Value string `xml:"value,attr"`
}

type Media struct {
	Primary
}

type MediaRef struct {
	Handle string `xml:"hlink,attr"`
}

type Repository struct {
	Primary
}

type RepositoryRef struct {
	Handle string    `xml:"hlink,attr"`
	Medium string    `xml:"medium,attr"`
	Notes  []NoteRef `xml:"noteref"`
}

type Note struct {
	Primary
}

type NoteRef struct {
	Handle string `xml:"hlink,attr"`
}
