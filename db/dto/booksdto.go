package dto

import "time"


type Books struct{
	BooksKind string `json:"bookskind"`
	Name  string `json:"name"`
	Detail string `json:"detail"`
}
type ResponsBooks struct{
	Id string `json:"Id"`
	BooksKind string `json:"bookskind"`
	Name  string `json:"name"`
	Detail string `json:"detail"`
	Time time.Time `json:"time"`
}



/* type BooksKind struct{
	Horror string `json:"horror"`
	Adventure string `json:"adventure"`
} */