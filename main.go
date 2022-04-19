package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"shopping-servis/db/dto"
	"shopping-servis/db/store"
)

var dbm *sql.DB
var q store.BooksRepo

func main() {
	fmt.Println("Server calisistiriliyor")

	store.Dbconfig()
	defer dbm.Close()

	router := gin.Default()

	router.POST("/booksPost", booksPost)

	router.Run("localhost:8080")
}

func booksPost(c *gin.Context) {
	var postbooks *dto.Books

	// Call BindJSON to bind the received JSON to
	// newAlbum.

	if err := c.BindJSON(&postbooks); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	book := &dto.Books{BooksKind: postbooks.BooksKind, Name: postbooks.Name, Detail: postbooks.Detail}

	result := q.CreateBooksItem(*book)

	c.JSON(http.StatusOK, gin.H{"data": result})
}
