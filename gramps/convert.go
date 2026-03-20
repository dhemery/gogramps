package gramps

import (
	"dhemery.com/gogramps/gen"
)

func convert(in *DB) (*gen.DB, error) {
	c := &converter{
		Gramps: in,
		Gen:    gen.NewDB(),
	}
	return c.convert()
}

type converter struct {
	Gramps *DB
	Gen    *gen.DB
}

func (c *converter) convert() (*gen.DB, error) {
	out := c.Gen
	return out, nil
}
