package store

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"log"
	"shopping-servis/db/dto"
	"time"
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

func (r *ShoppingRepo) compareRecords(user *dto.LoginUser) bool {
	registy := r.UserGet()
	for _, y := range *registy {
		if y.UserName == user.UsersName && y.UserPassword == user.UserPassword {
			return true
		}
	}
	return false
}

func (r *ShoppingRepo) UserLogin(user *dto.LoginUser) string {
	var token string
	var userlogin dto.LoginUser
	b := r.compareRecords(user)
	if b {
		token, _ = GenerateJWT(&userlogin)

	}
	return token
}
func GenerateJWT(user *dto.LoginUser) (string, error) {
	var mySigningKey = []byte(user.UsersName)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["username"] = user.UsersName
	claims["userpassword"] = user.UserPassword
	claims["exp"] = time.Now().Add(time.Minute * 20).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		fmt.Errorf("Something Went Wrong: %s", err.Error())
		return "", err
	}
	return tokenString, nil
}
