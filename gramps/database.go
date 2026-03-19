// Package gramps unmarshals Gramps data structures from XML.
package gramps

import (
	"encoding/xml"
)

type DataMap struct {
	Header Header
	Tags   []Tag

	Citations    map[string]*Citation
	Events       map[string]*Event
	Families     map[string]*Family
	Media        map[string]*Media
	Notes        map[string]*Note
	People       map[string]*Person
	Places       map[string]*Place
	Repositories map[string]*Repository
	Sources      map[string]*Source
}

func NewDataMap(db *Database) *DataMap {
	m := &DataMap{
		Citations:    map[string]*Citation{},
		Events:       map[string]*Event{},
		Families:     map[string]*Family{},
		Media:        map[string]*Media{},
		Notes:        map[string]*Note{},
		People:       map[string]*Person{},
		Places:       map[string]*Place{},
		Repositories: map[string]*Repository{},
		Sources:      map[string]*Source{},
	}

	m.Header = db.Header
	m.Tags = db.Tags

	for _, item := range db.Citations {
		m.Citations[item.Handle] = &item
	}

	for _, item := range db.Events {
		m.Events[item.Handle] = &item
	}

	for _, item := range db.Families {
		m.Families[item.Handle] = &item
	}

	for _, item := range db.Media {
		m.Media[item.Handle] = &item
	}

	for _, item := range db.Notes {
		m.Notes[item.Handle] = &item
	}

	for _, item := range db.People {
		m.People[item.Handle] = &item
	}

	for _, item := range db.Places {
		m.Places[item.Handle] = &item
	}

	for _, item := range db.Repositories {
		m.Repositories[item.Handle] = &item
	}

	for _, item := range db.Sources {
		m.Sources[item.Handle] = &item
	}

	return m
}

type Database struct {
	XMLName xml.Name `xml:"database"`
	Header  Header
	Tags    []Tag `xml:"tags>tag"`

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
	Type        string        `xml:"type"`
	Date        DateVal       `xml:"dateval"`
	Place       PlaceRef      `xml:"place"`
	Description string        `xml:"description"`
	Attributes  []Attribute   `xml:"attribute"`
	Citations   []CitationRef `xml:"citationref"`
	Notes       []NoteRef     `xml:"noteref"`
	Media       []MediaRef    `xml:"mediaref"`
}

type EventRef struct {
	Handle     string        `xml:"hlink,attr"`
	Role       string        `xml:"role,attr"`
	Attributes []Attribute   `xml:"attribute"`
	Citations  []CitationRef `xml:"citationref"`
	Notes      []NoteRef     `xml:"noteref"`
}

type DateVal struct {
	Val  string `xml:"val,attr"`
	Type string `xml:"type,attr"`
}

type Person struct {
	Primary
	Gender     string        `xml:"gender"`
	Name       PersonName    `xml:"name"`
	ChildOf    []FamilyRef   `xml:"childof"`
	ParentIn   []FamilyRef   `xml:"parentin"`
	Attributes []Attribute   `xml:"attribute"`
	Citations  []CitationRef `xml:"citationref"`
	Events     []EventRef    `xml:"eventref"`
	Media      []MediaRef    `xml:"mediaref"`
	Notes      []NoteRef     `xml:"noteref"`
	URLs       []URL         `xml:"url"`
}

type PersonRef struct {
	Handle string `xml:"hlink,attr"`
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

type Attribute struct {
	Type      string        `xml:"type,attr"`
	Value     string        `xml:"value,attr"`
	Citations []CitationRef `xml:"citationref"`
	Notes     []NoteRef     `xml:"noteref"`
}

type Family struct {
	Primary
	Rel        Rel           `xml:"rel"`
	Father     PersonRef     `xml:"father"`
	Mother     PersonRef     `xml:"mother"`
	Children   []PersonRef   `xml:"childref"`
	Attributes []Attribute   `xml:"attribute"`
	Citations  []CitationRef `xml:"citationref"`
	Events     []EventRef    `xml:"eventref"`
	Media      []MediaRef    `xml:"mediaref"`
	Notes      []NoteRef     `xml:"noteref"`
}

type FamilyRef struct {
	Handle string `xml:"hlink,attr"`
}

type Rel struct {
	Type string `xml:"type,attr"`
}

type Citation struct {
	Primary
	Page       string      `xml:"page"`
	Confidence uint8       `xml:"confidence"`
	Date       DateVal     `xml:"datestr"`
	Attributes []Attribute `xml:"attribute"`
	Media      []MediaRef  `xml:"mediaref"`
	Notes      []NoteRef   `xml:"noteref"`
	Sources    []SourceRef `xml:"sourceref"`
}

type CitationRef struct {
	Handle string `xml:"hlink,attr"`
}

type Source struct {
	Primary
	Title        string          `xml:"stitle"`
	Author       string          `xml:"sauthor"`
	PubInfo      string          `xml:"spubinfo"`
	Attributes   []Attribute     `xml:"attribute"`
	Media        []MediaRef      `xml:"mediaref"`
	Notes        []NoteRef       `xml:"noteref"`
	Repositories []RepositoryRef `xml:"reporef"`
}

type SourceRef struct {
	Handle string `xml:"hlink,attr"`
}

type Place struct {
	Primary
	Type          string        `xml:"type,attr"`
	Name          PlaceName     `xml:"pname"`
	Coordinates   Coordinates   `xml:"coord"`
	EncompassedBy []PlaceRef    `xml:"placeref"`
	Citations     []CitationRef `xml:"citationref"`
	Media         []MediaRef    `xml:"mediaref"`
	Notes         []NoteRef     `xml:"noteref"`
	URLs          []URL         `xml:"url"`
}

type PlaceRef struct {
	Handle string  `xml:"hlink,attr"`
	Date   DateVal `xml:"dateval"`
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
	File       MediaFile     `xml:"file"`
	Date       DateVal       `xml:"dateval"`
	Attributes []Attribute   `xml:"attribute"`
	Citations  []CitationRef `xml:"citationref"`
	Notes      []NoteRef     `xml:"noteref"`
}

type MediaRef struct {
	Handle string `xml:"hlink,attr"`
}

type MediaFile struct {
	Source      string `xml:"src,attr"`
	Mime        string `xml:"mime,attr"`
	Description string `xml:"description,attr"`
	Checksum    string `xml:"checksum,attr"`
}
type Repository struct {
	Primary
	Name      string    `xml:"rname"`
	Type      string    `xml:"type"`
	Addresses []Address `xml:"address"`
	Notes     []NoteRef `xml:"noteref"`
	URLs      []URL     `xml:"url"`
}

type Address struct {
}

type RepositoryRef struct {
	Handle string    `xml:"hlink,attr"`
	Medium string    `xml:"medium,attr"`
	Notes  []NoteRef `xml:"noteref"`
}

type URL struct {
	HREF string `xml:"href,attr"`
	Type string `xml:"type,attr"`
}
type Note struct {
	Primary
	Type string `xml:"type,attr"`
	Text string `xml:"text"`
}

type NoteRef struct {
	Handle string `xml:"hlink,attr"`
}
