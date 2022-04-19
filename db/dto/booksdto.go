package dto


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
}



/* type BooksKind struct{
	Horror string `json:"horror"`
	Adventure string `json:"adventure"`
} */