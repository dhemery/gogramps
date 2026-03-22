// Package gen represents all of the genealogy data.
package gen

type Gen struct {
	Citations    []*Citation
	Events       []*Event
	Families     []*Family
	Media        []*Media
	Notes        []*Note
	People       []*Person
	Places       []*Place
	Repositories []*Repository
	Sources      []*Source
	Tags         []*Tag
}

type Citation struct{}
type Event struct{}
type Family struct{}
type Media struct{}
type Note struct{}
type Person struct{}
type Place struct{}
type Repository struct{}
type Source struct{}
type Tag struct{}
