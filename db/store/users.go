package store

import (
	"context"
	"shopping-servis/db/dto"
	"time"
	"log"
)


func (r *ShoppingRepo) CreateUsersItem(request dto.User) *dto.ResponsUser {
	var lastInsertID string
	var err error
	if lastInsertID, err = r.insertUsersItem(request); err != nil {
		return nil
	}
	return &dto.ResponsUser{
		Id:        lastInsertID,
		UserName: request.UserName,
		UserLastName:      request.UserLastName,
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
		// var fileResponse []string
		
		err:=rows.Scan(&response.Id, &response.UserName, &response.UserLastName,&response.Time)
		if err!=nil{
			log.Fatalf(err.Error())
		}
		requests = append(requests, response)
	}

	return &requests
}
