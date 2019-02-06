package database

type Bot struct {
	Addr   string
	Pwd    string
	Port   string
	Status bool
}

type ListedBot struct {
	ID     uint64
	Addr   string
	Pwd    string
	Port   string
	Status bool
}
