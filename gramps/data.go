// Package gramps unmarshals Gramps data structures from XML.
package gramps

import (
	"encoding/xml"
	"errors"
)

var ErrHasUnknown = errors.New("has unknown fields or attrs")

type UnknownField struct {
	XMLName xml.Name
	Value   string `xml:",innerxml"`
}

// Unknown collects and reports unknown XML fields and attributes.
type Unknown struct {
	UnknownFields []UnknownField `xml:",any"`
	UnknownAttrs  []string       `xml:",any,attr"`
}

func (u Unknown) CheckUnknown() error {
	if len(u.UnknownFields) > 0 || len(u.UnknownAttrs) > 0 {
		return ErrHasUnknown
	}
	return nil
}

// Privacy indicates whether an object is private
// and therefore should not be published.
type Privacy struct {
	Private bool `xml:"priv,attr"`
}

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

// Tags is a mix-in that collects the containing object's tags.
type Tags struct {
	Tags []TagRef `xml:"tagref"`
}

type Tag struct {
	TableObject
	Name     string `xml:"name,attr"`
	Color    string `xml:"color,attr"`
	Priority uint   `xml:"priority,attr"`
}

type TagRef struct {
	TagHandle string `xml:"hlink,attr"`
}

type Event struct {
	PrimaryObject
	Type        string        `xml:"type"`
	Date        DateVal       `xml:"dateval"`
	Place       PlaceRef      `xml:"place"`
	Description string        `xml:"description"`
	Attributes  []Attribute   `xml:"attribute"`
	Citations   []CitationRef `xml:"citationref"`
	Notes       []NoteRef     `xml:"noteref"`
	Media       []MediaRef    `xml:"objref"`
}

type EventRef struct {
	EventHandle string        `xml:"hlink,attr"`
	Role        string        `xml:"role,attr"`
	Attributes  []Attribute   `xml:"attribute"`
	Citations   []CitationRef `xml:"citationref"`
	Notes       []NoteRef     `xml:"noteref"`
}

// A DateVal represents a date or range of dates.
type DateVal struct {
	Val  string `xml:"val,attr"`
	Type string `xml:"type,attr"`
}

type Person struct {
	PrimaryObject
	Unknown
	Privacy
	Gender        string         `xml:"gender"`
	Name          PersonName     `xml:"name"`
	Addresses     []Address      `xml:"address"`
	Associations  []PersonRef    `xml:"personref"`
	ChildOf       []FamilyRef    `xml:"childof"`
	ParentIn      []FamilyRef    `xml:"parentin"`
	Attributes    []Attribute    `xml:"attribute"`
	Citations     []CitationRef  `xml:"citationref"`
	Events        []EventRef     `xml:"eventref"`
	LdsOrdinances []LDSOrdinance `xml:"lds_ord"`
	Media         []MediaRef     `xml:"objref"`
	Notes         []NoteRef      `xml:"noteref"`
	URLs          []URL          `xml:"url"`
}

type PersonRef struct {
	PersonHandle string `xml:"hlink,attr"`
}

type PersonName struct {
	First     string        `xml:"first"`
	Call      string        `xml:"call"`
	Nick      string        `xml:"nick"`
	Surname   string        `xml:"surname"`
	Suffix    string        `xml:"suffix"`
	DateStr   DateVal       `xml:"datestr"`
	Citations []CitationRef `xml:"citationref"`
	Notes     []NoteRef     `xml:"noteref"`
}

// LDSOrdinance matches the `lds_ord` element. I don't use it, so just match and ignore.
type LDSOrdinance struct {
}

type Attribute struct {
	Type      string        `xml:"type,attr"`
	Value     string        `xml:"value,attr"`
	Citations []CitationRef `xml:"citationref"`
	Notes     []NoteRef     `xml:"noteref"`
}

type Family struct {
	PrimaryObject
	Rel        Rel           `xml:"rel"`
	Father     PersonRef     `xml:"father"`
	Mother     PersonRef     `xml:"mother"`
	Children   []PersonRef   `xml:"childref"`
	Attributes []Attribute   `xml:"attribute"`
	Citations  []CitationRef `xml:"citationref"`
	Events     []EventRef    `xml:"eventref"`
	Media      []MediaRef    `xml:"objref"`
	Notes      []NoteRef     `xml:"noteref"`
}

type FamilyRef struct {
	FamilyHandle string `xml:"hlink,attr"`
}

type Rel struct {
	Type string `xml:"type,attr"`
}

type Citation struct {
	PrimaryObject
	Page       string      `xml:"page"`
	Confidence uint8       `xml:"confidence"`
	Date       DateVal     `xml:"datestr"`
	Attributes []Attribute `xml:"attribute"`
	Media      []MediaRef  `xml:"objref"`
	Notes      []NoteRef   `xml:"noteref"`
	Sources    []SourceRef `xml:"sourceref"`
}

type CitationRef struct {
	CitationHandle string `xml:"hlink,attr"`
}

type Source struct {
	PrimaryObject
	Title        string          `xml:"stitle"`
	Author       string          `xml:"sauthor"`
	PubInfo      string          `xml:"spubinfo"`
	Attributes   []Attribute     `xml:"attribute"`
	Media        []MediaRef      `xml:"objref"`
	Notes        []NoteRef       `xml:"noteref"`
	Repositories []RepositoryRef `xml:"reporef"`
}

type SourceRef struct {
	SourceHandle string `xml:"hlink,attr"`
}

type Place struct {
	PrimaryObject
	Type          string        `xml:"type,attr"`
	Name          PlaceName     `xml:"pname"`
	Coordinates   Coordinates   `xml:"coord"`
	EncompassedBy []PlaceRef    `xml:"placeref"`
	Citations     []CitationRef `xml:"citationref"`
	Media         []MediaRef    `xml:"objref"`
	Notes         []NoteRef     `xml:"noteref"`
	URLs          []URL         `xml:"url"`
}

type PlaceRef struct {
	PlaceHandle string  `xml:"hlink,attr"`
	Date        DateVal `xml:"dateval"`
}

type Coordinates struct {
	Longitude string `xml:"long,attr"`
	Latitude  string `xml:"lat,attr"`
}

type PlaceName struct {
	Value string `xml:"value,attr"`
}

type Media struct {
	PrimaryObject
	File       MediaFile     `xml:"file"`
	Date       DateVal       `xml:"dateval"`
	Attributes []Attribute   `xml:"attribute"`
	Citations  []CitationRef `xml:"citationref"`
	Notes      []NoteRef     `xml:"noteref"`
}

type MediaRef struct {
	MediaHandle string `xml:"hlink,attr"`
}

type MediaFile struct {
	Source      string `xml:"src,attr"`
	Mime        string `xml:"mime,attr"`
	Description string `xml:"description,attr"`
	Checksum    string `xml:"checksum,attr"`
}
type Repository struct {
	PrimaryObject
	Name      string    `xml:"rname"`
	Type      string    `xml:"type"`
	Addresses []Address `xml:"address"`
	Notes     []NoteRef `xml:"noteref"`
	URLs      []URL     `xml:"url"`
}

type Address struct {
}

type RepositoryRef struct {
	RepositoryHandle string    `xml:"hlink,attr"`
	Medium           string    `xml:"medium,attr"`
	Notes            []NoteRef `xml:"noteref"`
}

type URL struct {
	HREF string `xml:"href,attr"`
	Type string `xml:"type,attr"`
}
type Note struct {
	PrimaryObject
	Type string `xml:"type,attr"`
	Text string `xml:"text"`
}

type NoteRef struct {
	NoteHandle string `xml:"hlink,attr"`
}
