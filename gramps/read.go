package gramps

import (
	"encoding/xml"
	"io"
	"os"
)

func Read(fname string) (*DB, error) {
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

	return grampsDB, err

}
