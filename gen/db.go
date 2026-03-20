package gen

type DB struct {
	Header struct {
		Created struct {
			Date    string
			Version string
		}
		Researcher struct{}
		MediaPath  string
	}
	HomePerson *Person
	People     map[string]*Person
}

func NewDB() *DB {
	return &DB{
		People: map[string]*Person{},
	}
}
