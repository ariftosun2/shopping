package main

import (
	"fmt"
	"fmt"
	"log"
	"net/http"
	"shopping-servis/db/dto"
	"shopping-servis/db/store"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

var q store.ShoppingRepo

func main() {
	router := gin.Default()

	store.OpenConnection()

	//authorization system
	protected := router.Group("/", authorizationMiddleware)
	//they will be loaded when opening the page
	router.GET("/", HelloWeb)
	//books
	router.GET("/booksGet", booksGet)
	protected.POST("/booksPost", booksPost)
	protected.PATCH("/booksUpdate/:id", booksUpdate)

	//users
	router.POST("/usersLogin", userLogin)
	router.POST("/usersRecord", userRecord)
	protected.POST("/logout", logout)

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

func HelloWeb(c *gin.Context) {
	booksrespons := q.BooksGet()
	c.JSON(http.StatusOK, gin.H{"data": booksrespons})
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
func booksUpdate(c *gin.Context) {
	id := c.Param("id")
	var updateuser *dto.Books

	if err := c.BindJSON(&updateuser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	books := &dto.Books{
		BooksKind: updateuser.BooksKind,
		Name:      updateuser.Name,
		Detail:    updateuser.Detail,
	}
	result := q.BooksUpdate(*books, id)
	c.JSON(http.StatusOK, gin.H{"data": result})
}

func userRecord(c *gin.Context) {
	var postuser *dto.User

	// Call BindJSON to bind the received JSON to
	// userpost.

	if err := c.BindJSON(&postuser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	password, _ := HashPassword(postuser.UserPassword)

	users := &dto.User{
		UserName:     postuser.UserName,
		UserLastName: postuser.UserLastName,
		UserPassword: password,
	}

	a := q.UserValidate(users)
	var result *dto.ResponsUser
	if a {
		result = q.CreateUsersItem(*users)
	}
	//dto create

	fmt.Println(result)
	c.JSON(http.StatusOK, gin.H{"data": result})
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
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
	token, userLogin := q.UserLogin(users)
	fmt.Println(userLogin)
	c.SetCookie("username", loginuser.UsersName, 3600, "", "", false, true)
	c.SetCookie("token", token, 3600, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{"token": token})

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
func logout(c *gin.Context) {
	user, err1 := c.Cookie("username")
	token, err2 := c.Cookie("token")

	if err1 == nil && err2 == nil {
		// Clear the cookies and
		// respond with an HTTP success status
		c.SetCookie("username", "", -1, "", "", false, true)
		c.SetCookie("token", "", -1, "", "", false, true)
		fmt.Println("cerezler silindi", token, ":", user, ":")
		c.JSON(http.StatusOK, nil)
	} else {
		// Respond with an HTTP error
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}
