package convert

import (
	"dhemery.com/gogramps/gramps"
)

type GrampsMap struct {
	Header gramps.Header
	Tags   []gramps.Tag

	Citations    map[string]*gramps.Citation
	Events       map[string]*gramps.Event
	Families     map[string]*gramps.Family
	Media        map[string]*gramps.Media
	Notes        map[string]*gramps.Note
	People       map[string]*gramps.Person
	Places       map[string]*gramps.Place
	Repositories map[string]*gramps.Repository
	Sources      map[string]*gramps.Source
}

func NewGrampsMap(grampsDB *gramps.DB) *GrampsMap {
	m := &GrampsMap{
		Header:       grampsDB.Header,
		Tags:         grampsDB.Tags,
		Citations:    map[string]*gramps.Citation{},
		Events:       map[string]*gramps.Event{},
		Families:     map[string]*gramps.Family{},
		Media:        map[string]*gramps.Media{},
		Notes:        map[string]*gramps.Note{},
		People:       map[string]*gramps.Person{},
		Places:       map[string]*gramps.Place{},
		Repositories: map[string]*gramps.Repository{},
		Sources:      map[string]*gramps.Source{},
	}

	for _, item := range grampsDB.Citations {
		m.Citations[item.Handle] = &item
	}

	for _, item := range grampsDB.Events {
		m.Events[item.Handle] = &item
	}

	for _, item := range grampsDB.Families {
		m.Families[item.Handle] = &item
	}

	for _, item := range grampsDB.Media {
		m.Media[item.Handle] = &item
	}

	for _, item := range grampsDB.Notes {
		m.Notes[item.Handle] = &item
	}

	for _, item := range grampsDB.People {
		m.People[item.Handle] = &item
	}

	for _, item := range grampsDB.Places {
		m.Places[item.Handle] = &item
	}

	for _, item := range grampsDB.Repositories {
		m.Repositories[item.Handle] = &item
	}

	for _, item := range grampsDB.Sources {
		m.Sources[item.Handle] = &item
	}

	return m
}
