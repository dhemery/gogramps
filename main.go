package main

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"os"

	"dhemery.com/gogramps/gramps"
)

func main() {
	if len(os.Args) < 2 {
		exitError(errors.New("missing file argument"))
		os.Exit(2)
	}
	fname := os.Args[1]

	grampsFile, err := os.Open(fname)
	if err != nil {
		exitError(err)
	}
	defer grampsFile.Close()

	xmlBytes, err := io.ReadAll(grampsFile)
	if err != nil {
		exitError(err)
	}

	db := new(gramps.Database)

	err = xml.Unmarshal(xmlBytes, db)
	if err != nil {
		exitError(err)
	}

	dm := gramps.NewDataMap(db)
	asJSON, err := json.Marshal(dm)

	if err != nil {
		exitError(err)
	}

	fmt.Println(string(asJSON))
}

func exitError(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(2)
}
