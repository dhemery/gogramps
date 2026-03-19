package gramps

// A primaryMap is a map of the Gramps "primary" objects, indexed by their handles.
type primaryMap struct {
	Citations    map[string]*Citation
	Events       map[string]*Event
	Families     map[string]*Family
	Media        map[string]*Media
	Notes        map[string]*Note
	People       map[string]*Person
	Places       map[string]*Place
	Repositories map[string]*Repository
	Sources      map[string]*Source
}

func newPrimaryMap(db *DB) *primaryMap {
	m := &primaryMap{
		Citations:    map[string]*Citation{},
		Events:       map[string]*Event{},
		Families:     map[string]*Family{},
		Media:        map[string]*Media{},
		Notes:        map[string]*Note{},
		People:       map[string]*Person{},
		Places:       map[string]*Place{},
		Repositories: map[string]*Repository{},
		Sources:      map[string]*Source{},
	}

	for _, item := range db.Citations {
		m.Citations[item.Handle] = &item
	}

	for _, item := range db.Events {
		m.Events[item.Handle] = &item
	}

	for _, item := range db.Families {
		m.Families[item.Handle] = &item
	}

	for _, item := range db.Media {
		m.Media[item.Handle] = &item
	}

	for _, item := range db.Notes {
		m.Notes[item.Handle] = &item
	}

	for _, item := range db.People {
		m.People[item.Handle] = &item
	}

	for _, item := range db.Places {
		m.Places[item.Handle] = &item
	}

	for _, item := range db.Repositories {
		m.Repositories[item.Handle] = &item
	}

	for _, item := range db.Sources {
		m.Sources[item.Handle] = &item
	}

	return m
}
