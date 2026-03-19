package gramps

import (
	"fmt"
	"time"

	"dhemery.com/gogramps/gen"
)

func convert(in *DB) (*gen.DB, error) {
	c := &converter{
		Gramps: newPrimaryMap(in),
		Gen:    gen.NewDB(),
	}
	return c.convert()
}

type converter struct {
	Gramps *primaryMap
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

func (c *converter) convertPerson(in *Person, out *gen.Person) error {
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

func convertPrimary(in Primary) gen.Primary {
	return gen.Primary{
		ID:      in.ID,
		Changed: time.Unix(int64(in.Change), 0),
	}
}
