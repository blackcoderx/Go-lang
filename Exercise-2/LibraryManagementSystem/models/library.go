package models

type Library struct {
	Name  string
	Books []Book
}

func (l *Library) AddBook(book Book) {
	l.Books = append(l.Books, book)
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
