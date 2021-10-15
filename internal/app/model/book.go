package model

type Book struct {
	ID          int64  `json:"id" db:"id"`
	Title       string `json:"title" db:"title"`
	Author      string `json:"author" db:"author"`
	Genre       string `json:"genre" db:"genre"`
	Price       int64  `json:"price" db:"price"`
	Description string `json:"description" db:"description"`
}
