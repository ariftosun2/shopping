package store

import (
	"context"
	"shopping-servis/db/dto"
	"time"
)

type DBTX interface {
	QueryRowContext()
	QueryRow()
}

type BooksRepo struct {
	Create DBTX
}

func (r *BooksRepo) CreateBooksItem(request dto.Books) *dto.ResponsBooks {
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

func (r *BooksRepo) insertBooksItem(request dto.Books) (string, error) {
	Db := OpenConnection()

	sql := "INSERT INTO books_item(bookskind,bookname,detail,created) VALUES($1,$2,$3,$4) returning id;"
	row := Db.QueryRowContext(context.Background(), sql, request.BooksKind, request.Name, request.Detail, time.Now())
	var lastInsertID string

	return lastInsertID, row.Scan(&lastInsertID)
}
