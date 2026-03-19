package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"dhemery.com/gogramps/convert"
	"dhemery.com/gogramps/gramps"
)

func main() {
	if len(os.Args) < 2 {
		exitError(errors.New("missing file argument"))
	}

	grampsDB, err := gramps.Read(os.Args[1])
	if err != nil {
		exitError(err)
	}

	genDB, err := convert.Convert(grampsDB)
	if err != nil {
		exitError(err)
	}

	asJSON, err := json.Marshal(genDB)
	if err != nil {
		exitError(err)
	}

	fmt.Println(string(asJSON))
}

func exitError(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(2)
}
