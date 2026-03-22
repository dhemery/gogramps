// Package load loads data from a Gramps database.
package load

import (
	"database/sql"

	"github.com/dhemery/gogramps/gen"
)

func Gen(fname string) (*gen.Gen, error) {
	l := newLoader()
	if err := l.load(fname); err != nil {
		return nil, err
	}

	// TO DO: Unmarshal each record and resolve references.

	return l.gen, nil
}

type loader struct {
	gen *gen.Gen

	citations map[string]*gen.Citation
	people    map[string]*gen.Person

	citationRecords   map[string]string
	eventRecords      map[string]string
	familyRecords     map[string]string
	mediaRecords      map[string]string
	noteRecords       map[string]string
	personRecords     map[string]string
	placeRecords      map[string]string
	repositoryRecords map[string]string
	sourceRecords     map[string]string
	tagRecords        map[string]string
}

func newLoader() *loader {
	return &loader{
		gen:             new(gen.Gen),
		citationRecords: map[string]string{},
		personRecords:   map[string]string{},

		citations: map[string]*gen.Citation{},
		people:    map[string]*gen.Person{},
	}
}

func (l *loader) load(fname string) error {
	db, err := sql.Open("sqlite3", fname)
	if err != nil {
		return err
	}
	defer db.Close()

	// Citation
	// Event
	// Family
	// Media
	// Note

	citationRecords, err := l.readRecords(db, "person")
	l.citationRecords = citationRecords
	if err != nil {
		return err
	}

	for handle := range l.citationRecords {
		c := new(gen.Citation)
		l.citations[handle] = c
		l.gen.Citations = append(l.gen.Citations, c)
	}

	personRecords, err := l.readRecords(db, "person")
	l.personRecords = personRecords
	if err != nil {
		return err
	}

	for handle := range l.personRecords {
		p := new(gen.Person)
		l.people[handle] = p
		l.gen.People = append(l.gen.People, p)
	}

	// Place
	// Reference? I think this is a redundant table of all references.
	// Repository
	// Source
	// Tag

	return nil
}

func (l *loader) readRecords(db *sql.DB, tableName string) (map[string]string, error) {
	// Every record type has a handle, json_data, and other fields specific
	// to the subject. But the JSON has all of that data and more, so we
	// can ignore the fields and reconstruct the data from the JSON. This
	// allows us to read every record type using a query that differs only
	// in the name of the table to read.
	query := `SELECT handle, json_data FROM ` + tableName
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	records := make(map[string]string)

	for {
		if ok := rows.Next(); !ok {
			break
		}
		var handle string
		var data string
		if err := rows.Scan(&handle, &data); err != nil {
			return records, err
		}
		records[handle] = data
	}
	return records, nil
}
