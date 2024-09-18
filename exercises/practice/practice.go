package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Book struct {
	id               int
	title            string
	shortDescription string
	price            float64
}

func NewBook(id int, title string, shortDescription string, price float64) *Book {
	newBook := Book{
		id:               id,
		title:            title,
		shortDescription: shortDescription,
		price:            price,
	}
	return &newBook
}

func (book *Book) outputDetails() {
	fmt.Printf("The %v is %v with id of %v costs %v", book.title, book.shortDescription, book.id, book.price)
}

var reader = bufio.NewReader(os.Stdin)

func main() {
	var newBook Book
	newBook = Book{
		24,
		"A book",
		"Very interesting book",
		25.4,
	}
	fmt.Println(newBook)
	anotherBook := NewBook(29, "The second Book", "Not that much interesting", 9.01)
	anotherBook.outputDetails()

	var userBook Book
	id := getUserData("Enter the book ID: ")
	idNum, _ := strconv.Atoi(id)
	title := getUserData("Enter the title of the book: ")
	description := getUserData("Enter the description of the book: ")
	price := getUserData("Enter the price of the book: ")
	priceFloat, _ := strconv.ParseFloat(price, 64)
	userBook = Book{idNum, title, description, priceFloat}
	userBook.outputDetails()

}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	userData, _ := reader.ReadString('\n')
	cleanedData := strings.Replace(userData, "\n", "", -1)
	return cleanedData
}
