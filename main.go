package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"dhemery.com/gogramps/gramps"
)

func main() {
	if len(os.Args) < 2 {
		exitError(errors.New("missing file argument"))
	}

	db, err := gramps.Read(os.Args[1])
	if err != nil {
		exitError(err)
	}

	asJSON, err := json.Marshal(db)
	if err != nil {
		exitError(err)
	}

	fmt.Println(string(asJSON))
}

func exitError(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(2)
}
