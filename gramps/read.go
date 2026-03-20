package gramps

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
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
	if err != nil {
		return nil, err
	}

	unknowns := grampsDB.CollectUnknowns()
	if len(unknowns) > 0 {
		dump(unknowns)
		return nil, ErrHasUnknown
	}

	genDB, err := convert(grampsDB)
	if err != nil {
		return nil, err
	}

	return genDB, err

}
var ErrHasUnknown = errors.New("has unknown fields or attrs")

// Unknown collects and reports unknown XML fields and attributes.
type Unknown struct {
	UnknownFields []UnknownField `xml:",any" json:",omitempty"`
	UnknownAttrs  []string       `xml:",any,attr" json:",omitempty"`
}

type UnknownField struct {
	XMLName xml.Name
	Value   string `xml:",innerxml"`
}

func (u Unknown) hasUnknowns() bool {
	return len(u.UnknownFields) > 0 || len(u.UnknownAttrs) > 0
}

func dump(in any) {
	j, _ := json.Marshal(in)
	fmt.Println(string(j))
}
