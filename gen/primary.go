package gen

import "time"

type Primary struct {
	Handle  string    `json:"handle"`
	ID      string    `json:"id"`
	Changed time.Time `json:"changed"`
}
