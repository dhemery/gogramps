package gen

type DB struct {
	HomePerson *Person
	People map[string]*Person
}

func NewDB() *DB {
	return &DB{
		People: map[string]*Person{},
	}
}
