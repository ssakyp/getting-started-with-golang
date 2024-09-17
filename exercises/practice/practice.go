package main

type Book struct {
  id int
  title string
  shortDescription string
  price float64
}

func NewBook(id int, title string, shortDescription string, price float64) *Book {
  newBook := Book {
    id: id
    title: title
    shortDescription: shortDescription
    pirce: price
  }
  return &newBook
}

func (book *Book) outputDetails() {
  fmt.Printf("The %v is %v with id of %v costs %v", book.title, book.shortDescription, book.id, book.price)
}

func main() {
  var newBook Book {
    24,
    "A book"
    "Very interesting book"
    25.4
  }
  fmt.Println(newBook)
  anotherBook := NewBook(29, "The second Book", "Not that much interesting", 9.01)
  anotherBook.outputDetails()
}
