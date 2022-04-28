package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"shopping-servis/db/dto"
	"shopping-servis/db/store"
)

var q store.ShoppingRepo

func main() {
	router := gin.Default()

	store.OpenConnection()
	//router

	//books
	router.POST("/booksPost", booksPost)
	router.GET("/booksGet", booksGet)

	//users
	router.POST("/usersPost", userPost)
	router.GET("/usersGet", userGet)
	router.PATCH("/usersUpdate/:id", userUpdate)
	router.DELETE("/usersDelete/:id", userDelete)
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

func booksGet(c *gin.Context) {
	booksrespons := q.BooksGet()
	c.JSON(http.StatusOK, gin.H{"data": booksrespons})
}
func userPost(c *gin.Context) {
	var postuser *dto.User

	// Call BindJSON to bind the received JSON to
	// userpost.

	if err := c.BindJSON(&postuser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	users := &dto.User{
		UserName:     postuser.UserName,
		UserLastName: postuser.UserLastName}
	//dto create
	result := q.CreateUsersItem(*users)
	fmt.Println(result)
	c.JSON(http.StatusOK, gin.H{"data": result})
}

func userGet(c *gin.Context) {
	usersrespons := q.UserGet()
	c.JSON(http.StatusOK, gin.H{"data": usersrespons})
}
func userUpdate(c *gin.Context) {
	id := c.Param("id")
	var updateuser *dto.User

	if err := c.BindJSON(&updateuser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	users := &dto.User{
		UserName:     updateuser.UserName,
		UserLastName: updateuser.UserLastName}
	result := q.UserUpdate(*users, id)
	c.JSON(http.StatusOK, gin.H{"data": result})

}
func userDelete(c *gin.Context) {
	id := c.Param("id")
	result := q.UserDelete(id)
	c.JSON(http.StatusOK, gin.H{"data": result})

}
