package gramps

import (
	"encoding/xml"
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
