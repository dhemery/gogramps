package gramps

import (
	"encoding/xml"
	"errors"
	"io"
	"os"

	"dhemery.com/gogramps/gen"
)

func Read(fname string) (*gen.DB, error) {
	grampsFile, err := os.Open(fname)
	if err != nil {
		return nil, err
	}
	defer grampsFile.Close()

	xmlBytes, err := io.ReadAll(grampsFile)
	if err != nil {
		return nil, err
	}

	grampsDB := &DB{}
	err = xml.Unmarshal(xmlBytes, grampsDB)

	genDB, err := convert(grampsDB)
	if err != nil {
		return nil, err
	}

	return genDB, err

}

var ErrHasUnknown = errors.New("has unknown fields or attrs")

// Unknown collects and reports unknown XML fields and attributes.
type Unknown struct {
	UnknownFields []UnknownField `xml:",any"`
	UnknownAttrs  []string       `xml:",any,attr"`
}

type UnknownField struct {
	XMLName xml.Name
	Value   string `xml:",innerxml"`
}

func (u Unknown) CheckUnknown() error {
	if len(u.UnknownFields) > 0 || len(u.UnknownAttrs) > 0 {
		return ErrHasUnknown
	}
	return nil
}
