package main

import (
	"LibraryManagementSystem/models"
	"fmt"
)

func main() {
	bookOne := models.Book{
		Title:     "Go lang for dummies",
		ISBN:      "123455678AB",
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
		ISBN:      "98765432AB",
		Pages:     1100,
		Available: true,
		Author: models.Author{
			FirstName: "John",
			LastName:  "Doe",
			Country:   "USA",
		},
	}

	bookOne.CheckOut()

	fmt.Printf(" -----------------Books--------------------------------------------- \n")
	fmt.Println(bookOne.DisplayInfo())
	fmt.Println(bookTwo.DisplayInfo())
	fmt.Printf("--------------------------------------------------------------------")
}
