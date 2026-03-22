package load

import (
	"encoding/json/jsontext"
	"encoding/json/v2"
	"fmt"
	"time"

	"github.com/dhemery/gogramps/gen"
)

func (l *loader) unmarshalRecords() error {
	opts := grampsUnmarshalOptions(l)
	for handle, data := range l.citationRecords {
		citation, ok := l.citations[handle]
		if !ok {
			return fmt.Errorf("unmarshal has record but no citation: %s", handle)
		}
		if err := json.Unmarshal(data, citation, opts); err != nil {
			return err
		}
	}

	for handle, data := range l.eventRecords {
		event, ok := l.events[handle]
		if !ok {
			return fmt.Errorf("unmarshal has record but no event: %s", handle)
		}
		if err := json.Unmarshal(data, event, opts); err != nil {
			return err
		}
	}

	for handle, data := range l.familyRecords {
		family, ok := l.families[handle]
		if !ok {
			return fmt.Errorf("unmarshal has record but no family: %s", handle)
		}
		if err := json.Unmarshal(data, family, opts); err != nil {
			return err
		}
	}

	for handle, data := range l.mediaRecords {
		media, ok := l.media[handle]
		if !ok {
			return fmt.Errorf("unmarshal has record but no media: %s", handle)
		}
		if err := json.Unmarshal(data, media, opts); err != nil {
			return err
		}
	}

	for handle, data := range l.noteRecords {
		note, ok := l.notes[handle]
		if !ok {
			return fmt.Errorf("unmarshal has record but no note: %s", handle)
		}
		if err := json.Unmarshal(data, note, opts); err != nil {
			return err
		}
	}

	for handle, data := range l.personRecords {
		person, ok := l.people[handle]
		if !ok {
			return fmt.Errorf("unmarshal has record but no person: %s", handle)
		}
		if err := json.Unmarshal(data, person, opts); err != nil {
			return err
		}
	}

	for handle, data := range l.placeRecords {
		place, ok := l.places[handle]
		if !ok {
			return fmt.Errorf("unmarshal has record but no place: %s", handle)
		}
		if err := json.Unmarshal(data, place, opts); err != nil {
			return err
		}
	}

	for handle, data := range l.sourceRecords {
		source, ok := l.sources[handle]
		if !ok {
			return fmt.Errorf("unmarshal has record but no source: %s", handle)
		}
		if err := json.Unmarshal(data, source, opts); err != nil {
			return err
		}
	}

	for handle, data := range l.repositoryRecords {
		repository, ok := l.repositories[handle]
		if !ok {
			return fmt.Errorf("unmarshal has record but no repository: %s", handle)
		}
		if err := json.Unmarshal(data, repository, opts); err != nil {
			return err
		}
	}

	for handle, data := range l.tagRecords {
		tag, ok := l.tags[handle]
		if !ok {
			return fmt.Errorf("unmarshal has record but no tag: %s", handle)
		}
		if err := json.Unmarshal(data, tag, opts); err != nil {
			return err
		}
	}
	return nil
}

func grampsUnmarshalOptions(l *loader) jsontext.Options {
	unmarshalers := json.WithUnmarshalers(
		json.JoinUnmarshalers(
			json.UnmarshalFromFunc(unmarshalTime),
			json.UnmarshalFromFunc(l.unmarshalPersonRef),
			json.UnmarshalFromFunc(l.unmarshalPlaceRef),
			json.UnmarshalFromFunc(l.unmarshalSourceRef),
		))
	return json.JoinOptions(
		json.DefaultOptionsV2(),
		// json.OmitZeroStructFields(true),
		unmarshalers,
	)
}

// unmarshalTime unmarshals a time.Time from Gramps JSON, which encodes each
// record's modification time as the number of seconds (integer) from the UNIX
// epoch.
func unmarshalTime(d *jsontext.Decoder, t *time.Time) error {
	token, err := d.ReadToken()
	if err != nil {
		return err
	}
	*t = time.Unix(token.Int(), 0)
	return nil
}

func (l *loader) unmarshalPersonRef(d *jsontext.Decoder, s *gen.PersonRef) error {
	k := d.PeekKind()
	if k != jsontext.KindString {
		return fmt.Errorf("could not unmarshal as person handle: %s", k)
	}
	token, err := d.ReadToken()
	if err != nil {
		return err
	}
	handle := token.String()
	if handle == "" {
		return nil
	}
	person, ok := l.people[handle]
	if !ok {
		return fmt.Errorf("could not find referenced person: %s", handle)
	}
	s.Person = person
	return nil
}

func (l *loader) unmarshalPlaceRef(d *jsontext.Decoder, s *gen.PlaceRef) error {
	k := d.PeekKind()
	if k != jsontext.KindString {
		return fmt.Errorf("could not unmarshal as place handle: %s", k)
	}
	token, err := d.ReadToken()
	if err != nil {
		return err
	}
	handle := token.String()
	if handle == "" {
		return nil
	}
	place, ok := l.places[handle]
	if !ok {
		return fmt.Errorf("could not find referenced place: %s", handle)
	}
	s.Place = place
	return nil
}

func (l *loader) unmarshalSourceRef(d *jsontext.Decoder, s *gen.SourceRef) error {
	k := d.PeekKind()
	if k != jsontext.KindString {
		return fmt.Errorf("could not unmarshal as source handle: %s", k)
	}
	token, err := d.ReadToken()
	if err != nil {
		return err
	}
	handle := token.String()
	source, ok := l.sources[handle]
	if !ok {
		return fmt.Errorf("could not find referenced source: %s", handle)
	}
	s.Source = source
	return nil
}
