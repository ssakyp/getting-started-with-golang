package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Product struct {
	id               string
	title            string
	description 	 string
	price            float64
}

func NewProduct(id string, t string, s string, p float64) *Product {
	return &Product{id, t, s, p}
}

func (prod *Product) printData() {
	fmt.Printf("ID: %v\n", prod.id)
	fmt.Printf("Title: %v\n", prod.title)
	fmt.Printf("Description: %v\n", prod.description)
	fmt.Printf("Price:  USD %.2f\n", prod.price)
}

var reader = bufio.NewReader(os.Stdin)

func main() {
	firstProduct := Product {
		"first-product",
		"A book",
		"A nice book",
		10.99
	}
	secondProduct := NewProduct("second-product", "A Carpet", "A beautiful carpet", 99.99)
	//fmt.Println(firstProduct)
	//fmt.Println(secodnProduct)

	firstProduct.printData()
	secondProduct.printData()
	
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
