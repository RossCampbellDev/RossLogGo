package types

type Entry struct {
	ID        string
	Title     string
	Body      string
	Tags      []string
	Datestamp string
}

type User struct {
	ID       string
	Username string
	Passhash string
}
