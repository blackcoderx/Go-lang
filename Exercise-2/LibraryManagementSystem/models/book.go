package models

import "fmt"

type Book struct {
	Title     string
	ISBN      string
	Pages     int
	Available bool
	Author    Author
}

func (b Book) DisplayInfo() string {
	status := "Available"
	authorFullName := b.Author.FullName()
	if !b.Available {
		status = "Checked Out"
	}

	return fmt.Sprintf("%s by %s (ISBN: %s, Pages: %d, Status: %s)\n ",
		b.Title,
		authorFullName,
		b.ISBN,
		b.Pages,
		status,
	)
}

func (b *Book) CheckOut() {
	defer fmt.Println("Checkout complete:", b.Title)

	b.Available = false
}
