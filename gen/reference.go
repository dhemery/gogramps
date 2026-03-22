package gen

import "encoding/json/jsontext"

type PlaceRef struct {
	Place *Place
}

func (r *PlaceRef) MarshalJSONTo(e*jsontext.Encoder) error {
	if r.Place == nil {
		return e.WriteToken(jsontext.String(""))
	}
	return e.WriteToken(jsontext.String(r.Place.Handle))
}

type SourceRef struct {
	Source *Source
}

func (r *SourceRef) MarshalJSONTo(e*jsontext.Encoder) error {
	return e.WriteToken(jsontext.String(r.Source.Handle))
}
