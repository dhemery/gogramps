package gramps

import (
	"encoding/xml"
	"log/slog"
)

type Address struct {
}

type Attribute struct {
	Type      string        `xml:"type,attr"`
	Value     string        `xml:"value,attr"`
	Citations []CitationRef `xml:"citationref"`
	Notes     []NoteRef     `xml:"noteref"`
}

type CitationRef struct {
	CitationHandle string `xml:"hlink,attr"`
}

type Coordinates struct {
	Longitude string `xml:"long,attr"`
	Latitude  string `xml:"lat,attr"`
}

// A DateVal represents a date or range of dates,
// possibly with quality such as "estimated" or "calculated"
// and modifiers such as "about" or "before."
type DateVal struct {
	Quality  string
	Modifier string
	Date     Date
}

func (dv *DateVal) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for _, a := range start.Attr {
		name := a.Name.Local
		if name == "xmlns" {
			continue
		}
		line, column := d.InputPos()
		switch name {
		case "quality":
			dv.Quality = a.Value
		case "type":
			dv.Modifier = a.Value
		case "val":
			if err := dv.Date.UnmarshalXMLAttr(a); err != nil {
				slog.Warn("could not parse DateVal date",
					"line", line, "column", column,
					"attr", a.Name.Local, "value", a.Value)
			}
		default:
			slog.Warn("unknown DateVal attr",
				"line", line, "column", column,
				"attr", a.Name.Local, "value", a.Value)
		}
	}
	// No inner elements, so skip through the end element.
	return d.Skip()
}

// An EventRef is a references from a person to an event in which the person played a role.
type EventRef struct {
	EventHandle string        `xml:"hlink,attr"`
	Role        string        `xml:"role,attr"`
	Attributes  []Attribute   `xml:"attribute"`
	Citations   []CitationRef `xml:"citationref"`
	Notes       []NoteRef     `xml:"noteref"`
}

type FamilyRef struct {
	FamilyHandle string `xml:"hlink,attr"`
}

type FamilyType struct {
	Type string `xml:"type,attr"`
}

// LDSOrdinance matches the `lds_ord` element. I don't use it, so just match and ignore.
type LDSOrdinance struct {
}

type MediaFile struct {
	Source      string `xml:"src,attr"`
	Mime        string `xml:"mime,attr"`
	Description string `xml:"description,attr"`
	Checksum    string `xml:"checksum,attr"`
}

type MediaRef struct {
	MediaHandle string `xml:"hlink,attr"`
}
type NoteRef struct {
	NoteHandle string `xml:"hlink,attr"`
}

type PersonName struct {
	Privacy
	Type       string        `xml:"type,attr"`
	Alt        string        `xml:"alt,attr"`
	Title      string        `xml:"title"`
	First      string        `xml:"first"`
	Surname    string        `xml:"surname"`
	Suffix     string        `xml:"suffix"`
	Date       []DateVal     `xml:"dateval"`
	Call       string        `xml:"call"`
	Nick       string        `xml:"nick"`
	FamilyNick string        `xml:"familynick"`
	Citations  []CitationRef `xml:"citationref"`
	Notes      []NoteRef     `xml:"noteref"`
}

type PersonRef struct {
	PersonHandle string `xml:"hlink,attr"`
}

type PlaceName struct {
	Value string `xml:"value,attr"`
}

// PlaceRef reprents that a place is or was encompassed by an encompassing place.
type PlaceRef struct {
	// Reference to the encompassing place.
	PlaceHandle string `xml:"hlink,attr"`
	// The date or dates of the relationship.
	Date []DateVal `xml:"dateval"`
}

// Privacy indicates whether an object is private
// and therefore should not be published.
type Privacy struct {
	Private bool `xml:"priv,attr"`
}

type RepositoryRef struct {
	RepositoryHandle string    `xml:"hlink,attr"`
	Medium           string    `xml:"medium,attr"`
	Notes            []NoteRef `xml:"noteref"`
}

type SourceRef struct {
	SourceHandle string `xml:"hlink,attr"`
}

type TagRef struct {
	TagHandle string `xml:"hlink,attr"`
}

// Tags is a mix-in that collects the containing element's
// references to tags.
type Tags struct {
	Tags []TagRef `xml:"tagref"`
}

type URL struct {
	HREF string `xml:"href,attr"`
	Type string `xml:"type,attr"`
}
