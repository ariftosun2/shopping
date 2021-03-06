package store

import (
	"context"
	"log"
	"shopping-servis/db/dto"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

func (r *ShoppingRepo) CreateUsersItem(request dto.User) *dto.ResponsUser {
	var lastInsertID string
	var err error
	if lastInsertID, err = r.insertUsersItem(request); err != nil {
		return nil
	}
	return &dto.ResponsUser{
		Id:           lastInsertID,
		UserName:     request.UserName,
		UserLastName: request.UserLastName,
		UserPassword: request.UserPassword,
	}
}
func ValidateRecord(user *dto.User) (*[]dto.ResponsUser, error) {
	Db := OpenConnection()

	rows, err := Db.Query("SELECT * FROM user_item WHERE username = $1", user.UserName)
	if err != nil {
		return nil, err
	}

	var requests []dto.ResponsUser

	defer Db.Close()
	for rows.Next() {
		response := dto.ResponsUser{}

		err := rows.Scan(&response.Id, &response.UserName, &response.UserLastName, &response.UserPassword, &response.Time)
		if err != nil {
			log.Fatalf(err.Error())
		}
		requests = append(requests, response)
	}

	return &requests, nil
}

func (r *ShoppingRepo) UserValidate(user *dto.User) bool {
	registy, err := ValidateRecord(user)
	if err != nil {
		panic(err)
	}
	for _, v := range *registy {
		if v.UserName != "" {
			return false
		}
	}

	return true
}

func (r *ShoppingRepo) insertUsersItem(request dto.User) (string, error) {
	Db := OpenConnection()

	sql := "INSERT INTO user_item(username,lastname,userpassword,created) VALUES($1,$2,$3,$4) returning id;"
	row := Db.QueryRowContext(context.Background(), sql, request.UserName, request.UserLastName, request.UserPassword, time.Now())
	defer Db.Close()
	var lastInsertID string

	return lastInsertID, row.Scan(&lastInsertID)
}

func (r *ShoppingRepo) UserGet() *[]dto.ResponsUser {
	Db := OpenConnection()
	sql := "Select * from user_item"

	rows, _ := Db.Query(sql)

	var requests []dto.ResponsUser

	defer Db.Close()
	for rows.Next() {
		response := dto.ResponsUser{}

		err := rows.Scan(&response.Id, &response.UserName, &response.UserLastName, &response.UserPassword, &response.Time)
		if err != nil {
			log.Fatalf(err.Error())
		}
		requests = append(requests, response)
	}

	return &requests
}
func (r *ShoppingRepo) UserUpdate(request dto.User, id string) *[]dto.ResponsUser {
	Db := OpenConnection()
	sqlStatement := `UPDATE user_item SET  username= $2, lastname = $3,created=$5,userpassword=$4 WHERE id = $1 RETURNING id,username,lastname,created;`
	rows, err := Db.Query(sqlStatement, id, request.UserName, request.UserLastName, request.UserPassword, time.Now())
	if err != nil {
		log.Fatal(err)
	}
	var requests []dto.ResponsUser

	defer Db.Close()
	for rows.Next() {
		response := dto.ResponsUser{}

		err := rows.Scan(&response.Id, &response.UserName, &response.UserLastName, &response.UserPassword, &response.Time)
		if err != nil {
			log.Fatalf(err.Error())
		}
		requests = append(requests, response)
	}

	return &requests
}
func (r *ShoppingRepo) UserDelete(id string) *[]dto.ResponsUser {
	Db := OpenConnection()
	sql := `DELETE FROM user_item WHERE id=$1;`
	rows, err := Db.Query(sql, id)
	if err != nil {
		log.Fatal(err)
	}
	defer Db.Close()
	var requests []dto.ResponsUser

	defer Db.Close()
	for rows.Next() {
		response := dto.ResponsUser{}

		err := rows.Scan(&response.Id, &response.UserName, &response.UserLastName, &response.Time)
		if err != nil {
			log.Fatalf(err.Error())
		}

		requests = append(requests, response)
	}
	return &requests
}

// Check password data and local control
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
func ValidateLogin(user *dto.LoginUser) (*[]dto.ResponsUser, error) {
	Db := OpenConnection()

	rows, err := Db.Query("SELECT * FROM user_item WHERE username = $1", user.UsersName)
	if err != nil {
		return nil, err
	}

	var requests []dto.ResponsUser

	defer Db.Close()
	for rows.Next() {
		response := dto.ResponsUser{}

		err := rows.Scan(&response.Id, &response.UserName, &response.UserLastName, &response.UserPassword, &response.Time)
		if err != nil {
			log.Fatalf(err.Error())
		}
		requests = append(requests, response)
	}

	return &requests, nil
}

//login username and password check
func (r *ShoppingRepo) compareRecords(user *dto.LoginUser) (bool, *dto.LoginUser) {
	registy, err := ValidateLogin(user)
	if err != nil {
		panic(err)
	}
	for _, y := range *registy {

		if y.UserName != "" {
			checkpassword := CheckPasswordHash(user.UserPassword, y.UserPassword)
			if checkpassword {
				return true, &dto.LoginUser{
					UsersName:    y.UserName,
					UserPassword: y.UserPassword,
				}
			}
		}

	}
	return false, &dto.LoginUser{}
}

//login result token
func (r *ShoppingRepo) UserLogin(user *dto.LoginUser) (string, *dto.LoginUser) {

	var token string
	var b bool
	var userlogin *dto.LoginUser
	b, userlogin = r.compareRecords(user)

	if b {
		token, _ = GenerateJWT(userlogin)

	}
	return token, userlogin
}

//token generate
func GenerateJWT(user *dto.LoginUser) (string, error) {
	var mySigningKey = []byte("secretkey")
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["username"] = user.UsersName
	claims["userpassword"] = user.UserPassword
	claims["exp"] = time.Now().Add(time.Minute * 20).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		log.Fatalf("Something Went Wrong: %s", err.Error())
		return "", err
	}
	return tokenString, nil
}
