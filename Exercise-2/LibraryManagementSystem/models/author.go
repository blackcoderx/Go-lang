package models

type Author struct {
	FirstName string
	LastName  string
	Country   string
}

func (a Author) FullName() string {
	return a.FirstName + " " + a.LastName
}
