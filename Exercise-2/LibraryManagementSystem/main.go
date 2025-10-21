package main

import (
	"LibraryManagementSystem/models"
	"fmt"
)

func makeLateFeeCalculator(dailyRate float64) func(int) float64 {
	return func(daysLate int) float64 {
		return float64(daysLate) * dailyRate
	}
}

func main() {
	bookOne := models.Book{
		Title:     "Go lang for dummies",
		ISBN:      "AB001",
		Pages:     550,
		Available: true,
		Author: models.Author{
			FirstName: "Emmanuel",
			LastName:  "Ziggah",
			Country:   "Ghana",
		},
	}

	bookTwo := models.Book{
		Title:     "Rust for dummies",
		ISBN:      "AB001",
		Pages:     1100,
		Available: true,
		Author: models.Author{
			FirstName: "John",
			LastName:  "Doe",
			Country:   "USA",
		},
	}

	bookOne.CheckOut()

	library := models.NewLibrary("City Library")
	library.AddBook(bookOne)
	library.AddBook(bookTwo)

	fmt.Printf(" -----------------Books--------------------------------------------- \n")
	fmt.Println(bookOne.DisplayInfo())
	fmt.Println(bookTwo.DisplayInfo())
	fmt.Printf("--------------------------Total Books in My library ------------------\n")
	fmt.Println("There are ", library.TotalBooks(), "books in my library")
	fmt.Printf("\n")
	fmt.Printf("---------------------------Available Books --------------------------- \n")
	fmt.Printf("\n")
	for _, book := range library.AvailableBooks() {
		fmt.Println(book.DisplayInfo())
	}
	fmt.Printf("--------------------------------------------------------------------\n")

	book, found := library.FindByISBN("AB001")
	if found {

		fmt.Println("Found Book with ISBN - ", book.ISBN, "\n", book.DisplayInfo())
	}

	fmt.Printf("------------------------charges for keeping books-------------------------\n")
	lateFeeCalculator := makeLateFeeCalculator(0.25)

	daysLate := 5
	lateFee := lateFeeCalculator(daysLate)

	fmt.Println("GHS", lateFee)
}
