package gen

import "time"

type Primary struct {
	ID      string    `json:"id"`
	Changed time.Time `json:"changed"`
}
