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
			json.UnmarshalFromFunc(l.unmarshalCitationHandle),
			json.UnmarshalFromFunc(l.unmarshalFamilyHandle),
			json.UnmarshalFromFunc(l.unmarshalMediaHandle),
			json.UnmarshalFromFunc(l.unmarshalNoteHandle),
			json.UnmarshalFromFunc(l.unmarshalPersonHandle),
			json.UnmarshalFromFunc(l.unmarshalPlaceHandle),
			json.UnmarshalFromFunc(l.unmarshalSourceHandle),
			json.UnmarshalFromFunc(l.unmarshalTagHandle),
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

func unmarshalHandle[V any](name string, d *jsontext.Decoder, values map[string]*V) (*V, error) {
	k := d.PeekKind()
	if k != jsontext.KindString {
		return nil, fmt.Errorf("can not unmarshal %s as %s handle", k, name)
	}
	token, err := d.ReadToken()
	if err != nil {
		return nil, err
	}
	handle := token.String()
	if handle == "" {
		return nil, nil
	}
	value, ok := values[handle]
	if !ok {
		return nil, fmt.Errorf("can not find %s referenced by handle %s", name, handle)
	}
	return value, nil
}

func (l *loader) unmarshalCitationHandle(d *jsontext.Decoder, r *gen.CitationHandle) error {
	citation, err := unmarshalHandle("citation", d, l.citations)
	if err != nil {
		return err
	}
	r.Value = citation
	return nil
}

func (l *loader) unmarshalFamilyHandle(d *jsontext.Decoder, r *gen.FamilyHandle) error {
	family, err := unmarshalHandle("family", d, l.families)
	if err != nil {
		return err
	}
	r.Value = family
	return nil
}

func (l *loader) unmarshalNoteHandle(d *jsontext.Decoder, r *gen.NoteHandle) error {
	note, err := unmarshalHandle("note", d, l.notes)
	if err != nil {
		return err
	}
	r.Value = note
	return nil
}

func (l *loader) unmarshalMediaHandle(d *jsontext.Decoder, r *gen.MediaHandle) error {
	media, err := unmarshalHandle("media", d, l.media)
	if err != nil {
		return err
	}
	r.Value = media
	return nil
}

func (l *loader) unmarshalPersonHandle(d *jsontext.Decoder, r *gen.PersonHandle) error {
	person, err := unmarshalHandle("person", d, l.people)
	if err != nil {
		return err
	}
	r.Value = person
	return nil
}

func (l *loader) unmarshalPlaceHandle(d *jsontext.Decoder, r *gen.PlaceHandle) error {
	place, err := unmarshalHandle("place", d, l.places)
	if err != nil {
		return err
	}
	r.Value = place
	return nil
}

func (l *loader) unmarshalSourceHandle(d *jsontext.Decoder, r *gen.SourceHandle) error {
	source, err := unmarshalHandle("source", d, l.sources)
	if err != nil {
		return err
	}
	r.Value = source
	return nil
}

func (l *loader) unmarshalTagHandle(d *jsontext.Decoder, r *gen.TagHandle) error {
	tag, err := unmarshalHandle("tag", d, l.tags)
	if err != nil {
		return err
	}
	r.Value = tag
	return nil
}
