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

	if err := unmarshalRecords("citation", l.citationRecords, l.citations, opts); err != nil {
		return err
	}

	if err := unmarshalRecords("event", l.eventRecords, l.events, opts); err != nil {
		return err
	}

	if err := unmarshalRecords("family", l.familyRecords, l.families, opts); err != nil {
		return err
	}

	if err := unmarshalRecords("media", l.mediaRecords, l.media, opts); err != nil {
		return err
	}

	if err := unmarshalRecords("note", l.noteRecords, l.notes, opts); err != nil {
		return err
	}

	if err := unmarshalRecords("person", l.personRecords, l.people, opts); err != nil {
		return err
	}

	if err := unmarshalRecords("place", l.placeRecords, l.places, opts); err != nil {
		return err
	}

	if err := unmarshalRecords("source", l.sourceRecords, l.sources, opts); err != nil {
		return err
	}

	if err := unmarshalRecords("repository", l.repositoryRecords, l.repositories, opts); err != nil {
		return err
	}

	if err := unmarshalRecords("tag", l.tagRecords, l.tags, opts); err != nil {
		return err
	}

	return nil
}

func unmarshalRecords[T any](name string, records map[string][]byte, values map[string]*T, opts json.Options) error {
	for handle, data := range records {
		value, ok := values[handle]
		if !ok {
			return fmt.Errorf("unmarshal: loader has %s record but no %s: %s",
				name, name, handle)
		}
		if err := json.Unmarshal(data, value, opts); err != nil {
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
