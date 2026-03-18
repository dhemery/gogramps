package main

import (
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

	// dec := xml.NewDecoder(grampsFile)

	xmlBytes, err := io.ReadAll(grampsFile)
	if err != nil {
		exitError(err)
	}

	db := gramps.Database{}

	err = xml.Unmarshal(xmlBytes, &db)
	// err = dec.Decode(&db)
	if err != nil {
		exitError(err)
	}

	fmt.Printf("%#v\n", db)
}

func exitError(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(2)
}
