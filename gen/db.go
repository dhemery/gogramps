package gen

type DB struct {
	People map[string]*Person
}

func NewDB() *DB {
	return &DB{
		People: map[string]*Person{},
	}
}
