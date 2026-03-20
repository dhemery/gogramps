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
	out.HomePerson = out.People[c.Gramps.HomePerson.Handle]

	return out, nil
}

func (c *converter) convertPerson(in *Person, out *gen.Person) error {
	for _, inName := range in.Names {
		outName := convertPersonName(inName)
		out.Names = append(out.Names, outName)
	}

	out.Primary = convertPrimary(in.PrimaryRecord)
	out.Gender = in.Gender
	return nil
}

func convertPersonName(in PersonName) gen.PersonName{
	return gen.PersonName{
		Private: in.Private,
		Title:   in.Title,
		First:   in.First,
		Surname: in.Surname,
		Suffix:  in.Suffix,
		Call:    in.Call,
		Nick:    in.Nick,
	}
}

func convertPrimary(in PrimaryRecord) gen.Primary {
	return gen.Primary{
		ID:      in.ID,
		Changed: time.Unix(int64(in.Change), 0),
	}
}

