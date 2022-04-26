package store

import (
	"context"
	"log"
	"shopping-servis/db/dto"
	"time"
)

type DBTX interface {
	QueryRowContext()
	QueryRow()
}

type ShoppingRepo struct {
	Create DBTX
}

func (r *ShoppingRepo) CreateBooksItem(request dto.Books) *dto.ResponsBooks {
	var lastInsertID string
	var err error
	if lastInsertID, err = r.insertBooksItem(request); err != nil {
		return nil
	}
	return &dto.ResponsBooks{
		Id:        lastInsertID,
		BooksKind: request.BooksKind,
		Name:      request.Name,
		Detail:    request.Detail,
	}
}



func (r *ShoppingRepo) insertBooksItem(request dto.Books) (string, error) {
	 Db := OpenConnection()

	sql := "INSERT INTO books_item(bookskind,bookname,detail,created) VALUES($1,$2,$3,$4) returning id;"
	row := Db.QueryRowContext(context.Background(), sql, request.BooksKind, request.Name, request.Detail, time.Now())
	defer Db.Close()
	var lastInsertID string

	return lastInsertID, row.Scan(&lastInsertID)
}
func (r *ShoppingRepo) BooksGet() *[]dto.ResponsBooks {
	Db := OpenConnection()
	sql := "Select * from books_item"

	rows, _ := Db.Query(sql)

	var requests []dto.ResponsBooks

	defer Db.Close()
	for rows.Next() {
		response := dto.ResponsBooks{}
		// var fileResponse []string
		
		err:=rows.Scan(&response.Id, &response.BooksKind, &response.Name, &response.Detail,&response.Time)
		if err!=nil{
			log.Fatalf(err.Error())
		}
		requests = append(requests, response)
	}

	return &requests
}
