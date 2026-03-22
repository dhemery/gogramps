package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	if len(os.Args) < 2 {
		exitError(errors.New("usage: gogramps dbfile"))
	}
	gramps, err := loadGramps(os.Args[1])
	if err != nil {
		exitError(err)
	}

	j, err := json.Marshal(gramps)
	if err != nil {
		exitError(err)
	}

	fmt.Println(string(j))
}

type personRecord struct {
	Data   string
	Person *Person
}

func exitError(err error) {
	slog.Error(err.Error())
	os.Exit(2)
}

type records struct {
	People map[string]personRecord
}

func newRecords() *records {
	return &records{
		People: map[string]personRecord{},
	}
}

type Gramps struct {
	People []*Person
}

type Person struct {
}

func loadGramps(fname string) (*Gramps, error) {
	// Phase 1: Load the records.
	records := newRecords()
	if err := records.load(fname); err != nil {
		return nil, err
	}

	// Phase 2: Convert to Gramps
	gramps := new(Gramps)
	records.copyTo(gramps)

	return gramps, nil
}

func (r *records) load(fname string) error {
	db, err := sql.Open("sqlite3", fname)
	if err != nil {
		return err
	}
	defer db.Close()

	return r.loadPeople(db)
}

func (r *records) loadPeople(db *sql.DB) error {
	query := `SELECT handle, json_data FROM person`
	rows, err := db.Query(query)
	if err != nil {
		return err
	}

	for {
		if ok := rows.Next(); !ok {
			break
		}
		person := new(Person)
		record := personRecord{Person: person}

		var handle string
		if err := rows.Scan(&handle, &record.Data); err != nil {
			return err
		}
		r.People[handle] = record
	}
	return nil
}

func (r *records) copyTo(gramps *Gramps) {
	for _, pr := range r.People {
		gramps.People = append(gramps.People, pr.Person)
	}
}
