// Package convert converts Gramps data into Gen data.
package convert

import (
	"fmt"
	"time"

	"dhemery.com/gogramps/gen"
	"dhemery.com/gogramps/gramps"
)

func Convert(in *gramps.DB) (*gen.DB, error) {
	c := &converter{
		Gramps: NewGrampsMap(in),
		Gen:    gen.NewDB(),
	}
	return c.convert()
}

type converter struct {
	Gramps *GrampsMap
	Gen    *gen.DB
}

func (c *converter) convert() (*gen.DB, error) {
	out := c.Gen
	for handle := range c.Gramps.People {
		out.People[handle] = &gen.Person{}
	}

	for handle, in := range c.Gramps.People {
		out, ok := c.Gen.People[handle]
		if !ok {
			return nil, fmt.Errorf("converting %v: gen DB has no such person", handle)
		}
		if err := c.convertPerson(in, out); err != nil {
			return nil, err
		}
	}

	return out, nil
}

func (c *converter) convertPerson(in *gramps.Person, out *gen.Person) error {
	inName := in.Name
	outName := gen.PersonName{
		First:   inName.First,
		Surname: inName.Surname,
		Suffix:  inName.Suffix,
		Call:    inName.Call,
		Nick:    inName.Nick,
	}

	out.Primary = convertPrimary(in.Primary)
	out.Name = outName
	out.Gender = in.Gender
	return nil
}

func convertPrimary(in gramps.Primary) gen.Primary {
	return gen.Primary{
		Handle:  in.Handle,
		ID:      in.ID,
		Changed: time.Unix(int64(in.Change), 0),
	}
}
