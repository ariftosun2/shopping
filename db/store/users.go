package store

import (
	"context"
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
	}
}

func (r *ShoppingRepo) insertUsersItem(request dto.User) (string, error) {
	Db := OpenConnection()

	sql := "INSERT INTO user_item(username,lastname,created) VALUES($1,$2,$3) returning id;"
	row := Db.QueryRowContext(context.Background(), sql, request.UserName, request.UserLastName, time.Now())
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

		err := rows.Scan(&response.Id, &response.UserName, &response.UserLastName, &response.Time)
		if err != nil {
			log.Fatalf(err.Error())
		}
		requests = append(requests, response)
	}

	return &requests
}
func (r *ShoppingRepo) UserUpdate(request dto.User, id string) *[]dto.ResponsUser {
	Db := OpenConnection()
	sqlStatement := `UPDATE user_item SET  username= $2, lastname = $3,created=$4 WHERE id = $1 RETURNING id,username,lastname,created;`
	rows, err := Db.Query(sqlStatement, id, request.UserName, request.UserLastName, time.Now())
	if err != nil {
		log.Fatal(err)
	}
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
