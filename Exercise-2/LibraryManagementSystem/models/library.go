package models

import "fmt"

type Library struct {
	Name        string
	Books       []Book
	BooksByISBN map[string]*Book
}

func NewLibrary(name string) *Library {
	return &Library{
		Name:        name,
		Books:       []Book{},
		BooksByISBN: make(map[string]*Book),
	}
}
func (l *Library) AddBook(book Book) {
	defer fmt.Println("Book added:", book.Title)

	l.Books = append(l.Books, book)
	l.BooksByISBN[book.ISBN] = &l.Books[len(l.Books)-1]
	// "&" saves the memory address of the Book at len-1
	// so any changes that happens to the orginal book affects the varible
}

func (l Library) TotalBooks() int {
	return len(l.Books)
}

func (l Library) AvailableBooks() []Book {
	availableBooks := make([]Book, 0)

	for _, book := range l.Books {
		if book.Available {
			availableBooks = append(availableBooks, book)
		}
	}
	return availableBooks
}

func (l *Library) FindByISBN(isbn string) (*Book, bool) {
	book, exists := l.BooksByISBN[isbn]
	return book, exists
}
