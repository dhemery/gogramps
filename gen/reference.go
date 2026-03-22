package gen

import "encoding/json/jsontext"

type PersonRef struct {
	Person *Person
}

func (r *PersonRef) MarshalJSONTo(e *jsontext.Encoder) error {
	var handle string
	if r.Person != nil {
		handle = r.Person.Handle
	}
	return e.WriteToken(jsontext.String(handle))
}

type PlaceRef struct {
	Place *Place
}

func (r *PlaceRef) MarshalJSONTo(e *jsontext.Encoder) error {
	var handle string
	if r.Place != nil {
		handle = r.Place.Handle
	}
	return e.WriteToken(jsontext.String(handle))
}

type SourceRef struct {
	Source *Source
}

func (r *SourceRef) MarshalJSONTo(e *jsontext.Encoder) error {
	var handle string
	if r.Source != nil {
		handle = r.Source.Handle
	}
	return e.WriteToken(jsontext.String(handle))
}

type TagRef struct {
	Tag *Tag
}


func (r *TagRef) MarshalJSONTo(e *jsontext.Encoder) error {
	var handle string
	if r.Tag != nil {
		handle = r.Tag.Handle
	}
	return e.WriteToken(jsontext.String(handle))
}
