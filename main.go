package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"shopping-servis/db/dto"
	"shopping-servis/db/store"
)

var q store.BooksRepo

func main() {
	router := gin.Default()

	store.OpenConnection()
	//router
	router.POST("/booksPost", booksPost)

	router.Run("localhost:8080")
}

func booksPost(c *gin.Context) {
	var postbooks *dto.Books

	// Call BindJSON to bind the received JSON to
	// bookspost.

	if err := c.BindJSON(&postbooks); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	//dto create
	book := &dto.Books{BooksKind: postbooks.BooksKind, Name: postbooks.Name, Detail: postbooks.Detail}

	result := q.CreateBooksItem(*book)
	fmt.Println(result)
	c.JSON(http.StatusOK, gin.H{"data": result})
}
