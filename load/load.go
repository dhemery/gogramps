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

	citations    map[string]*gen.Citation
	events       map[string]*gen.Event
	families     map[string]*gen.Family
	media        map[string]*gen.Media
	notes        map[string]*gen.Note
	people       map[string]*gen.Person
	places       map[string]*gen.Place
	repositories map[string]*gen.Repository
	sources      map[string]*gen.Source
	tags         map[string]*gen.Tag

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
		gen:               new(gen.Gen),
		citationRecords:   map[string]string{},
		eventRecords:      map[string]string{},
		familyRecords:     map[string]string{},
		mediaRecords:      map[string]string{},
		noteRecords:       map[string]string{},
		personRecords:     map[string]string{},
		placeRecords:      map[string]string{},
		repositoryRecords: map[string]string{},
		sourceRecords:     map[string]string{},
		tagRecords:        map[string]string{},

		citations:    map[string]*gen.Citation{},
		events:       map[string]*gen.Event{},
		families:     map[string]*gen.Family{},
		media:        map[string]*gen.Media{},
		notes:        map[string]*gen.Note{},
		people:       map[string]*gen.Person{},
		places:       map[string]*gen.Place{},
		repositories: map[string]*gen.Repository{},
		sources:      map[string]*gen.Source{},
		tags:         map[string]*gen.Tag{},
	}
}

func (l *loader) load(fname string) error {
	db, err := sql.Open("sqlite3", fname)
	if err != nil {
		return err
	}
	defer db.Close()

	citationRecords, err := l.readRecords(db, "person")
	l.citationRecords = citationRecords
	if err != nil {
		return err
	}

	for handle := range l.citationRecords {
		citation := new(gen.Citation)
		l.citations[handle] = citation
		l.gen.Citations = append(l.gen.Citations, citation)
	}

	eventRecords, err := l.readRecords(db, "event")
	l.eventRecords = eventRecords
	if err != nil {
		return err
	}

	for handle := range l.eventRecords {
		event := new(gen.Event)
		l.events[handle] = event
		l.gen.Events = append(l.gen.Events, event)
	}

	familyRecords, err := l.readRecords(db, "family")
	l.familyRecords = familyRecords
	if err != nil {
		return err
	}

	for handle := range l.familyRecords {
		family := new(gen.Family)
		l.families[handle] = family
		l.gen.Families = append(l.gen.Families, family)
	}

	mediaRecords, err := l.readRecords(db, "media")
	l.mediaRecords = mediaRecords
	if err != nil {
		return err
	}

	for handle := range l.mediaRecords {
		media := new(gen.Media)
		l.media[handle] = media
		l.gen.Media = append(l.gen.Media, media)
	}

	noteRecords, err := l.readRecords(db, "note")
	l.noteRecords = noteRecords
	if err != nil {
		return err
	}

	for handle := range l.noteRecords {
		note := new(gen.Note)
		l.notes[handle] = note
		l.gen.Notes = append(l.gen.Notes, note)
	}

	personRecords, err := l.readRecords(db, "person")
	l.personRecords = personRecords
	if err != nil {
		return err
	}

	for handle := range l.personRecords {
		person := new(gen.Person)
		l.people[handle] = person
		l.gen.People = append(l.gen.People, person)
	}

	placeRecords, err := l.readRecords(db, "place")
	l.placeRecords = placeRecords
	if err != nil {
		return err
	}

	for handle := range l.placeRecords {
		place := new(gen.Place)
		l.places[handle] = place
		l.gen.Places = append(l.gen.Places, place)
	}

	repositoryRecords, err := l.readRecords(db, "repository")
	l.repositoryRecords = repositoryRecords
	if err != nil {
		return err
	}

	for handle := range l.repositoryRecords {
		repository := new(gen.Repository)
		l.repositories[handle] = repository
		l.gen.Repositories = append(l.gen.Repositories, repository)
	}

	sourceRecords, err := l.readRecords(db, "source")
	l.sourceRecords = sourceRecords
	if err != nil {
		return err
	}

	for handle := range l.sourceRecords {
		source := new(gen.Source)
		l.sources[handle] = source
		l.gen.Sources = append(l.gen.Sources, source)
	}

	tagRecords, err := l.readRecords(db, "tag")
	l.tagRecords = tagRecords
	if err != nil {
		return err
	}

	for handle := range l.tagRecords {
		tag := new(gen.Tag)
		l.tags[handle] = tag
		l.gen.Tags = append(l.gen.Tags, tag)
	}

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
