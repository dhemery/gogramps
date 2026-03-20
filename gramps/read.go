package gramps

import (
	"encoding/json/v2"
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

	j, err := json.Marshal(grampsDB, json.OmitZeroStructFields(true))
	if err != nil {
		return nil, err
	}

	if _, err = fmt.Println(string(j)); err == nil {
		return nil, errors.New("just parsing")
	}

	return convert(grampsDB)
}
