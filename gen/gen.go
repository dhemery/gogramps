// Package gen represents all of the genealogy data.
package gen

import (
	"time"
)

type Gen struct {
	Citations    []*Citation   `json:"citations"`
	Events       []*Event      `json:"events"`
	Families     []*Family     `json:"families"`
	Media        []*Media      `json:"media"`
	Notes        []*Note       `json:"notes"`
	People       []*Person     `json:"people"`
	Places       []*Place      `json:"places"`
	Repositories []*Repository `json:"repositories"`
	Sources      []*Source     `json:"sources"`
	Tags         []*Tag        `json:"tags"`
}

type TableObject struct {
	Handle string    `json:"handle"`
	Change time.Time `json:"change"`
}

type GrampsObject struct {
	TableObject
	ID string `json:"gramps_id"`
}

type Family struct{}
type Media struct{}
type Note struct{}
type Person struct{}
type Repository struct{}
type Tag struct{}
