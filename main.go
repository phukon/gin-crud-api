package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

/**
* when it's being serialized, the field names are converted to lowercase
* or when taking the json object and converting it into a book struct
* do the opposite
*
* The reason these are uppercase because it makes them all exported fields
* That means it's kinda like a public field, it can be viewed outside of the
* the current module.
 */

type book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

var books = []book{
	{ID: "1", Title: "Guns Germs and Steel", Author: "Jared Diamond", Quantity: 12},
	{ID: "2", Title: "Kafka on The Shore", Author: "Haruki Murakami", Quantity: 4},
	{ID: "3", Title: "Norwegian Woods", Author: "Haruki Murakami", Quantity: 5},
	{ID: "4", Title: "1984", Author: "George Orwell", Quantity: 2},
}

/**
* The context holds all the request headers, query params and related stuff.
 */

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func getBookById(id string) (*book, error) {
	for i, b := range books {
		if b.ID == id {
			return &books[i], nil
		}
	}

	return nil, errors.New("The requested book not found")
}

func bookById(c *gin.Context) {
	id := c.Param("id")
	book, err := getBookById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, book)
}

func createBook(c *gin.Context) {
	var newBook book

	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.GET("/books/:id", bookById)
	router.POST("/books", createBook)
	router.Run("localhost:8080")
}
