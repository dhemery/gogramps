// Package adapter converts Gramps data into Gen data.
package adapter

import (
	"time"

	"dhemery.com/gogramps/gen"
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

func Adapt(in *gramps.Database) *gen.DB {
	out := gen.NewDB()
	for _, p := range in.People {
		out.People[p.Handle] = newPerson(&p)
	}

	return out
}

func newPerson(in *gramps.Person) *gen.Person {
	inName := in.Name
	outName := gen.PersonName{
		First:   inName.First,
		Surname: inName.Surname,
		Suffix:  inName.Suffix,
		Call:    inName.Call,
		Nick:    inName.Nick,
	}

	return &gen.Person{
		Primary: adaptPrimary(in.Primary),
		Name:   outName,
		Gender: in.Gender,
	}
}

func adaptPrimary(in gramps.Primary) gen.Primary {
	return gen.Primary{
		Handle:  in.Handle,
		ID:      in.ID,
		Changed: time.Unix(int64(in.Change), 0),
	}
}

func NewGrampsMap(grampsDB *gramps.Database) *GrampsMap {
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
