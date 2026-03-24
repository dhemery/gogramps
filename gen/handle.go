package gen

import "encoding/json/jsontext"

type CitationHandle struct {
	Value *Citation
}

func (h *CitationHandle) MarshalJSONTo(e *jsontext.Encoder) error {
	var handle string
	if h.Value != nil {
		handle = h.Value.Handle
	}
	return e.WriteToken(jsontext.String(handle))
}

type EventHandle struct {
	Value *Event
}

func (h *EventHandle) MarshalJSONTo(e *jsontext.Encoder) error {
	var handle string
	if h.Value != nil {
		handle = h.Value.Handle
	}
	return e.WriteToken(jsontext.String(handle))
}

type FamilyHandle struct {
	Value *Family
}

func (h *FamilyHandle) MarshalJSONTo(e *jsontext.Encoder) error {
	var handle string
	if h.Value != nil {
		handle = h.Value.Handle
	}
	return e.WriteToken(jsontext.String(handle))
}

type MediaHandle struct {
	Value *Media
}

func (h *MediaHandle) MarshalJSONTo(e *jsontext.Encoder) error {
	var handle string
	if h.Value != nil {
		handle = h.Value.Handle
	}
	return e.WriteToken(jsontext.String(handle))
}

type NoteHandle struct {
	Value *Note
}

func (h *NoteHandle) MarshalJSONTo(e *jsontext.Encoder) error {
	var handle string
	if h.Value != nil {
		handle = h.Value.Handle
	}
	return e.WriteToken(jsontext.String(handle))
}

type PersonHandle struct {
	Value *Person
}

func (h *PersonHandle) MarshalJSONTo(e *jsontext.Encoder) error {
	var handle string
	if h.Value != nil {
		handle = h.Value.Handle
	}
	return e.WriteToken(jsontext.String(handle))
}

type PlaceHandle struct {
	Value *Place
}

func (h *PlaceHandle) MarshalJSONTo(e *jsontext.Encoder) error {
	var handle string
	if h.Value != nil {
		handle = h.Value.Handle
	}
	return e.WriteToken(jsontext.String(handle))
}

type SourceHandle struct {
	Value *Source
}

func (h *SourceHandle) MarshalJSONTo(e *jsontext.Encoder) error {
	var handle string
	if h.Value != nil {
		handle = h.Value.Handle
	}
	return e.WriteToken(jsontext.String(handle))
}

type TagHandle struct {
	Value *Tag
}

func (r *TagHandle) MarshalJSONTo(e *jsontext.Encoder) error {
	var handle string
	if r.Value != nil {
		handle = r.Value.Handle
	}
	return e.WriteToken(jsontext.String(handle))
}
