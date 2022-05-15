package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"shopping-servis/db/dto"
	"shopping-servis/db/store"
	"strings"
	"time"
)

var q store.ShoppingRepo

func main() {
	router :=gin.Default()

	store.OpenConnection()
	//router

	protected := router.Group("/", authorizationMiddleware)

	//books
	router.POST("/booksPost", booksPost)
	router.GET("/booksGet", booksGet)

	//users
	router.POST("/usersLogin", userLogin)
	protected.POST("/usersPost", userPost)
	protected.GET("/usersGet", userGet)
	protected.PATCH("/usersUpdate/:id", userUpdate)
	protected.DELETE("/usersDelete/:id", userDelete)
	/* router.Run("localhost:8080") */

	httpServer := &http.Server{
		Handler:      router,
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
	}
	log.Fatal(httpServer.ListenAndServe())
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
		UserLastName: postuser.UserLastName,
		UserPassword: postuser.UserPassword,
	}
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
		UserLastName: updateuser.UserLastName,
		UserPassword: updateuser.UserPassword,
	}
	result := q.UserUpdate(*users, id)
	c.JSON(http.StatusOK, gin.H{"data": result})

}
func userDelete(c *gin.Context) {
	id := c.Param("id")
	result := q.UserDelete(id)
	c.JSON(http.StatusOK, gin.H{"data": result})
}

func userLogin(c *gin.Context) {

	var loginuser *dto.LoginUser

	if err := c.BindJSON(&loginuser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	users := &dto.LoginUser{
		UsersName:    loginuser.UsersName,
		UserPassword: loginuser.UserPassword,
	}
	//dto create
	result := q.UserLogin(users)
	fmt.Println(result)
	c.JSON(http.StatusOK, gin.H{"token": result})

}
func authorizationMiddleware(c *gin.Context) {
	s := c.Request.Header.Get("Authorization")

	token := strings.TrimPrefix(s, "Bearer ")

	if err := validateToken(token); err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
}

func validateToken(token string) error {
	_, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return []byte("secretkey"), nil
	})

	return err
}
