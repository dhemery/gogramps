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
	for handle, data := range l.sourceRecords {
		source, ok := l.sources[handle]
		if !ok {
			return fmt.Errorf("unmarshal has record but no source: %s", handle)
		}
		if err := json.Unmarshal(data, source, opts); err != nil {
			return err
		}
	}
	return nil
}

func grampsUnmarshalOptions(l *loader) jsontext.Options {
	return json.WithUnmarshalers(
		json.JoinUnmarshalers(
			json.UnmarshalFromFunc(unmarshalTime),
			json.UnmarshalFromFunc(l.unmarshalSourceRef),
		))
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
