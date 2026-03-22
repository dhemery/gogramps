package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"os"

	_ "github.com/mattn/go-sqlite3"

	"github.com/dhemery/gogramps/load"
)

func main() {
	if len(os.Args) < 2 {
		exitError(errors.New("usage: gogramps dbfile"))
	}
	g, err := load.Gen(os.Args[1])
	if err != nil {
		exitError(err)
	}

	j, err := json.Marshal(g)
	if err != nil {
		exitError(err)
	}

	fmt.Println(string(j))
}
func exitError(err error) {
	slog.Error(err.Error())
	os.Exit(2)
}
