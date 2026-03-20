// Package gramps unmarshals Gramps data structures from XML.
package gramps

import (
	"bytes"
	"encoding/json/jsontext"
	"encoding/xml"
	"log/slog"
	"time"
)

type DB struct {
	Header       Header                 `json:"header"`
	HomePerson   string                 `json:"home_person"`
	Tags         []Tag                  `json:"tags"`
	Citations    map[string]*Citation   `json:"citations"`
	Events       map[string]*Event      `json:"events"`
	Families     map[string]*Family     `json:"families"`
	Media        map[string]*Media      `json:"media"`
	Notes        map[string]*Note       `json:"notes"`
	People       map[string]*Person     `json:"people"`
	Places       map[string]*Place      `json:"places"`
	Repositories map[string]*Repository `json:"repositories"`
	Sources      map[string]*Source     `json:"sources"`
}

func (db *DB) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	db.Citations = map[string]*Citation{}
	db.Events = map[string]*Event{}
	db.Families = map[string]*Family{}
	db.Media = map[string]*Media{}
	db.Notes = map[string]*Note{}
	db.People = map[string]*Person{}
	db.Places = map[string]*Place{}
	db.Repositories = map[string]*Repository{}
	db.Sources = map[string]*Source{}

	slog.Debug("unmarshaling element", "element", start.Name.Local)
	for {
		for _, a := range start.Attr {
			if a.Name.Local == "xmlns" {
				continue
			}
			slog.Warn("unknown <database> attr", "attr", a.Name.Local, "value", a.Value)
		}

		token, err := d.Token()
		if err != nil {
			return err
		}
		switch e := token.(type) {
		case xml.StartElement:
			if err := db.UnmarshalElement(e, d); err != nil {
				return err
			}
		case xml.CharData:
			if len(bytes.TrimSpace(e)) > 0 {
				slog.Debug("ignoring non-empty <database> CharData", "data", e)
			}
		case xml.EndElement:
			if e == start.End() {
				return nil
			}
			slog.Debug("end of element", "element", e.Name.Local)
		case xml.Comment:
			slog.Debug("ignoring comment", "comment", string(e))
		case xml.Directive:
			slog.Debug("ignoring directive", "directive", string(e))
		case xml.ProcInst:
			slog.Debug("ignoring processing instruction", "inst", e.Inst)
		}
	}
}

func (db *DB) UnmarshalElement(e xml.StartElement, d *xml.Decoder) error {
	slog.Debug("unmarshaling <database> element", "element", e.Name.Local)
	switch e.Name.Local {

	// These elements are collections of records. No need to decode them directly.
	// The outer decoder will loop over the records, which we handle below.
	case "citations":
		fallthrough
	case "events":
		fallthrough
	case "families":
		fallthrough
	case "objects":
		fallthrough
	case "notes":
		fallthrough
	case "places":
		fallthrough
	case "repositories":
		fallthrough
	case "sources":
		// Warn of any attributes we aren't currently handling.
		for _, a := range e.Attr {
			if a.Name.Local == "xmlns" {
				continue
			}
			slog.Warn("unknown attr",
				"element", e.Name.Local,
				"attr", a.Name.Local,
				"value", a.Value)
		}

	case "citation":
		item := new(Citation)
		if err := d.DecodeElement(item, &e); err != nil {
			return err
		}
		db.Citations[item.Handle] = item

	case "event":
		item := new(Event)
		if err := d.DecodeElement(item, &e); err != nil {
			return err
		}
		db.Events[item.Handle] = item

	case "family":
		item := new(Family)
		if err := d.DecodeElement(&item, &e); err != nil {
			return err
		}
		db.Families[item.Handle] = item

	case "header":
		return d.DecodeElement(&db.Header, &e)

	case "note":
		item := new(Note)
		if err := d.DecodeElement(&item, &e); err != nil {
			return err
		}
		db.Notes[item.Handle] = item

	case "object":
		item := new(Media)
		if err := d.DecodeElement(&item, &e); err != nil {
			return err
		}
		db.Media[item.Handle] = item

	case "people":
		for _, a := range e.Attr {
			if a.Name.Local == "home" {
				db.HomePerson = a.Value
			} else {
				slog.Warn("skipping unknown <person> attr",
					"name", a.Name.Local, "value", a.Value)
			}
		}

	case "person":
		item := new(Person)
		if err := d.DecodeElement(&item, &e); err != nil {
			return err
		}
		db.People[item.Handle] = item

	case "placeobj":
		item := new(Place)
		if err := d.DecodeElement(&item, &e); err != nil {
			return err
		}
		db.Places[item.Handle] = item

	case "repository":
		item := new(Repository)
		if err := d.DecodeElement(&item, &e); err != nil {
			return err
		}
		db.Repositories[item.Handle] = item

	case "source":
		item := new(Source)
		if err := d.DecodeElement(item, &e); err != nil {
			return err
		}
		db.Sources[item.Handle] = item

	case "tags":
		return d.DecodeElement(&db.Tags, &e)

	// Skip these elements.
	case "bookmarks":
		fallthrough
	case "name-formats":
		fallthrough
	case "namemaps":
		return d.Skip()

	// Warn of any elements we aren't currently handling.
	default:
		slog.Warn("skipping unknown <database> element", "element", e.Name.Local)
		return d.Skip()
	}
	return nil
}

type Header struct {
	Created    Created    `xml:"created" json:"created"`
	Researcher Researcher `xml:"researcher" json:"researcher"`
	MediaPath  string     `xml:"mediapath" json:"media_path"`
}

type Created struct {
	Date    Date   `xml:"date,attr"`
	Version string `xml:"version,attr"`
}

type Researcher struct {
}

// Date is a time represented in XML and JSON in the [time.DateOnly] layout.
type Date time.Time

func (d *Date) UnmarshalXMLAttr(a xml.Attr) error {
	t, err := time.Parse(time.DateOnly, a.Value)
	if err != nil {
		return err
	}
	*d = Date(t)
	return nil
}

func (d Date) MarshalJSONTo(e *jsontext.Encoder) error {
	return e.WriteToken(jsontext.String(time.Time(d).Format(time.DateOnly)))
}
